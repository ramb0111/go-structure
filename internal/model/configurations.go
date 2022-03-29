package model

//Configurations Model for admin table
type Configurations struct {
	ID    uint   `gorm:"primary_key"`
	Key   string `gorm:"type:varchar(255)" json:"key"`
	Value string `gorm:"type:text" json:"value"`
}
