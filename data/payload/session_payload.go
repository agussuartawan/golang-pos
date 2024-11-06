package payload

import (
	"encoding/json"
	"github.com/agussuartawan/golang-pos/data/response"
)

type SessionCookie struct {
	SessionId string        `json:"sessionId"`
	Token     string        `json:"token"`
	User      response.User `json:"user"`
}

func (s *SessionCookie) ToJSON() (string, error) {
	jsonSession, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(jsonSession), nil
}
