package request

type UnitRequest struct {
	Name      string `json:"name" validate:"required,max=255"`
	BaseValue uint   `json:"baseValue" validate:"required"`
}
