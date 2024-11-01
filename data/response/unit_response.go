package response

type UnitResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	BaseValue uint   `json:"baseValue"`
}

type Unit struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
