package response

import "time"

type PermissionResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description"`
	CreatedAt time.Time `json:"createdAt"`
}

type Permission struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}