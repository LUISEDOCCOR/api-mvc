package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	var err error

	DB, err = gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in the databse 🚨")
	} else {
		fmt.Println("The database is ready to use 🚀")
	}
}
