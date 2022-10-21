package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmailAlreadyExists(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{
		"code":    http.StatusConflict,
		"message": "This email already exists.",
	})
}

func SuccessfullySignUp(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Successfully sign up.",
	})
}

func EmailDoesNotExist(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "This email is not registered.",
	})
}

func SuccessfullySignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully sign in.",
	})
}

func FailedSignIn(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "Failed to sign in.",
	})
}
