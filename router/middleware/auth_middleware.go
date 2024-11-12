package middleware

import (
	"encoding/json"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get session from cookie
		sessionJSON, err := c.Cookie("session")
		if err != nil {
			helper.JSON401(c, "session not found")
			return
		}

		// validate cookie and get session
		var session payload.SessionCookie
		if _, err := validateCookie(&session, sessionJSON); err != nil {
			helper.JSON401(c, err.Error())
			return
		}

		c.Set("session", session)
		c.Next()
	}
}

func Authorized(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get session from cookie
		sessionJSON, err := c.Cookie("session")
		if err != nil {
			helper.JSON401(c, "session not found")
			return
		}

		// validate cookie and get session
		var session payload.SessionCookie
		claims, err := validateCookie(&session, sessionJSON)
		if err != nil {
			helper.JSON401(c, err.Error())
			return
		}

		if claims.IsSuperAdmin {
			c.Set("session", session)
			c.Next()
			return
		}

		// check permission
		permitted, err := userrepository.IsHasPermission(session.User.Id, permission)
		if err != nil {
			helper.ThrowError(c, err)
			return
		}
		if !permitted {
			helper.JSON403(c, "You don't have permission")
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

func validateCookie(session *payload.SessionCookie, sessionJSON string) (*payload.ClaimPayload, error) {
	// deserialize sessionJSON
	if err := json.Unmarshal([]byte(sessionJSON), &session); err != nil {
		return nil, err
	}

	claims, err := validateToken(session.Token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
