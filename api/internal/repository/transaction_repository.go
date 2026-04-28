package repository

import (
	"context"
	"time"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TransactionRepository provides methods to interact with transactions in the database.
// It contains a reference to the GORM database connection and a Logger for logging.
type TransactionRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

// NewTransactionRepository creates a new instance of TransactionRepository with the provided database connection and logger.
func NewTransactionRepository(db *gorm.DB, logger *zap.Logger) *TransactionRepository {
	return &TransactionRepository{Storage: db, Logger: logger}
}

// CreateTransaction creates a new transaction in the database using the provided transaction model.
// It logs the process and returns the created transaction or an error if the operation fails.
func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	r.Logger.Info("creating transaction in repository")

	entityTransaction := entity.Transaction{
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
		CreationDate:    time.Now(),
	}

	if err := r.Storage.Create(&entityTransaction).Error; err != nil {
		r.Logger.Error("failed to create transaction in database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("transaction created in database successfully")
	return &model.Transaction{
		TransactionID:   entityTransaction.TransactionID,
		AccountID:       entityTransaction.AccountID,
		OperationTypeID: entityTransaction.OperationTypeID,
		Amount:          entityTransaction.Amount,
	}, nil
}
