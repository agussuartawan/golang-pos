package outletrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
)

func Create(model *models.Outlet) error {
	return config.DB.Create(&model).Error
}

func List(res *[]response.OutletResponse, param request.OutletParam) (int64, error) {
	query := config.DB.Model(&models.Outlet{}).Joins("Warehouse").Joins("Supervisor")

	if param.Name != nil {
		query = query.Where("outlets.name ILIKE ?", "%"+*param.Name+"%")
	}
	if param.WarehouseID != nil {
		query = query.Where("warehouse_id = ?", param.WarehouseID)
	}
	if param.SupervisorID != nil {
		query = query.Where("supervisor_id = ?", param.SupervisorID)
	}

	// count total data
	var total int64
	query.Count(&total)

	return total, param.Paginate(query).Order("created_at desc").Find(&res).Error
}
