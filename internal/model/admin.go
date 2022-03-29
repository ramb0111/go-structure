package model

import "time"

//Admin Model for admin table
type Admin struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(255)" json:"username"`
	Password  string `gorm:"type:varchar(255)" json:"password,omitempty"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
