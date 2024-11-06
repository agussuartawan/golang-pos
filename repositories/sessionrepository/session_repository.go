package sessionrepository

import (
	"time"

	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/models"
)

func Create(userId uint, expiredAt time.Time, ipAddress string) (string, error) {
	session := models.Session{
		UserId:    userId,
		IpAddress: &ipAddress,
		ExpiredAt: expiredAt,
	}
	err := config.DB.Create(&session).Error
	return session.ID, err
}

func ClearSession(userId uint) error {
	return config.DB.Where("user_id = ?", userId).Delete(&models.Session{}).Error
}

func DeleteSession(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Session{}).Error
}
