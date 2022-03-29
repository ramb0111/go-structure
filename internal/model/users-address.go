package model

import "time"

//UsersAddress Model for customer address table
type UsersAddress struct {
	ID        uint `gorm:"primary_key"`
	UserID    int
	User      Users
	Address1  string `gorm:"size:255"`
	Address2  string `gorm:"size:255"`
	City      string `gorm:"size:255"`
	PinCode   string `gorm:"size:255"`
	State     string `gorm:"size:255"`
	IsDairy   bool   `sql:"default:false"`
	DairyName string `gorm:"size:255"`
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
