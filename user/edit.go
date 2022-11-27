package user

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/here"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

type EditUser struct {
	Images *multipart.FileHeader `form:"image"`
}

func Edit(c *gin.Context) {
	imageBase := os.Getenv("HERE_IMAGE_PATH")

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var editUser EditUser
	err := c.Bind(&editUser)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	profileImage := here.CreateUniqueFileName(editUser.Images.Filename)

	if err := c.SaveUploadedFile(editUser.Images, imageBase+"/"+profileImage); err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	var user response.EditUser

	mysqlClient := repository.Mysql()
	err = mysqlClient.Model(&models.User{}).Where("uid = ?", uid).Update("profile_image", profileImage).Scan(&user).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.EditUserInformation(c, user, status.StatusOK)
}
