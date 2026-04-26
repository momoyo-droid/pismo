package repository

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountInterface interface {
	CreateAccount(ctx context.Context, account *model.Account) error
}

type AccountRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

func NewAccountRepository(db *gorm.DB, logger *zap.Logger) *AccountRepository {
	return &AccountRepository{Storage: db, Logger: logger}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *model.Account) error {
	r.Logger.Info("CreateAccount repository called")

	var entityAccount entity.Account

	if err := copier.Copy(&entityAccount, account); err != nil {
		r.Logger.Error("Failed to copy account data", zap.Error(err))
		return err
	}

	if err := r.Storage.Create(&entityAccount).Error; err != nil {
		r.Logger.Error("Failed to create account in database", zap.Error(err))
		return err
	}

	r.Logger.Info("Account created in database successfully")
	return nil
}
