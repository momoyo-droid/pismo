package model

import "github.com/shopspring/decimal"

// Transaction represents a financial transaction with an ID,
// associated AccountID, OperationTypeID, and Amount.
type Transaction struct {
	TransactionID   uint64
	AccountID       uint64
	OperationTypeID int
	Amount          decimal.Decimal
}
