package database

import (
	"fmt"
	"github.com/Strke12i/go_docker/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// ConnectToDatabase function connects to the database
// and returns the database object
func ConnectToDatabase() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file!")
	}

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	fmt.Println(connectString)

	database, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(
		&models.Pet{},
		&models.Room{},
		&models.University{})

	if err != nil {
		panic("Failed to migrate database!")
	}

	return database
}
