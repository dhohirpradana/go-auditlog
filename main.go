package main

import (
	"github.com/joho/godotenv"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := ValidateEnvVars(); err != nil {
		log.Fatalf("Environment variable validation failed: %v", err)
	}

	router := gin.Default()

	// dynamic method
	router.Any("/:protocol/:url/*path", bridgeRequest)

	addr := ":8080"
	log.Printf("Starting server on %s...\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
