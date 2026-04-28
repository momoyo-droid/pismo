package model

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

func (o OperationType) IsDebit() bool {
	switch o {
	case Purchase, InstallmentPurchase, Withdrawal: // Debit transactions
		return true
	default:
		return false // Credit transactions (Payment)
	}
}
