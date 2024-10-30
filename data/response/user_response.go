package response

import (
	"time"
)

type AppendRolesResponse struct {
	UserId  uint   `json:"userId"`
	RoleIds []uint `json:"roleIds"`
}

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Phone       *string      `json:"phone"`
	CreatedAt   time.Time    `json:"createdAt"`
	Roles       []Role       `json:"roles,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}
