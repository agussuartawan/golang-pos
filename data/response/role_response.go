package response

import "time"

type RoleResponse struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	CreatedAt   time.Time   `json:"createdAt"`
	Permissions []Permission `json:"permissions"`
}

type AppendPermissionsResponse struct {
	Id uint `json:"roleId"`
	Permissions []uint `json:"permissionIds"`
}