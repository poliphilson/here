package here

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func List(c *gin.Context) {
	var err error
	date := c.Query("date")
	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var simpleHereList []response.SimpleHere

	if date != "" {
		err = config.DB.
			Model(&models.Here{}).
			Where("uid = ? AND is_deleted = ? AND DATE(created_at) = ? ", uid, false, date).
			Scan(&simpleHereList).Error
	} else {
		err = config.DB.
			Model(&models.Here{}).
			Where("uid = ? AND is_deleted = ?", uid, false).
			Scan(&simpleHereList).Error
	}
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.HeresOnMap(c, simpleHereList, status.StatusOK)
}
