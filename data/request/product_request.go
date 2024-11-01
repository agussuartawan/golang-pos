package request

import "time"

type ProductRequest struct {
	Name        string              `json:"name" validate:"required,max=255"`
	Size        uint                `json:"size" validate:"required"`
	UnitId      uint                `json:"unitId" validate:"required"`
	CategoryIds []uint              `json:"categoryIds"`
	Price       ProductPriceRequest `json:"price" validate:"required"`
}

type ProductPriceRequest struct {
	Value     float64   `json:"value" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate   time.Time `json:"endDate" validate:"required"`
}

type ProductParam struct {
	PaginationParam
	Name *string `form:"name"`
}
