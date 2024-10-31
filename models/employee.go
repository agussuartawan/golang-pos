package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	UserID   uint      `gorm:"unique"`
	Employee *Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	OutletID uint
	Outlet   Outlet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Identity uint    `gorm:"unique;index"`
	Name     string  `gorm:"type:varchar(255)"`
	Position string  `gorm:"type:varchar(255)"`
	Pin      *string `gorm:"type:varchar(255)"`
}
