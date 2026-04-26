package service

import (
	"context"
	"errors"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	AccountRepo     *repository.AccountRepository
	TransactionRepo *repository.TransactionRepository
	Logger          *zap.Logger
}

func NewService(accountRepo *repository.AccountRepository, transactionRepo *repository.TransactionRepository, logger *zap.Logger) *Service {
	return &Service{
		AccountRepo:     accountRepo,
		TransactionRepo: transactionRepo,
		Logger:          logger,
	}
}

func (s *Service) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	s.Logger.Info("CreateAccount service called")

	if account.DocumentNumber == "" {
		s.Logger.Fatal("Document number is empty")
		return nil, errors.New("document number is required")
	}

	s.Logger.Info("Validation passed, creating account in repository")

	return s.AccountRepo.CreateAccount(ctx, account)
}

func (s *Service) GetAccountByID(ctx context.Context, accountID string) (*model.Account, error) {
	s.Logger.Info("GetAccountByID service called", zap.String("account_id", accountID))

	if accountID == "" {
		s.Logger.Error("Account ID is empty")
		return nil, errors.New("account ID is required")
	}

	s.Logger.Info("Validation passed, fetching account from repository")

	return s.AccountRepo.GetAccountByID(ctx, accountID)
}
