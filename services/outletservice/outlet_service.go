package outletservice

import (
	"errors"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/outletrepository"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"github.com/agussuartawan/golang-pos/repositories/warehouserepository"
)

func Create(req request.OutletRequest) (*response.OutletResponse, error) {
	// pastikan warehouseId valid
	exists, err := warehouserepository.IsExist(req.WarehouseID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("warehouseId not found")
	}

	// pastikan supervisorId valid
	exists, err = userrepository.IsExists(req.SupervisorID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("supervisorId not found")
	}

	// save to db
	model := models.Outlet{
		WarehouseID:  req.WarehouseID,
		SupervisorID: req.SupervisorID,
		Name:         req.Name,
		Address:      req.Address,
	}
	if err = outletrepository.Create(&model); err != nil {
		return nil, err
	}

	// return response
	return &response.OutletResponse{
		ID:      model.ID,
		Name:    model.Name,
		Address: model.Address,
	}, nil
}
