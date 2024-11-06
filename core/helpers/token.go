package helper

import (
	"errors"
	"os"
	"time"

	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(sessionId string, isSuperAdmin bool, expiresAt time.Time) (string, error) {
	SecretKey := os.Getenv("SECRET_KEY")

	claims := &payload.ClaimPayload{
		SessionId:    sessionId,
		IsSuperAdmin: isSuperAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenStruct.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(token string) (*payload.ClaimPayload, error) {
	SecretKey := os.Getenv("SECRET_KEY")

	var claims payload.ClaimPayload
	tokenStruct, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt unexpected signing method")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tokenStruct.Valid {
		return nil, errors.New("jwt token is invalid")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("jwt token is expired")
	}

	return &claims, nil
}
