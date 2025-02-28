package permissionrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
)

func Create(model models.Permission) error {
	err := config.DB.Create(&model).Error
	return err
}

func List(param *request.PermissionParam) ([]response.PermissionResponse, error) {
	var permissions []response.PermissionResponse
	query := config.DB.Table("permissions").Where("deleted_at IS NULL")

	if param.Name != nil {
		query = query.Where("name ilike ?", "%"+*param.Name+"%")
	}

	if param.Query != nil {
		query = query.Where("name ilike '%?%' and description ilike '%?%'", *param.Query, *param.Query)
	}

	query = param.Paginate(query)
	switch param.SortBy {
	case "name":
		query = query.Order("name " + param.Sort)
	case "createdAt":
		query = query.Order("created_at " + param.Sort)
	default:
		query = query.Order("createdAt desc")
	}

	err := query.Find(&permissions).Error
	return permissions, err
}

func Gets(permissions *[]models.Permission, ids []uint) error {
	return config.DB.Find(&permissions, ids).Error
}
