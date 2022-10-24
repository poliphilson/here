package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/auth"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
)

func init() {
	db := repository.Connect()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Here{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Point{})
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", auth.SignUp)
	r.POST("/signin", auth.SignIn)
	r.POST("/signout", auth.SignOut)
	r.POST("/refresh", auth.RefreshAccessToken)

	authMiddle := r.Group("/")
	authMiddle.Use(auth.VerifyAccessToken)
	authMiddle.GET("/test", func(c *gin.Context) {
		data, _ := c.Get("email")
		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})
	r.Run()
}
