package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	UserID   uint    `gorm:"unique"`
	User     *User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	OutletID uint    `gorm:"not null"`
	Outlet   Outlet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Identity uint    `gorm:"unique;index;not null"`
	Name     string  `gorm:"type:varchar(255);not null"`
	Position string  `gorm:"type:varchar(255);not null"`
	Pin      *string `gorm:"type:varchar(255)"`
}
