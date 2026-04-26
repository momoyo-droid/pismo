package model

import "github.com/shopspring/decimal"

type Transaction struct {
	OperationType string
	Amount        decimal.Decimal
}
