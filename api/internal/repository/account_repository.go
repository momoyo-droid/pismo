package repository

import (
	"context"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AccountRepository provides methods to interact with the accounts in the database.
// It contains a reference to the GORM database connection and a Logger for logging.
type AccountRepository struct {
	Storage *gorm.DB
	Logger  *zap.Logger
}

// NewAccountRepository creates a new instance of AccountRepository with the provided database connection and logger.
func NewAccountRepository(db *gorm.DB, logger *zap.Logger) *AccountRepository {
	return &AccountRepository{Storage: db, Logger: logger}
}

// CreateAccount creates a new account in the database using the provided account model.
// It logs the process and returns the created account or an error if the operation fails.
func (r *AccountRepository) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	r.Logger.Info("CreateAccount repository called")

	entityAccount := entity.Account{
		DocumentNumber: account.DocumentNumber,
	}

	if err := r.Storage.Create(&entityAccount).Error; err != nil {
		r.Logger.Error("failed to create account in database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("account created in database successfully")
	return &model.Account{
		ID:             entityAccount.ID,
		DocumentNumber: entityAccount.DocumentNumber,
	}, nil
}

// GetAccountByID retrieves an account from the database by its ID.
// It logs the process and returns the account or an error if the operation fails.
func (r *AccountRepository) GetAccountByID(ctx context.Context, accountID uint64) (*model.Account, error) {
	r.Logger.Info("GetAccountByID repository called")

	var entityAccount entity.Account
	if err := r.Storage.First(&entityAccount, "id = ?", accountID).Error; err != nil {
		r.Logger.Error("failed to fetch account from database", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("account fetched from database successfully")
	return &model.Account{
		ID:             entityAccount.ID,
		DocumentNumber: entityAccount.DocumentNumber,
	}, nil
}
