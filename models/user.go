package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Email string `gorm:"not null; type:varchar(255)"`
	Phone *string `gorm:"type:varchar(50)"`
	Password string `gorm:"not null; type:varchar(255)"`
	Roles []Role `gorm:"many2many:user_roles;"`
}