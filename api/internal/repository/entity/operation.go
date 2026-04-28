package entity

type Operation struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Description string `gorm:"column:description;not null"`
}
