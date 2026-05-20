package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize Repository
	mongoURI := "mongodb://localhost:27017"
	dbName := "trading_db"

	repo, err := NewRepository(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 2. Initialize Service (Worker Pool & Cache)
	service := NewService(repo)

	// 3. Initialize Handler
	handler := NewHandler(service)

	// 4. Setup Router
	router := gin.Default()
	handler.SetupRoutes(router)

	// 5. Start Server
	log.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
