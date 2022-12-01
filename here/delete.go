package here

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func Delete(c *gin.Context) {
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

	err = config.DB.Model(&models.Here{}).Where("uid = ? AND hid = ?", uid, hid).Update("is_deleted", true).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.Ok(c, status.StatusOK)
}
