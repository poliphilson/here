package user

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/auth"
	"github.com/poliphilson/here/here"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

type UpdateUser struct {
	Images *multipart.FileHeader `form:"image"`
}

func Update(c *gin.Context) {
	imageBase := os.Getenv("HERE_IMAGE_PATH")

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	email, exists := c.Get("email")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var updateUser UpdateUser
	err := c.Bind(&updateUser)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	profileImage := here.CreateUniqueFileName(updateUser.Images.Filename)

	if err := c.SaveUploadedFile(updateUser.Images, imageBase+"/"+profileImage); err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	mysqlClient := repository.Mysql()

	err = mysqlClient.Model(&models.User{}).Where("uid = ?", uid).Update("profile_image", profileImage).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	aToken, err := auth.CreateAccessToken(uid.(int), email.(string), profileImage)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Fail to create access token.")
		log.Println(err.Error())
		return
	}

	c.SetCookie("access_token", aToken, 60*60*72, "/", "localhost", false, true)

	response.Ok(c, status.StatusOK)
}
