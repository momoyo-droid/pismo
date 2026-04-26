package entity

import "github.com/shopspring/decimal"

type Transaction struct {
	OperationType string          `gorm:"column:operation_type;not null"`
	Amount        decimal.Decimal `gorm:"column:amount;not null"`
}
