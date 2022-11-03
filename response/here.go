package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeresOnMap(c *gin.Context, heres []SimpleHere, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"data":      heres,
	})
}

func DetailHereOnMap(c *gin.Context, here DetailHere, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"data":      here,
	})
}
