package main

import (
	"fmt"
	"log"
	"otto-test-go/internal/api"
)

func main() {
    router := api.SetupRouter() // Using Gin
    
    // Start server
    fmt.Println("Server running on port 8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}