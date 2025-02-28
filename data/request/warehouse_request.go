package request

type WarehouseRequest struct {
	CompanyId int `json:"companyId" form:"companyId" validate:"required"`
	Name string `json:"name" form:"name" validate:"required,max=255"`
	Description *string `json:"description" form:"description"`
}

type WarehouseParam struct {
	CompanyId int
	Name string
}