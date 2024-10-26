package models

import (
	"strings"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Description *string `gorm:"type:varchar(255)"`
}

func (p *Permission) BeforeSave(tx *gorm.DB) error {
	p.Name = strings.ToLower(strings.TrimSpace(p.Name))
	return nil
}