package repository

import (
	"context"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

func NewAccountRepository(db *gorm.DB, logger *zap.Logger) *AccountRepository {
	return &AccountRepository{Storage: db, Logger: logger}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	r.Logger.Info("CreateAccount repository called")

	var entityAccount entity.Account

	entityAccount = entity.Account{
		DocumentNumber: account.DocumentNumber,
	}

	if err := r.Storage.Create(&entityAccount).Error; err != nil {
		r.Logger.Error("Failed to create account in database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("Account created in database successfully")
	return &model.Account{
		ID:             entityAccount.ID,
		DocumentNumber: entityAccount.DocumentNumber,
	}, nil
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, accountID uint64) (*model.Account, error) {
	r.Logger.Info("GetAccountByID repository called")

	var entityAccount entity.Account
	if err := r.Storage.First(&entityAccount, "id = ?", accountID).Error; err != nil {
		r.Logger.Error("Failed to fetch account from database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("Account fetched from database successfully")
	return &model.Account{
		ID:             entityAccount.ID,
		DocumentNumber: entityAccount.DocumentNumber,
	}, nil
}
