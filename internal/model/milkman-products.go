package model

import "time"

//MilkmanProducts Model for milkman products table
type MilkmanProducts struct {
	ID          uint `gorm:"primary_key"`
	MilkmanID   int
	Milkman     Users   `json:"milkman"`
	ProductName string  `gorm:"size:255"`
	ProductUnit string  `gorm:"size:255"`
	Rate        float32 `json:"rate" gorm:"type:decimal(10,8);"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
