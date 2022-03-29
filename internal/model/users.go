package model

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//Users Model for admin table
type Users struct {
	ID              uint           `gorm:"primary_key"`
	Mobile          string         `gorm:"size:255" form:"mobile" binding:"required"`
	Password        string         `gorm:"type:varchar(255)" json:"password,omitempty" form:"password" binding:"required"`
	Email           string         `gorm:"type:varchar(100)"`
	Role            pq.StringArray `gorm:"type:varchar(100)[]"` //CUSTOMER,MILKMAN
	Status          int
	Latitude        float32 `json:"latitude" gorm:"type:decimal(10,8);"`
	Longitude       float32 `json:"longitude" gorm:"type:decimal(10,8);"`
	VerificationOtp string  `gorm:"size:10"`
	OtpExpireAt     *time.Time
	SignUpCode      string          `gorm:"size:255" gorm:"default:null"`
	DairymanToken   string          `gorm:"size:255"  gorm:"default:null"`
	Preference      json.RawMessage `sql:"type:json" gorm:"default:'{}'"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

//UserData used for default user object
type UserData struct {
	ID         uint64          `json:"id"`
	Mobile     string          `json:"mobile"`
	Email      string          `json:"email"`
	Status     int             `json:"status"`
	Latitude   float32         `json:"latitude"`
	Longitude  float32         `json:"longitude"`
	Preference json.RawMessage `json:"preference"`
	Role       pq.StringArray  `json:"role"`
}

//BeforeCreate hook
func (user *Users) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
	return nil
}
