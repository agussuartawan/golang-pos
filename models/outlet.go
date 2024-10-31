package models

import "gorm.io/gorm"

type Outlet struct {
	gorm.Model
	WarehouseID  uint      `gorm:"index"`
	Warehouse    Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SupervisorID uint      `gorm:"index"`
	Supervisor   User      `gorm:"foreignKey:SupervisorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name         string    `gorm:"unique;type:varchar(255)"`
	Address      *string
}
