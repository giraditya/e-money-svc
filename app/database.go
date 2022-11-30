package app

import (
	"emoney-service/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if DB == nil {
		DBDriver := os.Getenv("DB_DRIVER")
		DBHost := os.Getenv("DB_HOST")
		DBUser := os.Getenv("DB_USER")
		DBPassword := os.Getenv("DB_PASSWORD")
		DBName := os.Getenv("DB_NAME")
		DBPort := os.Getenv("DB_PORT")
		DBUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

		DB, err = gorm.Open(mysql.Open(DBUrl))

		if err != nil {
			fmt.Println("Cannot connect to database ", DBDriver)
			log.Fatal("connection error:", err)
		} else {
			fmt.Println("We are connected to the database ", DBDriver)
		}

		DB.AutoMigrate(&models.User{}, &models.Balance{}, &models.BalanceHistory{}, &models.Transaction{})
	}
	return DB
}
