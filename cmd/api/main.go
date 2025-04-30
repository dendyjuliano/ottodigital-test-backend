package main

import (
	"log"
	"net/http"
	"otto-test-go/internal/api"
	"otto-test-go/internal/config"
	"otto-test-go/internal/db"
)

func main() {
    cfg := config.LoadConfig()
    
    // Run migrations on startup
    if err := db.RunMigrations(cfg); err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }
    
    router := api.SetupRouter()

    log.Printf("Starting server on port %s...", cfg.ServerPort)
    if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}