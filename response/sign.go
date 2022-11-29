package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmailAlreadyExists(c *gin.Context, code int) {
	c.JSON(http.StatusConflict, gin.H{
		"http_code": http.StatusConflict,
		"here_code": code,
		"message":   "This email already exists.",
	})
}

func SuccessfullySignUp(c *gin.Context, code int) {
	c.JSON(http.StatusCreated, gin.H{
		"http_code": http.StatusCreated,
		"here_code": code,
		"message":   "Successfully sign up.",
	})
}

func EmailDoesNotExist(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"http_code": http.StatusUnauthorized,
		"here_code": code,
		"message":   "This email is not registered.",
	})
}

func SuccessfullySignIn(c *gin.Context, user SignIn, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"message":   "Successfully sign in.",
		"data":      user,
	})
}

func FailedSignIn(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"http_code": http.StatusUnauthorized,
		"here_code": code,
		"message":   "Failed to sign in.",
	})
}

func SeccessfullySignOut(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"message":   "Successfully sign out.",
	})
}
