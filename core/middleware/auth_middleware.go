package middleware

import (
	"strings"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/repositories/sessionRepository"
	"github.com/agussuartawan/golang-pos/repositories/userRepository"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			helper.JSON401(c, "bearer token required")
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := validateToken(token)
		if err != nil {
			helper.ThrowError(c, err)
			c.Abort()
			return
		}

		var session payload.SessionPayload
		if err := sessionRepository.Get(&session, claims.SessionId); err != nil {
			helper.ThrowError(c, err)
			c.Abort()
			return
		}
		
		c.Set("session", session)
		c.Next()
	}
}

func Authorized(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			helper.JSON401(c, "bearer token required")
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := validateToken(token)
		if err != nil {
			helper.ThrowError(c, err)
			c.Abort()
			return
		}

		var session payload.SessionPayload
		if err := sessionRepository.Get(&session, claims.SessionId); err != nil {
			helper.ThrowError(c, err)
			c.Abort()
			return
		}

		// check permission
		permitted, err := userRepository.IsHasPermission(session.UserId, permission)
		if err != nil {
			helper.ThrowError(c, err)
			c.Abort()
			return
		}
		if !permitted {
			helper.JSON403(c, "You don't have permission")
			c.Abort()
			return
		}

		c.Set("session", session)
		c.Next()
	}
}

func validateToken(token string) (*payload.ClaimPayload, error) {
	claims, err := helper.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}