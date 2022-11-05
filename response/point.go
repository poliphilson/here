package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PointList(c *gin.Context, points []Point, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"data":      points,
	})
}
