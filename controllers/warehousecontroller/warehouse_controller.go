package warehousecontroller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/agussuartawan/golang-pos/core/errors"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	companyRepo "github.com/agussuartawan/golang-pos/repositories/companyrepository"
	warehouseRepo "github.com/agussuartawan/golang-pos/repositories/warehouserepository"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	var param request.CompanyParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	warehouses, err := warehouseRepo.List(param)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(warehouses))
}

func FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	warehouse, err := warehouseRepo.FindById(id)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.OK(warehouse))
}

func Create(ctx *gin.Context) {
	log.Println("Membuat warehouse baru...")

	// parse json into request
	req := request.WarehouseRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// validate request
	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}
	// pastikan companyId ada di DB
	if exists, err := companyRepo.IsExists(req.CompanyId); err != nil {
		helper.ThrowError(ctx, err)
		return
	} else if !exists {
		helper.ThrowError(ctx, errors.ErrCompanyNotFound)
		return
	}

	// map request to model
	model := models.Warehouse{
		CompanyId:   req.CompanyId,
		Name:        req.Name,
		Description: req.Description,
	}

	// store to database
	if err := warehouseRepo.Create(model); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// give response
	ctx.JSON(http.StatusOK, response.OK(req))
}

func Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
	}

	log.Printf("Mengupdate warehouse dengan id %v...", id)

	// parse json into request
	req := request.WarehouseRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// validate request
	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	// pastikan companyId ada di DB
	if exists, err := companyRepo.IsExists(req.CompanyId); err != nil {
		helper.ThrowError(ctx, err)
		return
	} else if !exists {
		helper.ThrowError(ctx, errors.ErrCompanyNotFound)
		return
	}

	if err := warehouseRepo.Update(id, req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(req))
}

func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
	}

	log.Printf("Menghapus warehouse dengan id %v...", id)

	if err := warehouseRepo.Delete(uint(id)); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(nil))
}
