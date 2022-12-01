package trash

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func PointList(c *gin.Context) {
	query1 := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(query1)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
	}

	query2 := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(query2)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
	}

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var pointList []response.Point

	err = config.DB.
		Model(&models.Point{}).Where("uid = ? AND is_deleted = ?", uid, true).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&pointList).
		Error

	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.PointList(c, pointList, status.StatusOK)
}

func HereList(c *gin.Context) {
	query1 := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(query1)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
	}

	query2 := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(query2)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
	}

	uid, exists := c.Get("uid")
	if !exists {
		response.InternalServerError(c, status.FailedSignIn)
		return
	}

	var simpleHereList []response.SimpleHere

	err = config.DB.
		Model(&models.Here{}).
		Where("uid = ? AND is_deleted = ?", uid, true).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&simpleHereList).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.HeresOnMap(c, simpleHereList, status.StatusOK)
}
