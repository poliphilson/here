package here

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func Detail(c *gin.Context) {
	hid := c.Param("hid")
	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var detailHere response.DetailHere

	mysqlClient := repository.Mysql()
	err := mysqlClient.Model(&models.Here{}).Where("uid = ? AND hid = ? AND is_deleted = ?", uid, hid, false).Scan(&detailHere.Here).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	if detailHere.Here.Image {
		err := mysqlClient.Model(&models.HereImage{}).Select("image").Where("hid = ?", hid).Scan(&detailHere.Images).Error
		if err != nil {
			response.InternalServerError(c, status.InternalError)
			log.Println(err)
			return
		}
	}

	if detailHere.Here.Video {
		err := mysqlClient.Model(&models.HereVideo{}).Select("video").Where("hid = ?", hid).Scan(&detailHere.Videos).Error
		if err != nil {
			response.InternalServerError(c, status.InternalError)
			log.Println(err)
			return
		}
	}

	response.DetailHereOnMap(c, detailHere, status.StatusOK)
}
