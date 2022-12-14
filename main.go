package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/auth"
	"github.com/poliphilson/here/config"
	"github.com/poliphilson/here/here"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/point"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/trash"
	"github.com/poliphilson/here/user"
)

func init() {
	config.DB = repository.Mysql()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	err = config.DB.AutoMigrate(&models.Here{})
	if err != nil {
		panic(err.Error())
	}
	err = config.DB.AutoMigrate(&models.Point{})
	if err != nil {
		panic(err.Error())
	}
	err = config.DB.AutoMigrate(&models.HereImage{})
	if err != nil {
		panic(err.Error())
	}
	err = config.DB.AutoMigrate(&models.HereVideo{})
	if err != nil {
		panic(err.Error())
	}
	err = config.DB.AutoMigrate(&models.HereAddress{})
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

	authMiddle := r.Group("/")
	authMiddle.Use(auth.VerifyAccessToken)
	authMiddle.Static("/image", hereImagePath)
	authMiddle.Static("/video", hereVideoPath)

	authMiddle.PATCH("/user", user.Edit)
	authMiddle.GET("/user", user.Detail)

	authMiddle.GET("/here", here.List)
	authMiddle.POST("/here", here.Upload)
	authMiddle.GET("/here/:hid", here.Detail)
	authMiddle.PATCH("/here/:hid", here.Edit)
	authMiddle.DELETE("/here/:hid", here.Delete)

	authMiddle.GET("/point", point.List)
	authMiddle.POST("/point", point.Create)
	authMiddle.PATCH("/point/:pid", point.Edit)
	authMiddle.DELETE("/point/:pid", point.Delete)

	authMiddle.GET("/trash/here", trash.HereList)
	authMiddle.PATCH("/trash/here/:hid", trash.HereRecovery)
	authMiddle.DELETE("/trash/here/:hid", trash.HereDelete)
	authMiddle.GET("/trash/point", trash.PointList)
	authMiddle.PATCH("/trash/point/:pid", trash.PointRecovery)
	authMiddle.DELETE("/trash/point/:pid", trash.PointDelete)
	r.Run(":8080")
}
