package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Description *string `gorm:"type:varchar(255)"`
}