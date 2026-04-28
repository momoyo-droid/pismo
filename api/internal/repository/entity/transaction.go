package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Transaction represents a financial transaction with an ID,
// associated AccountID, OperationTypeID, and Amount.
type Transaction struct {
	TransactionID   uint64          `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	AccountID       uint64          `gorm:"column:account_id;not null"`
	Account         Account         `gorm:"foreignKey:AccountID;references:ID"`
	OperationTypeID int             `gorm:"column:operation_type_id;not null"`
	Operation       Operation       `gorm:"foreignKey:OperationTypeID;references:ID"`
	Amount          decimal.Decimal `gorm:"column:amount;not null"`
	CreationDate    time.Time       `gorm:"column:creation_date;not null"`
}
