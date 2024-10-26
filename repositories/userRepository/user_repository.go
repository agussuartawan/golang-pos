package userRepository

import (
	"errors"

	"github.com/agussuartawan/golang-pos/config"
	"github.com/agussuartawan/golang-pos/models"
	"gorm.io/gorm"
)

func Create(model models.User) error {
	err := config.DB.Create(&model).Error
	return err
}

func AppendRoles(u models.User, roles []models.Role) error {
	return config.DB.Model(&u).Association("Roles").Append(roles)
}

func Get(u *models.User, id uint) error {
	err := config.DB.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("userId not found")
	}

	return err
}

func List() ([]models.User, error) {
	var users []models.User
	err := config.DB.Model(&users).
		Select("users.id, users.name, users.email, users.phone, users.created_at").
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("roles.id, roles.name")
		}).
		Where("users.deleted_at IS NULL").
		Order("users.created_at desc").
		Find(&users).
		Error

	return users, err
}