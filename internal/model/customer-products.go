package model

import "time"

//CustomerProducts Model for customer products table
type CustomerProducts struct {
	ID         uint `gorm:"primary_key"`
	CustomerID int
	Customer   MilkmanCustomers `json:"customer"`
	ProductID  int
	Product    MilkmanProducts `json:"milkman_products"`
	MorningQty float32         `json:"morning_qty" gorm:"type:decimal(10,8);"`
	EveningQty float32         `json:"evening_qty" gorm:"type:decimal(10,8);"`
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
