package userController

import (
	"log"
	"net/http"

	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	helper "github.com/agussuartawan/golang-pos/helpers"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/roleRepository"
	"github.com/agussuartawan/golang-pos/repositories/userRepository"
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
		Name: req.Name,
		Phone: req.Phone,
		Password: req.Password,
		Email: req.Email,
	}
	if err := userRepository.Create(user); err != nil {
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
	if err := userRepository.Get(&user, req.UserId); err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	roles := []models.Role{}
	if err := roleRepository.Gets(&roles, req.RoleIds); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := userRepository.AppendRoles(user, roles); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	res := response.AppendRolesResponse{
		UserId: user.ID,
		RoleIds: func (roles []models.Role) []uint {
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

	users, err := userRepository.List()
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	var userResponse []response.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, response.UserResponse{
			Id: user.ID,
			Name: user.Name,
			Email: user.Email,
			Phone: user.Phone,
			CreatedAt: user.CreatedAt,
			Roles: func (roles []models.Role) []response.Role {
				var res []response.Role
				for _, role := range roles {
					res = append(res, response.Role{
						Id: role.ID,
						Name: role.Name,
					})
				}
				return res
			}(user.Roles),
		})
	}

	ctx.JSON(http.StatusOK, response.OK(userResponse))
}