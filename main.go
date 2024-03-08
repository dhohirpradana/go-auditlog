package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// dynamic method
	router.Any("/:protocol/:url/*path", bridgeRequest)

	addr := "localhost:8080"
	log.Printf("Starting server on %s...\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
