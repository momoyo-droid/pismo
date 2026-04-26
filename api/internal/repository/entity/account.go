package entity

type Account struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	DocumentNumber string `gorm:"column:document_number;unique;not null"`
}