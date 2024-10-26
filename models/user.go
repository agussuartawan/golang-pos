package models

import (
	"strings"

	helper "github.com/agussuartawan/golang-pos/helpers"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Email string `gorm:"not null; type:varchar(255)"`
	Phone *string `gorm:"type:varchar(50)"`
	Password string `gorm:"not null; type:varchar(255)"`
	Roles []Role `gorm:"many2many:user_roles;"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Name = helper.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}