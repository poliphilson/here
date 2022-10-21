package sign

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"golang.org/x/crypto/bcrypt"
)

type signUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Up(c *gin.Context) {
	var form signUpForm
	err1 := c.BindJSON(&form)
	if err1 != nil {
		response.InternalServerError(c)
		log.Println("BindJson error.")
		return
	}

	db := repository.Connect()

	var exists bool
	err2 := db.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", form.Email).Find(&exists).Error
	if err2 != nil {
		response.InternalServerError(c)
		log.Println("Select email error.")
		return
	}

	if exists {
		response.EmailAlreadyExists(c)
		log.Println("Already exist email.")
		return
	}

	hashedpw, err3 := convertPasswordtoHash(form.Password)
	if err3 != nil {
		response.InternalServerError(c)
		log.Println("Convert password to hash error.")
		return
	}
	form.Password = hashedpw

	err4 := db.Create(&models.User{Email: form.Email, Password: form.Password}).Error
	if err4 != nil {
		response.InternalServerError(c)
		log.Println("Insert account error.")
		return
	}

	response.SuccessfullySignUp(c)
}

func convertPasswordtoHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
