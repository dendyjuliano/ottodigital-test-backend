package db

import (
	"database/sql"
	"fmt"
	"otto-test-go/internal/config"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var database *sql.DB

// Connect establishes a connection to the database
func Connect() error {
	return ConnectWithConfig(config.LoadConfig())
}

// ConnectWithConfig establishes a connection to the database using config
func ConnectWithConfig(cfg *config.Config) error {
	// MySQL connection string format: username:password@protocol(address)/dbname?param=value
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName,
		cfg.DbCharset, cfg.DbParseTime, cfg.DbLoc,
	)
	
	var err error
	database, err = sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	
	// Test the connection
	err = database.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return database
}