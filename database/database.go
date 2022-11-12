package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func Init() {
	var err error

	dsn := "host=localhost port=5431 user=postgres dbname=delivery password=example"

	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Database Connected")
	}
}

// GetDB ...
func GetDB() *gorm.DB {
	return DBConnection
}

// CloseDB ...
func CloseDB() {
	DBConnection, _ := DBConnection.DB()
	DBConnection.Close()
}
