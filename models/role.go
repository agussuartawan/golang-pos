package models

import (
	"strings"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Description *string `gorm:"type:varchar(255)"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (r *Role) BeforeSave(tx *gorm.DB) error {
	r.Name = strings.ToLower(strings.TrimSpace(r.Name))
	return nil
}