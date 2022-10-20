package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliphilson/here/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDb()
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func initDb() *gorm.DB {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
