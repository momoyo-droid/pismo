package repository

import (
	"context"
	"time"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type TransactionRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

func NewTransactionRepository(db *gorm.DB, logger *zap.Logger) *TransactionRepository {
	return &TransactionRepository{Storage: db, Logger: logger}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	r.Logger.Info("Creating transaction in repository")

	var entityTransaction entity.Transaction

	entityTransaction = entity.Transaction{
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
		CreationDate:    time.Now(),
	}

	if err := r.Storage.Create(&entityTransaction).Error; err != nil {
		r.Logger.Error("Failed to create transaction in database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("Transaction created in database successfully")
	return &model.Transaction{
		TransactionID:   entityTransaction.TransactionID,
		AccountID:       entityTransaction.AccountID,
		OperationTypeID: entityTransaction.OperationTypeID,
		Amount:          entityTransaction.Amount,
	}, nil
}
