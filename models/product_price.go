package models

import (
	"gorm.io/gorm"
	"time"
)

type ProductPrice struct {
	gorm.Model
	ProductID uint      `gorm:"not null"`
	Value     float64   `gorm:"type:numeric(10,2);not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}
