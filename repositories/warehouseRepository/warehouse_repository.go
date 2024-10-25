package warehouseRepository

import (
	"github.com/agussuartawan/golang-pos/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/errors"
	"github.com/agussuartawan/golang-pos/models"
)

var TABLE = "warehouses"

func Create(model models.Warehouse) error {
	result := config.DB.Create(&model)
	return result.Error
}

func Delete(id int) error {
	if exists, err := IsExist(id); err != nil {
		return err
	} else if !exists {
		return errors.ErrWarehouseNotFound
	}

	result := config.DB.Delete(&models.Warehouse{})
	return result.Error
}

func Update(id int, req request.WarehouseRequest) error {
	var model models.Warehouse
	if result := config.DB.First(&model); result.Error != nil {
		return result.Error
	}

	result := config.DB.Model(&model).Updates(req)
	return result.Error
}

func IsExist(id int) (bool, error) {
	var exists bool
	if result := config.DB.Model(&models.Warehouse{}).
		Where("id = ?", id).
		Limit(1).
		Scan(&exists); result.Error != nil {
		return false, result.Error
	}

	return exists, nil
}

func List(param request.CompanyParam) ([]response.WarehouseResponse, error) {
	var warehouses []response.WarehouseResponse
	result := config.DB.Table(TABLE).
		Joins("Company").
		Where(TABLE+".deleted_at IS NULL")

	if param.CompanyId != nil {
		result = result.Where(TABLE+".company_id = ?", *param.CompanyId)
	}
	if param.Name != nil {
		result = result.Where(TABLE+".name like ?", "%"+*param.Name+"%")
	}
	err := 	result.Order(TABLE+".created_at DESC").Find(&warehouses).Error

	return warehouses, err
}

func FindById(id int) (response.WarehouseResponse, error) {
	var warehouse response.WarehouseResponse
	result := config.DB.Table(TABLE).
		Joins("Company").
		Where(TABLE+".deleted_at is null and "+TABLE+".id = ?", id).
		First(&warehouse)

	return warehouse, result.Error
}