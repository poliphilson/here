package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func SignOut(c *gin.Context) {
	c.SetCookie("access-token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh-token", "", -1, "/", "localhost", false, true)

	response.SeccessfullySignOut(c, status.StatusOK)
}
