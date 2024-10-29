package companyRepository

import (
	"github.com/agussuartawan/golang-pos/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/errors"
	"github.com/agussuartawan/golang-pos/models"
)

const TABLE = "companies"

func List() ([]response.CompanyResponse, error) {
	var companies []response.CompanyResponse
	result := config.DB.Table(TABLE).Where("deleted_at IS NULL").Find(&companies).Order("created_at desc")
	return companies, result.Error
}

func FindById(id int) (response.CompanyResponse, error) {
	var company response.CompanyResponse
	result := config.DB.Table(TABLE).
		Where("deleted_at IS NULL").
		First(&company, id)
	return company, result.Error
}

func Create(model models.Company) error {
	result := config.DB.Create(&model)
	return result.Error
}

func Update(id int, request request.CompanyRequest) error {
	var company models.Company
	err := config.DB.First(&company, id).Error
	if err != nil {
		return err
	}

	result := config.DB.Model(&company).Updates(request)
	return result.Error
}

func Delete(id int) error {
	exists, err := IsExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.ErrCompanyNotFound
	}

	result := config.DB.Delete(&models.Company{}, id)
	return result.Error
}

func IsExists(id int) (bool, error) {
	var exists bool
	if result := config.DB.Model(&models.Company{}).
		Select("1").
		Where("id = ?", id).
		Limit(1).
		Scan(&exists); result.Error != nil {
		return false, result.Error
	}

	return exists, nil
}