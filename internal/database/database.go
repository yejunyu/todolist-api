package database

import (
	"fmt"
	"log"
	"todolist-api/internal/models"
	"todolist-api/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error
	cfg := config.Cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}
	log.Println("Database migrated successfully")
	return db, nil
}
