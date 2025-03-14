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
	r.POST("/users", models.CreateUser(db)) // ✅ Success
		// HTTP/1.1 201 Created
		// Content-Type: application/json; charset=utf-8
		// Date: Fri, 14 Mar 2025 03:02:47 GMT
		// Content-Length: 209
		// Connection: close
		//
		// {
		//   "id": 17,
		//   "password": "hashed_password_1235",
		//   "sap_user_code": "SAP012",
		//   "created_at": "2025-03-14T03:02:47.077073006Z",
		//   "updated_at": "0001-01-01T00:00:00Z",
		//   "deleted_at": {
		// 	"Time": "0001-01-01T00:00:00Z",
		// 	"Valid": false
		//   }
		// }	
	r.GET("/users", models.GetUsers(db)) // error
		// HTTP/1.1 500 Internal Server Error
		// Content-Type: application/json; charset=utf-8
		// Date: Thu, 13 Mar 2025 10:43:15 GMT
		// Content-Length: 148
		// Connection: close
		//
		// {
		// "error": "sql: Scan error on column index 4, name \"updated_at\": unsupported Scan, storing driver.Value type \u003cnil\u003e into type *time.Time"
		// }
	r.GET("/users/:id", models.GetUser(db)) // error
		// HTTP/1.1 500 Internal Server Error
		// Content-Type: application/json; charset=utf-8
		// Date: Thu, 13 Mar 2025 10:45:56 GMT
		// Content-Length: 148
		// Connection: close
		//
		// {
		// "error": "sql: Scan error on column index 4, name \"updated_at\": unsupported Scan, storing driver.Value type \u003cnil\u003e into type *time.Time"
		// }
	r.PUT("/users/:id", models.UpdateUser(db)) // ✅ Success
		// HTTP/1.1 200 OK
		// Content-Type: application/json; charset=utf-8
		// Date: Thu, 13 Mar 2025 10:39:33 GMT
		// Content-Length: 39
		// Connection: close
		//
		// {
		// "message": "User updated successfully"
		// }
	r.DELETE("/users/:id", models.DeleteUser(db)) // ✅ Success
		// {
		// 	"message": "User deleted successfully"
		// }

	// Start server
	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
