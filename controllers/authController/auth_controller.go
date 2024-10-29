package authController

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/payload"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/sessionRepository"
	"github.com/agussuartawan/golang-pos/repositories/userRepository"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	// prepare data from env
	JWTExpiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// bind json to struct and validate
	req := request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	// get user from DB
	user := response.UserLoginResponse{}
	if err := userRepository.GetByEmail(&user, req.Email); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// validate password
	if err := user.ValidatePassword(req.Password); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// create session and clear previous session
	if err := sessionRepository.ClearSession(user.ID); err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	expiredAt := time.Now().AddDate(0, 0, JWTExpiration)
	sessionId, err := sessionRepository.Create(user.ID, expiredAt, ctx.ClientIP())
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// generate jwt
	token, err := helper.CreateToken(sessionId, expiredAt)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// give response
	res := response.LoginResponse{
		Token: token,
	}
	log.Printf("%v telah login", user.Name)
	ctx.JSON(http.StatusOK, response.OK(res))
}

func Profile(ctx *gin.Context) {
	session, exists := ctx.Get("session")
	if !exists {
		helper.ThrowError(ctx, errors.New("session not found"))
		return
	}

	sessionStruct := session.(payload.SessionPayload)
	user := models.User{}
	if err := userRepository.GetProfile(sessionStruct.User.Id, &user); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// mapping models.User ke response.ProfileResponse
	profileResponse := response.ProfileResponse{
		Id: user.ID,
		Name: sessionStruct.User.Name,
		Email: sessionStruct.User.Email,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Roles: func (roles []models.Role) []response.RolePermission {
			res := []response.RolePermission{}
			for _, role := range roles {
				res = append(res, response.RolePermission{
					Id: role.ID,
					Name: role.Name,
					Permissions: func (permissions []models.Permission) []response.Permission {
						res := []response.Permission{}
						for _, permission := range permissions {
							res = append(res, response.Permission{
								Id: permission.ID,
								Name: permission.Name,
							})
						}
						return res
					}(role.Permissions),
				})
			}
			return res
		}(user.Roles),
	}

	ctx.JSON(http.StatusOK, response.OK(profileResponse))
}