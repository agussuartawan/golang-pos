package response

type OutletResponse struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Address      *string    `json:"address"`
	WarehouseId  int        `json:"-"`
	Warehouse    *Warehouse `json:"warehouse,omitempty"`
	SupervisorId int        `json:"-"`
	Supervisor   *User      `json:"supervisor,omitempty"`
}
