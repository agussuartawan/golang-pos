package unitrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
)

func Create(model *models.Unit) error {
	return config.DB.Create(&model).Error
}

func List(res *[]response.UnitResponse) error {
	return config.DB.Model(&models.Unit{}).Where("deleted_at is null").Find(&res).Order("created_at desc").Error
}
