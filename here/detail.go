package here

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"gorm.io/gorm"
)

func Detail(c *gin.Context) {
	temp := c.Param("hid")
	hid, err := strconv.Atoi(temp)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var detailHere response.DetailHere

	mysqlClient := repository.Mysql()
	err = mysqlClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Here{}).Where("uid = ? AND hid = ? AND is_deleted = ?", uid, hid, false).Scan(&detailHere.Here).Error
		if err != nil {
			return err
		}

		err = tx.Model(&models.HereAddress{}).Where("hid = ?", hid).Scan(&detailHere.Address).Error
		if err != nil {
			return err
		}

		if detailHere.Here.Image {
			err := tx.Model(&models.HereImage{}).Select("image").Where("hid = ?", hid).Scan(&detailHere.Images).Error
			if err != nil {
				return err
			}
		}

		if detailHere.Here.Video {
			err := tx.Model(&models.HereVideo{}).Select("video").Where("hid = ?", hid).Scan(&detailHere.Videos).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err.Error())
		response.InternalServerError(c, status.InternalError)
		return
	}

	response.DetailHereOnMap(c, detailHere, status.StatusOK)
}
