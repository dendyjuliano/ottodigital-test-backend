package db

import (
	"database/sql"
	"fmt"
	"log"
	"otto-test-go/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations performs database migrations
func RunMigrations(cfg *config.Config) error {
    log.Println("Running database migrations...")

    // Create DB connection for migrations
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?multiStatements=true",
        cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error opening database connection: %v", err)
    }
    defer db.Close()

    // Create database if it doesn't exist
    _, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DbName))
    if err != nil {
        return fmt.Errorf("error creating database: %v", err)
    }

    // Connect to the specific database for migrations
    dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&charset=%s&parseTime=%s&loc=%s",
        cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName,
        cfg.DbCharset, cfg.DbParseTime, cfg.DbLoc)
    
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error opening database: %v", err)
    }

    driver, err := mysql.WithInstance(db, &mysql.Config{})
    if err != nil {
        return fmt.Errorf("could not start sql migration: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations", // path to migration files
        "mysql",             // database type
        driver,              // database instance
    )
    if err != nil {
        return fmt.Errorf("error creating migration instance: %v", err)
    }

    // Run migrations
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("an error occurred while running migrations: %v", err)
    }

    log.Println("Migrations completed successfully")
    return nil
}

// RollbackMigrations rolls back all migrations
func RollbackMigrations(cfg *config.Config) error {
    log.Println("Rolling back database migrations...")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&charset=%s&parseTime=%s&loc=%s",
        cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName,
        cfg.DbCharset, cfg.DbParseTime, cfg.DbLoc)
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error opening database: %v", err)
    }

    driver, err := mysql.WithInstance(db, &mysql.Config{})
    if err != nil {
        return fmt.Errorf("could not start sql migration: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations", // path to migration files
        "mysql",             // database type
        driver,              // database instance
    )
    if err != nil {
        return fmt.Errorf("migration failed: %v", err)
    }

    // Run migrations down
    if err := m.Down(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("an error occurred while rolling back migrations: %v", err)
    }

    log.Println("Rollback completed successfully")
    return nil
}