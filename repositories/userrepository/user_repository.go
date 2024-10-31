package userrepository

import (
	"errors"
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/models"
	"gorm.io/gorm"
)

func Create(model models.User) error {
	err := config.DB.Create(&model).Error
	return err
}

func FirstOrCreate(model models.User) error {
	return config.DB.FirstOrCreate(&model, models.User{Email: model.Email}).Error
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

func GetByEmail(user *models.User, email string) error {
	err := config.DB.Model(&models.User{}).
		Select("users.id, users.name, users.password").
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("roles.id, roles.name")
		}).
		Preload("Roles.Permissions", func(db *gorm.DB) *gorm.DB {
			return db.Select("permissions.id, permissions.name")
		}).
		Where("email = ?", email).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("email not found")
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
		Preload("Roles.Permissions", func(db *gorm.DB) *gorm.DB {
			return db.Select("permissions.id, permissions.name")
		}).
		Where("users.deleted_at IS NULL").
		Order("users.created_at desc").
		Find(&users).
		Error

	return users, err
}

func IsHasPermission(id uint, permission string) (bool, error) {
	var exists bool
	if result := config.DB.Model(&models.User{}).
		Select("1").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("users.id = ?", id).
		Where("users.deleted_at IS NULL").
		Where("permissions.name = ?", permission).
		Limit(1).
		Scan(&exists); result.Error != nil {
		return false, result.Error
	}

	return exists, nil
}

func IsHasRole(id uint, role string) (bool, error) {
	var yes bool
	if err := config.DB.Model(&models.User{}).
		Select("1").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("users.deleted_at IS NULL").
		Where("users.id = ?", id).
		Where("roles.name = ?", role).
		Scan(&yes).
		Error; err != nil {
		return false, err
	}

	return yes, nil
}

func GetProfile(id uint, user *models.User) error {
	return config.DB.Model(&models.User{}).
		Select("users.id, users.phone, users.created_at, users.updated_at").
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("roles.id, roles.name")
		}).
		Preload("Roles.Permissions", func(db *gorm.DB) *gorm.DB {
			return db.Select("permissions.id, permissions.name")
		}).
		Preload("Roles.Permissions").
		First(&user, id).
		Error
}

func IsExists(id uint) (bool, error) {
	var exists bool
	if err := config.DB.Model(&models.User{}).
		Select("1").
		Where("users.id = ?", id).
		Limit(1).
		Scan(&exists).
		Error; err != nil {
		return false, err
	}

	return exists, nil
}
