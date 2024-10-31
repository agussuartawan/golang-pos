package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	ProductStocks []ProductStock `gorm:"foreignKey:SupplierID"`
	Name          string         `gorm:"type:varchar(100);not null"`
	Address       *string
	Phone         *string `gorm:"type:varchar(100)"`
	Pic           *string `gorm:"type:varchar(255)"`
}
