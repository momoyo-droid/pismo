package model

type OperationType int

// OperationType represents the type of a financial operation,
// such as purchase, installment purchase, withdrawal, or payment.
const (
	Purchase OperationType = iota + 1
	InstallmentPurchase
	Withdrawal
	Payment
)

// Operation represents a financial operation with an ID and a Description.
type Operation struct {
	ID          int
	Description string
}

// IsValid checks if the OperationType is one of the defined valid types.
func (o OperationType) IsValid() bool {
	switch o {
	case Purchase, InstallmentPurchase, Withdrawal, Payment:
		return true
	default:
		return false
	}
}

// IsDebit checks if the OperationType is a debit type (Purchase, InstallmentPurchase, Withdrawal).
// Returns true for debit operations and false for credit operations (Payment).
func (o OperationType) IsDebit() bool {
	switch o {
	case Purchase, InstallmentPurchase, Withdrawal: // Debit transactions
		return true
	default:
		return false // Credit transactions (Payment)
	}
}
