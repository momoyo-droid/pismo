package entity

// Operation represents a financial operation with an ID and a Description.
type Operation struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Description string `gorm:"column:description;not null"`
}
