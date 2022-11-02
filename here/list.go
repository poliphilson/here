package here

import (
	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func List(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var heres []models.Here

	mysqlClient := repository.Mysql()
	err := mysqlClient.Where("uid = ?", uid).Find(&heres).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		return
	}

	response.HeresOnMap(c, heres, status.StatusOK)
}
