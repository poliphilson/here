package repository

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Mysql() *gorm.DB {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Redis() *redis.Client {
	host := os.Getenv("REDIS_DBHOST")
	port := os.Getenv("REDIS_DBPORT")
	pass := os.Getenv("REDIS_DBPASS")
	dbName, err := strconv.Atoi(os.Getenv("REDIS_DBNAME"))
	if err != nil {
		panic(err.Error())
	}

	db := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       dbName,
	})
	return db
}
