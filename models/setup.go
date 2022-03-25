package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "lalit:1233@tcp(dreamy_spence:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil { // incase we fail to connect to db, we print error on console
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&User{}) //migrate database schema using AutoMigrate , for each model we have to call this method
	database.AutoMigrate(&UpdateUser{})

	DB = database
}
