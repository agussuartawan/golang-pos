package response

import (
	"time"
)

type LoginResponse struct {
	Name        string   `json:"name"`
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type ProfileResponse struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Phone       *string      `json:"phone"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Roles       []Role       `json:"roles,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}
