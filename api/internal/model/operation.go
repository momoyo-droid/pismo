package model

import "github.com/shopspring/decimal"

type OperationType int

const (
	Purchase OperationType = iota + 1
	InstallmentPurchase
	Withdrawal
	Payment
)

type Operation struct {
	ID          int
	Description string
}

func (o OperationType) IsValid() bool {
	switch o {
	case Purchase, InstallmentPurchase, Withdrawal, Payment:
		return true
	default:
		return false
	}
}

func (o OperationType) IsDebitOrCredit() decimal.Decimal{
	switch o {
		case Purchase, InstallmentPurchase, Withdrawal: // Debit transactions
			return decimal.NewFromInt(-1)
	}
	return decimal.NewFromInt(1) // Credit transactions
}