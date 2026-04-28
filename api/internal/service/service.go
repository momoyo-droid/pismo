package service

import (
	"context"
	"errors"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// AccountRepository provides methods to interact with the accounts in the database.
// It contains a reference to the GORM database connection and a Logger for logging.
//
//go:generate moq -out mocks/account_repository_mock.go -pkg mocks . AccountRepository
type AccountRepository interface {
	CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error)
	GetAccountByID(ctx context.Context, accountID uint64) (*model.Account, error)
}

// TransactionRepository provides methods to interact with transactions in the database.
// It contains a reference to the GORM database connection and a Logger for logging.
//
//go:generate moq -out mocks/transaction_repository_mock.go -pkg mocks . TransactionRepository
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

// Service is the main service layer of the application.
// It contains references to the account and transaction repositories, as well as a logger for logging operations.
type Service struct {
	AccountRepo     AccountRepository
	TransactionRepo TransactionRepository
	Logger          *zap.Logger
}

// NewService creates a new instance of Service with the provided account repository, transaction repository, and logger.
func NewService(accountRepo AccountRepository, transactionRepo TransactionRepository, logger *zap.Logger) *Service {
	return &Service{
		AccountRepo:     accountRepo,
		TransactionRepo: transactionRepo,
		Logger:          logger,
	}
}

// CreateAccount handles the creation of a new account.
// It validates the account data and then creates the account in the repository.
// It returns the created account or an error if any validation fails.
func (s *Service) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	s.Logger.Info("CreateAccount service called")

	if account.DocumentNumber == "" {
		s.Logger.Fatal("document number is empty")
		return nil, errors.New("document number is required")
	}

	s.Logger.Info("validation passed, creating account in repository")

	return s.AccountRepo.CreateAccount(ctx, account)
}

// GetAccountByID retrieves an account by its ID.
// It validates the account ID and then fetches the account from the repository.
// It returns the account or an error if the account ID is invalid or if the repository operation fails.
func (s *Service) GetAccountByID(ctx context.Context, accountID uint64) (*model.Account, error) {
	s.Logger.Info("GetAccountByID service called")

	if accountID == 0 {
		s.Logger.Error("invalid account ID")
		return nil, errors.New("account ID is required")
	}

	s.Logger.Info("validation passed, fetching account from repository")

	return s.AccountRepo.GetAccountByID(ctx, accountID)
}

// CreateTransaction handles the creation of a new transaction.
// It validates the transaction data, checks if the associated account exists,
// and then creates the transaction in the repository.
// It returns the created transaction or an error if any validation fails.
func (s *Service) CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	s.Logger.Info("CreateTransaction service called")

	if transaction.AccountID == 0 {
		s.Logger.Error("invalid account ID")
		return nil, errors.New("invalid account ID")
	}

	account, err := s.AccountRepo.GetAccountByID(ctx, transaction.AccountID)

	if err != nil {
		s.Logger.Error("failed to fetch account for transaction", zap.Error(err))
		return nil, errors.New("account not found")
	}

	if account == nil {
		s.Logger.Error("account not found for transaction")
		return nil, errors.New("account not found")
	}

	op := model.OperationType(transaction.OperationTypeID)
	if !op.IsValid() {
		s.Logger.Error("invalid operation type ID")
		return nil, errors.New("invalid operation type ID")
	}

	if transaction.Amount.LessThanOrEqual(decimal.Zero) {
		s.Logger.Error("invalid transaction amount")
		return nil, errors.New("transaction amount must be greater than zero")
	}

	if op.IsDebit() {
		transaction.Amount = transaction.Amount.Neg()
	}

	s.Logger.Info("validation passed, creating transaction in repository")

	return s.TransactionRepo.CreateTransaction(ctx, transaction)
}
