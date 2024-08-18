package config

import (
	"auth_service/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var MySQL *gorm.DB

func ConnectMySQL() {
	var username = os.Getenv("MYSQL_USER")
	var password = os.Getenv("MYSQL_PASSWORD")
	var host = os.Getenv("MYSQL_HOST")
	var port = os.Getenv("MYSQL_PORT")
	var db = os.Getenv("MYSQL_DB")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, db)
	log.Println(dsn)
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	MySQL.AutoMigrate(&model.User{}, &model.Session{})
}
