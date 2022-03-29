package model

import "time"

//MilkmanCustomers Model for customer address table
type MilkmanCustomers struct {
	ID            uint `gorm:"primary_key"`
	UserID        int
	User          Users `json:"customer"`
	MilkmanID     int
	Milkman       Users   `json:"milkman"`
	BalanceAmount float32 `json:"balance_amount" gorm:"type:decimal(10,8);"`
	CreatedAt     time.Time
	DeletedAt     *time.Time
}
