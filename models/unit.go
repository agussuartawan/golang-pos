package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name      string `gorm:"unique;type:varchar(255);not null"`
	BaseValue uint   `gorm:"type:int;not null;comment:menyimpan base value, misal unit kg berarti 1000gram"`
}
