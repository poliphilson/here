package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/sign"
)

func main() {
	initRepository()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", sign.Up)
	r.POST("/signin", sign.In)
	r.Run()
}

func initRepository() {
	db := repository.Connect()
	err1 := db.AutoMigrate(&models.User{})
	if err1 != nil {
		panic(err1.Error())
	}
	err2 := db.AutoMigrate(&models.Here{})
	if err2 != nil {
		panic(err2.Error())
	}
	err3 := db.AutoMigrate(&models.Point{})
	if err3 != nil {
		panic(err3.Error())
	}
}
