package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"http_code": http.StatusInternalServerError,
		"here_code": code,
		"message":   "Internal server error.",
	})
}
