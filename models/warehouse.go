package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	CompanyId int `gorm:"not null"`
	Company Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name string `gorm:"not null; type:varchar(255)"`
	Description *string
}