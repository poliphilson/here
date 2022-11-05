package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/auth"
	"github.com/poliphilson/here/here"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/point"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/user"
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
	authMiddle.PATCH("/user", user.Update)
	authMiddle.POST("/here", here.Upload)
	authMiddle.GET("/here", here.List)
	authMiddle.DELETE("/here/:hid", here.Delete)
	authMiddle.GET("/here/:hid", here.Detail)
	authMiddle.GET("/point", point.List)
	authMiddle.POST("/point", point.Create)
	authMiddle.PATCH("/point/:pid", point.Edit)
	authMiddle.DELETE("/point/:pid", point.Delete)
	r.Run()
}
