package payload

import "github.com/dgrijalva/jwt-go"

type ClaimPayload struct {
	SessionId    string `json:"session"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
	jwt.StandardClaims
}
