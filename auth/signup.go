package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"golang.org/x/crypto/bcrypt"
)

type signUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var form signUpForm

	err := c.BindJSON(&form)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("BindJson error.")
		log.Println(err.Error())
		return
	}

	db := repository.Mysql()

	var exists bool
	err = db.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", form.Email).Find(&exists).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Select email error.")
		log.Println(err.Error())
		return
	}

	if exists {
		response.EmailAlreadyExists(c, status.AlreadyExistsEmail)
		log.Println("Already exist email.")
		return
	}

	hashedpw, err := convertPasswordtoHash(form.Password)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Convert password to hash error.")
		log.Println(err.Error())
		return
	}
	form.Password = hashedpw

	err = db.Create(&models.User{Email: form.Email, Password: form.Password}).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println("Insert account error.")
		log.Println(err.Error())
		return
	}

	response.SuccessfullySignUp(c, status.StatusOK)
}

func convertPasswordtoHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
