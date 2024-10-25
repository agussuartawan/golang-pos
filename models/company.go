package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Phone *string `gorm:"type:varchar(50)"`
	Address *string 
	Email *string `gorm:"type:varchar(255)"`
}