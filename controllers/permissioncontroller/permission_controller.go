package permissioncontroller

import (
	"log"
	"net/http"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/permissionrepository"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	log.Println("Membuat permission baru...")

	req := request.PermissionRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	permission := models.Permission{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := permissionrepository.Create(permission); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(req))
}

func List(ctx *gin.Context) {
	log.Println("Mengambil list permission...")

	var param request.PermissionParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	permissions, err := permissionrepository.List(&param)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	helper.JSONPaginate(ctx, param.PaginationParam, permissions)
}
