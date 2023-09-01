package config

import (
	"fmt"

	"github.com/tahamazari/echo-framework/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "postgres"
	password := "1234"
	dbName := "echo_bookstore"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Karachi", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	err := database.AutoMigrate(&models.Book{}) // Add other models here
	if err != nil {
		fmt.Println(err)
	}
}

func DB() *gorm.DB {
	return database
}
