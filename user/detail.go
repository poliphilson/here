package user

import (
	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func Detail(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var user response.User

	mysqlClient := repository.Mysql()
	err := mysqlClient.Model(&models.User{}).Where("uid = ?", uid).Scan(&user).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		return
	}

	response.UserInformation(c, user, status.StatusOK)
}
