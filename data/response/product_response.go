package response

type ProductResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Size  uint    `json:"size"`
	Unit  Unit    `json:"unit"`
}
