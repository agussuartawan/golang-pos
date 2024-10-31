package request

type OutletRequest struct {
	WarehouseID  uint    `json:"warehouseId" validate:"required"`
	SupervisorID uint    `json:"supervisorId" validate:"required"`
	Name         string  `json:"name" validate:"required,max=255"`
	Address      *string `json:"address"`
}

type OutletParam struct {
	Name         *string `form:"name"`
	WarehouseID  *int    `form:"warehouseId"`
	SupervisorID *int    `form:"supervisorId"`
	PaginationParam
}
