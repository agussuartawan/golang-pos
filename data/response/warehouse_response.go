package response

type WarehouseResponse struct {
	Id          int                    `json:"id"`
	Name        string                 `json:"name"`
	Description *string                `json:"description"`
	CompanyId   int                    `json:"-"`
	Company     Company  			`json:"company"`
}