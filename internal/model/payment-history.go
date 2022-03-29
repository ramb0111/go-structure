package model

import (
	"time"
)

//PaymentHistory Model for payment history table
type PaymentHistory struct {
	ID            uint `gorm:"primary_key"`
	CustomerID    int
	Customer      MilkmanCustomers
	MilkmanID     int
	Milkman       Users
	PaymentAmount float32 `json:"payment_amount" gorm:"type:decimal(10,8);"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
