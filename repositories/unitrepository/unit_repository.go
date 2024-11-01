package unitrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
)

func Create(model *models.Unit) error {
	return config.DB.Create(&model).Error
}

func List(res *[]response.UnitResponse, param *request.UnitParam) error {
	var query = config.DB.Model(&models.Unit{})

	if param.Name != nil {
		query = query.Where("name ILIKE '%?%'", *param.Name)
	}

	return param.Paginate(query).Find(&res).Order("created_at desc").Error
}

func IsExists(id uint) (bool, error) {
	var exists bool
	if err := config.DB.Model(&models.Unit{}).
		Select("1").
		Where("id = ?", id).
		Where("deleted_at is null").
		Scan(&exists).
		Error; err != nil {
		return false, err
	}

	return exists, nil
}
