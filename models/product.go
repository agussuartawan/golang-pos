package models

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"gorm.io/gorm"
)

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

func (p *Product) BeforeSave(tx *gorm.DB) error {
	p.Name = helper.TrimSpace(p.Name)
	return nil
}
