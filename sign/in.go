package sign

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"golang.org/x/crypto/bcrypt"
)

type signInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func In(c *gin.Context) {
	var form signInForm
	err1 := c.BindJSON(&form)
	if err1 != nil {
		response.InternalServerError(c)
		log.Println("BindJson error.")
		return
	}

	db := repository.Connect()

	var user models.User
	err2 := db.Where("email = ?", form.Email).Find(&user).Error
	if err2 != nil {
		response.InternalServerError(c)
		log.Println("Select email error.")
		return
	}

	if user.Email == "" {
		response.EmailDoesNotExist(c)
		log.Println("Not registered email.")
		return
	}

	result := checkPassword(user.Password, form.Password)
	if !result {
		response.FailedSignIn(c)
		return
	}

	response.SuccessfullySignIn(c)
}

func checkPassword(hashVal, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
