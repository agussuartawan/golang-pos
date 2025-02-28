package usercontroller

import (
	"log"
	"net/http"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/rolerepository"
	"github.com/agussuartawan/golang-pos/repositories/userrepository"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	log.Println("Membuat user baru...")

	var req request.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	user := models.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Email:    req.Email,
	}
	if err := userrepository.Create(user); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(nil))
}

func AppendRoles(ctx *gin.Context) {
	log.Println("Mengappend role ke user...")

	req := request.AppendRoleRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	user := models.User{}
	if err := userrepository.Get(&user, req.UserId); err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	var roles []models.Role
	if err := rolerepository.Gets(&roles, req.RoleIds); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := userrepository.AppendRoles(user, roles); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	res := response.AppendRolesResponse{
		UserId: user.ID,
		RoleIds: func(roles []models.Role) []uint {
			ids := make([]uint, len(roles))
			for i, role := range roles {
				ids[i] = role.ID
			}
			return ids
		}(roles),
	}
	ctx.JSON(http.StatusOK, response.OK(res))
}

func List(ctx *gin.Context) {
	log.Println("Mengambil list user...")

	users, err := userrepository.List()
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
		}

		for _, role := range user.Roles {
			userResponse.Roles = append(userResponse.Roles, response.Role{
				Id:   role.ID,
				Name: role.Name,
			})
			for _, permission := range role.Permissions {
				userResponse.Permissions = append(userResponse.Permissions, response.Permission{
					Id:   permission.ID,
					Name: permission.Name,
				})
			}
		}
		userResponses = append(userResponses, userResponse)
	}

	ctx.JSON(http.StatusOK, response.OK(userResponses))
}
