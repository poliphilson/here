package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
)

func HeresOnMap(c *gin.Context, heres []models.Here, code int) {
	c.JSON(http.StatusOK, gin.H{
		"http_code": http.StatusOK,
		"here_code": code,
		"data":      heres,
	})
}
