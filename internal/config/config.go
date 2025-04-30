package config

import (
	"os"
)

type Config struct {
    ServerPort string
    DbHost     string
    DbPort     string
    DbUser     string
    DbPassword string
    DbName     string
    DbCharset  string
    DbParseTime string
    DbLoc       string
}

func LoadConfig() *Config {
    return &Config{
        ServerPort:  getEnv("SERVER_PORT", "8080"),
        DbHost:      getEnv("DB_HOST", "localhost"),
        DbPort:      getEnv("DB_PORT", "3306"),        
        DbUser:      getEnv("DB_USER", "root"),       
        DbPassword:  getEnv("DB_PASSWORD", ""),        
        DbName:      getEnv("DB_NAME", "ottotest"),
        DbCharset:   getEnv("DB_CHARSET", "utf8mb4"),
        DbParseTime: getEnv("DB_PARSE_TIME", "True"),
        DbLoc:       getEnv("DB_LOC", "Local"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}