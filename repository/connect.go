package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	/*
		user := os.Getenv("DBUSER")
		pass := os.Getenv("DBPASS")
		host := os.Getenv("DBHOST")
		port := os.Getenv("DBPORT")
		dbName := os.Getenv("DBNAME")

		dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	*/

	dsn := "root:rbdhks12@tcp(34.64.219.30:3306)/how_about_here?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
