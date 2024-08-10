package db

import (
	"fmt"
	"log"

	"20MIS0158_Backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(cfg *config.Config) {
	user := cfg.Database.User
	password := cfg.Database.Password
	dbName := cfg.Database.DBName
	host := cfg.Database.Host
	port := cfg.Database.Port

	if user == "" || password == "" || dbName == "" || host == "" || port == "" {
		log.Fatal("One or more required database configuration parameters are missing")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbName, host, port)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get raw database connection:", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Database connected")
}
