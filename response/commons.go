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

func BadRequset(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"http_code": http.StatusBadRequest,
		"here_code": code,
		"message":   "Bad request.",
	})
}

func Ok(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"message":   "Ok.",
	})
}
