package helper

import (
	"errors"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/gin-gonic/gin"
)

func GetPrincipal(ctx *gin.Context, session *payload.SessionCookie) error {
	s, exists := ctx.Get("session")
	if !exists {
		return errors.New("session not exists")
	}

	sessionStruct := s.(payload.SessionCookie)
	*session = sessionStruct

	return nil
}
