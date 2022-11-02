package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/auth"
	"github.com/poliphilson/here/here"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
)

func init() {
	mysqlClient := repository.Mysql()
	err := mysqlClient.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	err = mysqlClient.AutoMigrate(&models.Here{})
	if err != nil {
		panic(err.Error())
	}
	err = mysqlClient.AutoMigrate(&models.Point{})
	if err != nil {
		panic(err.Error())
	}
	err = mysqlClient.AutoMigrate(&models.HereImage{})
	if err != nil {
		panic(err.Error())
	}
	err = mysqlClient.AutoMigrate(&models.HereVideo{})
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	hereImagePath := os.Getenv("HERE_IMAGE_PATH")
	hereVideoPath := os.Getenv("HERE_VIDEO_PATH")

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", auth.SignUp)
	r.POST("/signin", auth.SignIn)
	r.POST("/signout", auth.SignOut)
	r.POST("/refresh", auth.RefreshAccessToken)

	r.Static("/image", hereImagePath)
	r.Static("/video", hereVideoPath)

	authMiddle := r.Group("/")
	authMiddle.Use(auth.VerifyAccessToken)
	authMiddle.GET("/test", func(c *gin.Context) {
		data, _ := c.Get("email")
		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})
	authMiddle.POST("/upload", here.Upload)
	authMiddle.GET("/list", here.List)
	r.Run()
}
