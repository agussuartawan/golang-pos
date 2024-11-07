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

	if param.Query != nil {
		query = query.Where("name ilike '%?%' and base_value = ?", *param.Query, *param.Query)
	}

	query = param.Paginate(query)
	switch param.SortBy {
	case "name":
		query = query.Order("name " + param.SortBy)
	case "baseValue":
		query = query.Order("base_value " + param.SortBy)
	case "createdAt":
		query = query.Order("created_at " + param.SortBy)
	default:
		query = query.Order("created_at asc")

	}

	return query.Find(&res).Error
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
