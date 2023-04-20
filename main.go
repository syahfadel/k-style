package main

import (
	"fmt"
	"log"
	"unnispick/entities"
	"unnispick/routers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "3306"
	dbName   = "unnispick"
	db       *gorm.DB
	err      error
)

func init() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbPort, dbName)

	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database : ", err)
	}

	db.Debug().AutoMigrate(entities.Member{}, entities.Product{}, entities.ReviewProduct{}, entities.LikeReview{})
}
func main() {
	routers.StartService(db).Start(":4000")
}
