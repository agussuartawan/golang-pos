package models

import "gorm.io/gorm"

type ProductStock struct {
	gorm.Model
	ProductID   uint      `gorm:"index;not null"`
	SupplierID  uint      `gorm:"index;not null"`
	WarehouseID uint      `gorm:"index;not null"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      uint      `gorm:"index;not null;comment:user yang melakukan aksi"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty         uint      `gorm:"not null"`
}
