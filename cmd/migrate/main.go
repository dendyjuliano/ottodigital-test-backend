package main

import (
	"flag"
	"log"
	"otto-test-go/internal/config"
	"otto-test-go/internal/db"
)

func main() {
    var migrationType string
    flag.StringVar(&migrationType, "type", "up", "Migration type: up or down")
    flag.Parse()
    
    cfg := config.LoadConfig()
    
    switch migrationType {
    case "up":
        if err := db.RunMigrations(cfg); err != nil {
            log.Fatalf("Failed to run migrations: %v", err)
        }
        log.Println("Migrations completed successfully")
    case "down":
        if err := db.RollbackMigrations(cfg); err != nil {
            log.Fatalf("Failed to rollback migrations: %v", err)
        }
        log.Println("Rollback completed successfully")
    default:
        log.Fatalf("Unknown migration type: %s", migrationType)
    }
}