package point

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func List(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var pointList []response.Point

	mysqlClient := repository.Mysql()
	err := mysqlClient.Model(&models.Point{}).Where("uid = ? AND is_deleted = ?", uid, false).Scan(&pointList).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.PointList(c, pointList, status.StatusOK)
}
