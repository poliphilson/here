package point

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

type GetPoint struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Description string  `json:"description"`
}

func Create(c *gin.Context) {
	var getPoint GetPoint
	err := c.Bind(&getPoint)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var pointForm models.Point
	pointForm.Uid = uid.(int)
	pointForm.Location.X = getPoint.X
	pointForm.Location.Y = getPoint.Y
	pointForm.Description = getPoint.Description

	err = createPoint(pointForm)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.CreateOk(c, status.StatusOK)
}

func createPoint(point models.Point) error {
	err := config.DB.Create(&point).Error
	if err != nil {
		return err
	}
	return nil
}
