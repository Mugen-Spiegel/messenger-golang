package models

import (
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	createDatabase()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	database, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{}, &Contact{}, &Task{})
	if err != nil {
		return
	}

	DB = database
}

func createDatabase() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	is_exist := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME"))
	result := DB.Exec(is_exist)
	fmt.Println(result)
}
