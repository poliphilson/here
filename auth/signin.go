package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"golang.org/x/crypto/bcrypt"
)

type signInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	var form signInForm

	err := c.BindJSON(&form)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("BindJson error.")
		log.Println(err.Error())
		return
	}

	var user models.User
	err = config.DB.Where("email = ?", form.Email).Find(&user).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Select email error.")
		log.Println(err.Error())
		return
	}

	if user.Email == "" {
		response.EmailDoesNotExist(c, status.NotRegisteredEmail)
		log.Println("Not registered email.")
		return
	}

	result := checkPassword(user.Password, form.Password)
	if !result {
		response.FailedSignIn(c, status.FailedSignIn)
		return
	}

	aToken, err := CreateAccessToken(user.UID)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Fail to create access token.")
		log.Println(err.Error())
		return
	}

	rToken, err := CreateRefreshToken()
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Fail to create refresh token.")
		log.Println(err.Error())
		return
	}

	c.SetCookie("access_token", aToken, 60*60*72, "/", "localhost", false, true)
	c.SetCookie("refresh_token", rToken, 60*60*72, "/", "localhost", false, true)

	var userInformation response.SignIn
	userInformation.ProfileImage = user.ProfileImage

	response.SuccessfullySignIn(c, userInformation, status.StatusOK)
}

func checkPassword(hashVal, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
