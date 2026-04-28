package entity

// Account represents a bank account with an ID and a DocumentNumber.
type Account struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	DocumentNumber string `gorm:"column:document_number;unique;not null"`
}
