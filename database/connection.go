package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=postgres dbname=gorm port=5432"
var DB *gorm.DB
var err error

func DBConnection() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("Connecting to database")
	}
}
