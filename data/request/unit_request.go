package request

type UnitRequest struct {
	Name      string `json:"name" validate:"required,max=255"`
	BaseValue uint   `json:"baseValue" validate:"required"`
}

type UnitParam struct {
	Name *string `form:"name"`
	PaginationParam
}
