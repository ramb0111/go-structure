package model

import "github.com/jinzhu/gorm/dialects/postgres"

//Localization Model for admin table
type Localization struct {
	ID       uint           `gorm:"primary_key"`
	Language string         `gorm:"type:varchar(255)" json:"language"`
	Values   postgres.Jsonb `json:"values"`
}
