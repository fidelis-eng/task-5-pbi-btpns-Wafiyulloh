package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models "photoapi.com/ppapi/models"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "pgsqlpass"
	dbname   = "photoapi"
)

func ConnectDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})

	DB = db

}
