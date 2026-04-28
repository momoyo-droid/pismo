package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/service"
	"github.com/momoyo-droid/pismo/api/internal/service/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestService_CreateAccount_Success(t *testing.T) {
	ctx := context.Background()

	account := &model.Account{
		DocumentNumber: "12345678900",
	}

	want := &model.Account{
		ID:             1,
		DocumentNumber: "12345678900",
	}

	repo := &mocks.AccountRepositoryMock{
		CreateAccountFunc: func(ctx context.Context, acc *model.Account) (*model.Account, error) {
			assert.Equal(t, account.DocumentNumber, acc.DocumentNumber)
			return want, nil
		},
	}

	logger, _ := zap.NewDevelopment()
	service := service.NewService(repo, nil, logger)

	response, err := service.CreateAccount(ctx, account)

	assert.Equal(t, want, response)
	assert.NoError(t, err)
}

func TestService_GetAccountByID_Success(t *testing.T) {
	ctx := context.Background()

	accountID := uint64(1)

	want := &model.Account{
		ID:             accountID,
		DocumentNumber: "12345678900",
	}

	repo := &mocks.AccountRepositoryMock{
		GetAccountByIDFunc: func(ctx context.Context, id uint64) (*model.Account, error) {
			assert.Equal(t, accountID, id)
			return want, nil
		},
	}

	logger, _ := zap.NewDevelopment()
	service := service.NewService(repo, nil, logger)

	response, err := service.GetAccountByID(ctx, accountID)

	assert.Equal(t, want, response)
	assert.NoError(t, err)
}

func TestService_CreateTransaction_Success(t *testing.T) {
	ctx := context.Background()

	transaction := &model.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          decimal.NewFromFloat(123.45),
	}

	expectedAmount := decimal.NewFromFloat(-123.45)

	want := &model.Transaction{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          expectedAmount,
	}

	accountRepo := &mocks.AccountRepositoryMock{
		GetAccountByIDFunc: func(ctx context.Context, id uint64) (*model.Account, error) {
			assert.Equal(t, transaction.AccountID, id)
			return &model.Account{
				ID:             transaction.AccountID,
				DocumentNumber: "12345678900",
			}, nil
		},
	}

	transactionRepo := &mocks.TransactionRepositoryMock{
		CreateTransactionFunc: func(ctx context.Context, tx *model.Transaction) (*model.Transaction, error) {
			assert.Equal(t, want.AccountID, tx.AccountID)
			assert.Equal(t, want.OperationTypeID, tx.OperationTypeID)
			assert.True(t, expectedAmount.Equal(tx.Amount))
			return want, nil
		},
	}

	logger, _ := zap.NewDevelopment()
	service := service.NewService(accountRepo, transactionRepo, logger)

	response, err := service.CreateTransaction(ctx, transaction)

	assert.Equal(t, want, response)
	assert.NoError(t, err)

}

func TestService_CreateTransaction_Failure(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		transaction   *model.Transaction
		expectedError string
	}{
		{
			name: "should fail when account does not exist",
			transaction: &model.Transaction{
				AccountID:       999,
				OperationTypeID: 1,
				Amount:          decimal.NewFromFloat(123.45),
			},
			expectedError: "account not found",
		},
		{
			name: "should fail when operation type is invalid",
			transaction: &model.Transaction{
				AccountID:       1,
				OperationTypeID: 999,
				Amount:          decimal.NewFromFloat(123.45),
			},
			expectedError: "invalid operation type ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			accountRepo := &mocks.AccountRepositoryMock{
				GetAccountByIDFunc: func(ctx context.Context, id uint64) (*model.Account, error) {
					assert.Equal(t, tt.transaction.AccountID, id)
					if tt.expectedError == "account not found" {
						return nil, errors.New("account not found")
					}
					return &model.Account{
						ID:             tt.transaction.AccountID,
						DocumentNumber: "12345678900",
					}, nil
				},
			}

			logger, _ := zap.NewDevelopment()

			service := service.NewService(accountRepo, nil, logger)

			response, err := service.CreateTransaction(ctx, tt.transaction)

			assert.Nil(t, response)
			assert.EqualError(t, err, tt.expectedError)
		})
	}

}

func TestOperationType_Success(t *testing.T) {
	tests := []struct {
		name     string
		op       model.OperationType
		expected bool
	}{
		{
			name:     "should be valid purchase debit operation type",
			op:       model.Purchase,
			expected: true,
		},
		{
			name:     "should be valid installment purchase debit operation type",
			op:       model.InstallmentPurchase,
			expected: true,
		},
		{
			name:     "should be valid withdrawal debit operation type",
			op:       model.Withdrawal,
			expected: true,
		},
		{
			name:     "should be a invalid debit operation type",
			op:       model.Payment,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.op.IsDebit())
		})
	}
}
