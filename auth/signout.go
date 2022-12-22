package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func SignOut(c *gin.Context) {
	/*cookie, _ := c.Request.Cookie("refresh_token")
	rToken := cookie.Value

	DeleteRefreshTokenFromRedis(rToken)*/

	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	response.SeccessfullySignOut(c, status.StatusOK)
}
