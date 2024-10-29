package response

import "time"

type LoginResponse struct {
	Token string `json:"token"`
}

type ProfileResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone *string `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles []RolePermission `json:"roles"`
}