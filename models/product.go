package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UnitId     uint           `gorm:"index;not null"`
	Unit       Unit           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Categories []Category     `gorm:"many2many:product_categories;"`
	Prices     []ProductPrice `gorm:"foreignKey:ProductID;"`
	Stocks     []ProductStock `gorm:"foreignKey:ProductID;"`
	Name       string         `gorm:"type:varchar(255);not null"`
	Size       uint           `gorm:"not null"`
}
