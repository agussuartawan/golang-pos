package models

import (
	"gorm.io/gorm"
	"time"
)

type Price struct {
	gorm.Model
	Value     float64   `gorm:"type:double"`
	StartDate time.Time `gorm:"type:datetime"`
	EndDate   time.Time `gorm:"type:datetime"`
}
