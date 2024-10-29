package payload

import "github.com/dgrijalva/jwt-go"

type ClaimPayload struct {
	SessionId string `json:"session"`
	jwt.StandardClaims
}