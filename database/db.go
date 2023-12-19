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
func ConnectToDatabase() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		return nil, err
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
		fmt.Println("Could not connect to the database")
		return nil, err
	}

	err = database.AutoMigrate(
		&models.Pet{},
		&models.Room{},
		&models.University{},
		&models.Professor{},
		&models.Student{})

	if err != nil {
		fmt.Println("Could not migrate the database")
		return nil, err
	}

	return database, nil
}
