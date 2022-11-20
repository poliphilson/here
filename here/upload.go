package here

import (
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/poliphilson/here/datatype"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"gorm.io/gorm"
)

type GetHere struct {
	Contents   string                  `form:"contents"`
	X          float64                 `form:"x"`
	Y          float64                 `form:"y"`
	IsPrivated bool                    `form:"is_privated"`
	Images     []*multipart.FileHeader `form:"image[]"`
	Videos     []*multipart.FileHeader `form:"video[]"`
}

func Upload(c *gin.Context) {
	imageBase := os.Getenv("HERE_IMAGE_PATH")
	videoBase := os.Getenv("HERE_VIDEO_PATH")

	var getHere GetHere
	err := c.Bind(&getHere)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		return
	}

	hereForm := models.Here{}
	var imageArray []string
	var videoArray []string

	images := getHere.Images

	if len(images) > 0 {
		hereForm.Image = true

		for _, file := range images {
			fileName := filepath.Base(file.Filename)
			rename := CreateUniqueFileName(fileName)
			if err := c.SaveUploadedFile(file, imageBase+"/"+rename); err != nil {
				response.InternalServerError(c, status.InternalError)
				log.Println(err.Error())
				return
			}
			imageArray = append(imageArray, rename)
		}
	} else {
		hereForm.Image = false
	}

	videos := getHere.Videos

	if len(videos) > 0 {
		hereForm.Video = true

		for _, file := range videos {
			fileName := filepath.Base(file.Filename)
			rename := CreateUniqueFileName(fileName)
			if err := c.SaveUploadedFile(file, videoBase+"/"+rename); err != nil {
				response.InternalServerError(c, status.InternalError)
				log.Println(err.Error())
				return
			}
			videoArray = append(videoArray, rename)
		}
	} else {
		hereForm.Video = false
	}

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	hereForm.Uid = uid.(int)
	hereForm.Contents = getHere.Contents
	hereForm.Location = datatype.Location{X: getHere.X, Y: getHere.Y}
	hereForm.IsPrivated = getHere.IsPrivated

	simpleHere, err := createHere(hereForm, imageArray, videoArray)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.CreateHere(c, simpleHere, status.StatusOK)
}

func createHere(here models.Here, images []string, videos []string) (response.SimpleHere, error) {
	mysqlClient := repository.Mysql()
	simpleHere := response.SimpleHere{}

	err := mysqlClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&here).Scan(&simpleHere).Error
		if err != nil {
			return err
		}

		for _, imageName := range images {
			form := models.HereImage{
				Hid:   here.Hid,
				Image: imageName,
			}
			err := tx.Create(&form).Error
			if err != nil {
				return err
			}
		}

		for _, videoName := range videos {
			form := models.HereVideo{
				Hid:   here.Hid,
				Video: videoName,
			}
			err := tx.Create(&form).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return simpleHere, err
}

func CreateUniqueFileName(fileName string) string {
	prefix := uuid.New().String()
	suffix := time.Now().Format("20060102150405")
	return prefix + "-" + suffix + "-" + fileName
}
