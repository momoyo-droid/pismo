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
	if err := db.AutoMigrate(&entity.Account{}, &entity.Transaction{}, &entity.Operation{}); err != nil {
		return nil, fmt.Errorf("auto-migrate error: %w", err)
	}

	if err := createOperations(db); err != nil {
		return nil, fmt.Errorf("failed to create operations: %w", err)
	}

	return db, nil
}

// createOperations inserts predefined financial operations into the database if they do not already exist.
// It defines a list of operations (Purchase, Withdrawal, Installment Purchase, Payment) and uses
// the FirstOrCreate method to ensure that each operation is only created once in the database.
func createOperations(db *gorm.DB) error {
	ops := []entity.Operation{
		{ID: 1, Description: "Purchase"},
		{ID: 2, Description: "Withdrawal"},
		{ID: 3, Description: "Installmente Purchase"},
		{ID: 4, Description: "Payment"},
	}

	for _, op := range ops {
		if err := db.FirstOrCreate(&op, entity.Operation{ID: op.ID}).Error; err != nil {
			return fmt.Errorf("failed to create operation %d: %w", op.ID, err)
		}
	}
	return nil
}
