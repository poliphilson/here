package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInformation(c *gin.Context, user User, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"data":      user,
	})
}
