package user

import (
	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func Detail(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	profileImage, exists := c.Get("profile_image")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var user response.User
	user.Email = email.(string)
	user.ProfileImage = profileImage.(string)

	response.UserInformation(c, user, status.StatusOK)
}
