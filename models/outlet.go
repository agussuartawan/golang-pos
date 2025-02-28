package models

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"gorm.io/gorm"
)

type Outlet struct {
	gorm.Model
	WarehouseID  uint      `gorm:"index;not null"`
	Warehouse    Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SupervisorID uint      `gorm:"index;not null"`
	Supervisor   User      `gorm:"foreignKey:SupervisorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name         string    `gorm:"unique;type:varchar(255);not null"`
	Address      *string
}

func (o *Outlet) BeforeSave(tx *gorm.DB) (err error) {
	o.Name = helper.TrimSpace(o.Name)
	return nil
}
