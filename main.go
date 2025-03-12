package main

import (
	"log"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"try-go/database"
	"try-go/models"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve environment variables
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")

	// Initialize database connection
	db, err := database.InitDB(dbUser, dbPass, dbHost, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize router
	r := gin.Default()

	// Define routes
	r.POST("/users", models.CreateUser(db))
	r.GET("/users", models.GetUsers(db))
	r.GET("/users/:id", models.GetUser(db))
	r.PUT("/users/:id", models.UpdateUser(db))
	r.DELETE("/users/:id", models.DeleteUser(db))

	// Start server
	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
