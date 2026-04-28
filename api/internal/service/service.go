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

func (s *Service) GetAccountByID(ctx context.Context, accountID uint64) (*model.Account, error) {
	s.Logger.Info("GetAccountByID service called")

	if accountID == 0 {
		s.Logger.Error("Invalid account ID")
		return nil, errors.New("account ID is required")
	}

	s.Logger.Info("Validation passed, fetching account from repository")

	return s.AccountRepo.GetAccountByID(ctx, accountID)
}

func (s *Service) CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	s.Logger.Info("CreateTransaction service called")

	if transaction.AccountID == 0 {
		s.Logger.Error("Invalid account ID")
		return nil, errors.New("Invalid account ID")
	}

	account, err := s.AccountRepo.GetAccountByID(ctx, transaction.AccountID)

	if err != nil {
		s.Logger.Error("Failed to fetch account for transaction", zap.Error(err))
		return nil, errors.New("Account not found")
	}

	if account == nil {
		s.Logger.Error("Account not found for transaction")
		return nil, errors.New("Account not found")
	}

	op := model.OperationType(transaction.OperationTypeID)
	if !op.IsValid() {
		s.Logger.Error("Invalid operation type ID")
		return nil, errors.New("Invalid operation type ID")
	}

	transaction.Amount = transaction.Amount.Mul(op.IsDebitOrCredit())

	s.Logger.Info("Validation passed, creating transaction in repository")

	return s.TransactionRepo.CreateTransaction(ctx, transaction)
}
