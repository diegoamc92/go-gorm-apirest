package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=Dgo_88939981 dbname=gorm port=5432"
var DB *gorm.DB

func DBconnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connected")
	}
}
