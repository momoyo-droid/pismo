package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	CreateTransaction(documentNumber string) error
}

type TransactionRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

func NewTransactionRepository(db *gorm.DB, logger *zap.Logger) *TransactionRepository {
	return &TransactionRepository{Storage: db, Logger: logger}
}
