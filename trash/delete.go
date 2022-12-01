package trash

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
	"gorm.io/gorm"
)

func HereDelete(c *gin.Context) {
	temp := c.Param("hid")
	hid, err := strconv.Atoi(temp)
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

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("hid = ?", hid).Delete(&models.HereImage{}).Error
		if err != nil {
			return err
		}

		err = tx.Where("hid = ?", hid).Delete(&models.HereVideo{}).Error
		if err != nil {
			return err
		}

		result := tx.Where("uid = ? AND hid = ? AND is_deleted = ?", uid, hid, true).Delete(&models.Here{})
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("rollback deleted image and video")
		}

		return nil
	})
	if err != nil {
		if strings.HasPrefix(err.Error(), "rollback") {
			response.BadRequset(c, status.StatusOK)
			return
		} else {
			response.InternalServerError(c, status.InternalError)
			log.Println(err.Error())
			return
		}
	}

	response.Ok(c, status.StatusOK)
}

func PointDelete(c *gin.Context) {
	temp := c.Param("pid")
	pid, err := strconv.Atoi(temp)
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

	result := config.DB.Where("pid = ? AND uid = ? AND is_deleted = ?", pid, uid, true).Delete(&models.Point{})
	if result.Error != nil {
		response.InternalServerError(c, status.InternalError)
		log.Println(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		response.BadRequset(c, status.StatusOK)
		return
	}

	response.Ok(c, status.StatusOK)
}
