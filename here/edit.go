package here

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"gorm.io/gorm"
)

type editHere struct {
	Contents   string                  `form:"contents"`
	IsPrivated bool                    `form:"is_privated"`
	Images     []string                `form:"images[]"`
	NewImages  []*multipart.FileHeader `form:"new_image[]"`
}

func Edit(c *gin.Context) {
	imageBase := os.Getenv("HERE_IMAGE_PATH")

	temp := c.Param("hid")
	hid, err := strconv.Atoi(temp)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	var editHere editHere
	err = c.Bind(&editHere)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	fmt.Println(editHere.Images)

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Here{}).Where("hid = ? AND uid = ?", hid, uid).Update("contents", editHere.Contents).Error
		if err != nil {
			return err
		}

		err = tx.Model(&models.Here{}).Where("hid = ? AND uid = ?", hid, uid).Update("is_privated", editHere.IsPrivated).Error
		if err != nil {
			return err
		}

		if len(editHere.Images) == 0 {
			err = tx.Where("hid = ?", hid).Delete(&models.HereImage{}).Error
			if err != nil {
				return err
			}
		} else {
			type imageList struct {
				Images []string
			}

			var images imageList
			err = tx.Model(&models.HereImage{}).Select("image").Where("hid = ?", hid).Scan(&images.Images).Error
			if err != nil {
				return err
			}
			for _, imageName := range images.Images {
				exists := searchImage(editHere.Images, imageName)
				if !exists {
					err = tx.Where("hid = ? AND image = ?", hid, imageName).Delete(&models.HereImage{}).Error
					if err != nil {
						return err
					}
				}
			}
		}

		newImages := editHere.NewImages
		if len(newImages) > 0 {
			err = tx.Model(&models.Here{}).Where("hid = ? AND uid = ?", hid, uid).Update("image", true).Error
			if err != nil {
				return err
			}
			for _, file := range newImages {
				fileName := filepath.Base(file.Filename)
				rename := CreateUniqueFileName(fileName)
				form := models.HereImage{
					Hid:   hid,
					Image: rename,
				}
				err = tx.Create(&form).Error
				if err != nil {
					return err
				}

				if err := c.SaveUploadedFile(file, imageBase+"/"+rename); err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.Ok(c, status.StatusOK)
}

func searchImage(imageList []string, image string) bool {
	for _, imageName := range imageList {
		if imageName == image {
			return true
		}
	}

	return false
}
