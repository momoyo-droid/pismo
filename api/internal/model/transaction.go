package model

import "github.com/shopspring/decimal"

type Transaction struct {
	TransactionID   uint64
	AccountID       uint64
	OperationTypeID int
	Amount          decimal.Decimal
}
