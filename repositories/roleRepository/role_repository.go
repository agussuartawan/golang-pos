package roleRepository

import (
	"errors"

	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/models"
	"gorm.io/gorm"
)

func Create(model models.Role) error {
	result := config.DB.Create(&model)
	return result.Error
}

func List() ([]models.Role, error) {
	var roles []models.Role  // Gunakan struct model asli
    // Query menggunakan model asli untuk memastikan relasi many2many berjalan
    err := config.DB.Model(&models.Role{}).
        Select("roles.created_at, roles.id, roles.name, roles.description").
        Preload("Permissions", func(db *gorm.DB) *gorm.DB {
            return db.Select("permissions.id, permissions.name")
        }).
        Where("roles.deleted_at IS NULL").
        Order("roles.created_at desc").
        Find(&roles).
		Error

    return roles, err
}

func AppendPermissions(role models.Role, permissions []models.Permission) error {
	return config.DB.Model(&role).Association("Permissions").Append(permissions)
}

func DeletePermissions(role models.Role, permissions models.Permission) error {
	return config.DB.Model(&role).Association("Permissions").Delete(permissions)
}

func Get(role *models.Role, id uint) error {
	err := config.DB.First(&role, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("roleId not found")
	}

	return err
}

func Gets(roles *[]models.Role, ids []uint) error {
	return config.DB.Find(&roles, ids).Error
}