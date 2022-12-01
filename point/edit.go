package point

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

type EditPoint struct {
	Description string `json:"description"`
}

func Edit(c *gin.Context) {
	temp := c.Param("pid")
	pid, err := strconv.Atoi(temp)
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	var editPoint EditPoint
	err = c.Bind(&editPoint)
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

	err = config.DB.Model(&models.Point{}).Where("pid = ? AND uid = ?", pid, uid).Update("description", editPoint.Description).Error
	if err != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(err.Error())
		return
	}

	response.Ok(c, status.StatusOK)
}
