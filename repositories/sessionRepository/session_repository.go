package sessionRepository

import (
	"errors"
	"time"

	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/models"
	"gorm.io/gorm"
)

func Create(userId uint, expiredAt time.Time, ipAddress string) (string, error) {
	session := models.Session{
		UserId: userId,
		IpAddress: &ipAddress,
		ExpiredAt: expiredAt,
	}
	err := config.DB.Create(&session).Error
	return session.ID, err
}

func ClearSession(userId uint) error {
	return config.DB.Where("user_id = ?", userId).Delete(&models.Session{}).Error
}

func Get(session *payload.SessionPayload, id string) error {
	err := config.DB.Model(&models.Session{}).Joins("User").Where("sessions.id = ?", id).First(&session).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("session not found")
	}
	return err
}