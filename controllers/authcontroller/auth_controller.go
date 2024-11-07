package authcontroller

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/sessionrepository"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"github.com/agussuartawan/golang-pos/services/authservice"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"slices"
)

func Login(ctx *gin.Context) {
	req := request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// bind json to struct and validate
	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	var res response.LoginResponse
	var session payload.SessionCookie
	if err := authservice.Login(req, &res, &session); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// success login
	log.Printf("%v telah login", res.Name)
	maxAge := 30 * 24 * 60 * 60
	sessionJSON, err := session.ToJSON()
	if err != nil {
		slog.Error("when parsing session to JSON")
		helper.ThrowError(ctx, err)
		return
	}
	ctx.SetCookie("session", sessionJSON, maxAge, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, response.OK(res))
}

func Logout(ctx *gin.Context) {
	var session payload.SessionCookie
	if err := helper.GetPrincipal(ctx, &session); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := sessionrepository.DeleteSession(session.SessionId); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.SetCookie("session", "", -1, "/", "localhost", false, true)
	helper.JSON200(ctx, "Logout success")
}

func Profile(ctx *gin.Context) {
	var sessionStruct payload.SessionCookie
	if err := helper.GetPrincipal(ctx, &sessionStruct); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	user := models.User{}
	if err := userrepository.GetProfile(sessionStruct.User.Id, &user); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// mapping models.User ke response.ProfileResponse
	profileResponse := response.ProfileResponse{
		Id:        user.ID,
		Name:      sessionStruct.User.Name,
		Email:     sessionStruct.User.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	for _, role := range user.Roles {
		profileResponse.Roles = append(profileResponse.Roles, response.Role{Id: role.ID, Name: role.Name})
		for _, permission := range role.Permissions {
			p := response.Permission{Id: permission.ID, Name: permission.Name}
			if !slices.Contains(profileResponse.Permissions, p) {
				profileResponse.Permissions = append(profileResponse.Permissions, p)
			}
		}
	}

	ctx.JSON(http.StatusOK, response.OK(profileResponse))
}
