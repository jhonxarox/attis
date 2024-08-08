package main

import (
	"backend-assignment/models"
	"backend-assignment/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate will create or update the database schema
	err = db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func main() {
	r := gin.Default()
	initDB()

	// Initialize routes
	routes.InitializeRoutes(r, db)

	r.Run(":8080")
}
