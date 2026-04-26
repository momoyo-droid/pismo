package postgres

import (
	"fmt"

	"github.com/momoyo-droid/pismo/api/internal/config"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseConnection establishes a new database connection using the provided configuration.
// It returns a gorm.DB instance and an error if the connection fails.
// The function also performs auto-migration for the Seller and Owner models to
// ensure the database schema is up to date.
func NewDatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("create database connection error: %w", err)
	}
	// Create tables if they do not exist
	if err := db.AutoMigrate(&entity.Account{}, &entity.Transaction{}); err != nil {
		return nil, fmt.Errorf("auto-migrate error: %w", err)
	}

	return db, nil
}
