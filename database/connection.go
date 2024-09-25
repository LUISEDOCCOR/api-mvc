package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("dsn"),
	}))
	if err != nil {
		fmt.Println("Error in the databse 🚨")
	} else {
		fmt.Println("The database is ready to use 🚀")
	}
}
