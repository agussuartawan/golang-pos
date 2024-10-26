package models

import (
	"strings"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name string `gorm:"not null; type:varchar(255)"`
	Phone *string `gorm:"type:varchar(50)"`
	Address *string 
	Email *string `gorm:"type:varchar(255)"`
}

func (c *Company) BeforeSave(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	if c.Email != nil {
        trimmedEmail := strings.TrimSpace(*c.Email)
        c.Email = &trimmedEmail
    }
	if c.Phone != nil {
		trimmedPhone := strings.TrimSpace(*c.Phone)
		c.Phone = &trimmedPhone
	}
	return nil
}