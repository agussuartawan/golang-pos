package rolecontroller

import (
	"log"
	"net/http"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/permissionrepository"
	"github.com/agussuartawan/golang-pos/repositories/rolerepository"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	log.Println("Mengambil list role...")

	roles, err := rolerepository.List()
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// Mapping dari models.Role ke response.RoleResponse
	var roleResponses []response.RoleResponse
	for _, role := range roles {
		roleResponse := response.RoleResponse{
			Id:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			Permissions: []response.Permission{},
		}

		for _, permission := range role.Permissions {
			roleResponse.Permissions = append(roleResponse.Permissions, response.Permission{
				Id:   permission.ID,
				Name: permission.Name,
			})
		}

		roleResponses = append(roleResponses, roleResponse)
	}

	ctx.JSON(http.StatusOK, response.OK(roleResponses))
}

func FindById(ctx *gin.Context) {}

func Create(ctx *gin.Context) {
	log.Println("Membuat role baru...")

	request := request.RoleRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(request); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	role := models.Role{
		Name:        request.Name,
		Description: request.Description,
	}
	if err := rolerepository.Create(role); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(request))
}

func Update(ctx *gin.Context) {}

func Delete(ctx *gin.Context) {}

func AppendPermissions(ctx *gin.Context) {
	log.Println("Menambahkan permission ke role...")

	req := request.AppendPermissionRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	// get role and permissions model
	var role models.Role
	if err := rolerepository.Get(&role, req.RoleId); err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	var permissions []models.Permission
	if err := permissionrepository.Gets(&permissions, req.PermissionIds); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := rolerepository.AppendPermissions(role, permissions); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	res := response.AppendPermissionsResponse{
		Id: role.ID,
		Permissions: func(permissions []models.Permission) []uint {
			ids := make([]uint, len(permissions))
			for i, permission := range permissions {
				ids[i] = permission.ID
			}
			return ids
		}(permissions),
	}
	ctx.JSON(http.StatusOK, response.OK(res))
}
