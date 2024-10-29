package payload

import (
	"github.com/agussuartawan/golang-pos/data/response"
)

type SessionPayload struct {
	Id string `json:"session"`
	UserId uint `json:"-"`
	User response.User `json:"user"`
}