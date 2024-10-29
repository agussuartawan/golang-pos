package permissionController

import (
	"log"
	"net/http"

	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/permissionRepository"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	log.Println("Membuat permission baru...")

	request := request.PermissionRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	if err := helper.Validator(request); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	permission := models.Permission{
		Name: request.Name,
		Description: request.Description,
	}
	if err := permissionRepository.Create(permission); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(request))
}

func List(ctx *gin.Context) {
	log.Println("Mengambil list permission...")

	var param request.PermissionParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	permissions, err := permissionRepository.List(param)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(permissions))
}