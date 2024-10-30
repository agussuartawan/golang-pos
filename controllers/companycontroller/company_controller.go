package companycontroller

import (
	"log"
	"net/http"
	"strconv"

	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	companyRepo "github.com/agussuartawan/golang-pos/repositories/companyrepository"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	log.Println("Mengambil list company...")
	companies, err := companyRepo.List()
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(companies))
}

func FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
	}
	log.Println("Mengambil company dengan id: ", id)

	company, err := companyRepo.FindById(id)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.OK(company))
}

func Create(ctx *gin.Context) {
	log.Println("Membuat company baru...")

	// bind json to struct
	request := request.CompanyRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// validate request
	if err := helper.Validator(request); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	// map request to model
	company := models.Company{
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Address: request.Address,
	}

	// store to database
	if err := companyRepo.Create(company); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// give response
	ctx.JSON(http.StatusOK, response.OK(request))
}

func Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
	}
	log.Println("Mengambil company dengan id: ", id)

	// bind json to struct
	req := request.CompanyRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// validate req
	if err := helper.Validator(req); err != nil {
		helper.ThrowFormatInvalid(ctx, err)
		return
	}

	// store to database
	if err := companyRepo.Update(id, req); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// give response
	ctx.JSON(http.StatusOK, response.OK(req))
}

func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.ThrowError(ctx, err)
	}
	log.Println("Mengambil company dengan id: ", id)

	if err := companyRepo.Delete(id); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	// give response
	ctx.JSON(http.StatusOK, response.OK(nil))
}
