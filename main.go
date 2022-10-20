package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic("Init database error.")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Here{})
	db.AutoMigrate(&models.Point{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := "root:rbdhks12@tcp(34.64.219.30:3306)/how_about_here?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
