// create connection with database
//gorm.Open method creates new connection with databse
//print error msg if failed to connect
//migrate all models using AutoMigrate
//assign database to package level variable

package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db") //using SQLite database & store data inside test.db

	if err != nil { // incase we fail to connect to db, we print error on console
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&User{}) //migrate database schema using AutoMigrate , for each model we have to call this method

	DB = database
}
