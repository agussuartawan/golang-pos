package request

type CompanyRequest struct {
	Name    string `json:"name" form:"name" validate:"required,max=255"`
	Phone   *string `json:"phone" form:"phone" validate:"omitempty,max=50"`
	Address *string `json:"address" form:"address"`
	Email   *string `json:"email" form:"email" validate:"omitempty,max=255"`
}

type CompanyParam struct {
	CompanyId *int
	Name *string
}