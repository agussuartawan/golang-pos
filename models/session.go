package models

import "time"

type Session struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	UserId uint `gorm:"not null"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	IpAddress *string `gorm:"type:varchar(50)"`
	ExpiredAt time.Time
	CreatedAt time.Time
}