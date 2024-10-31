package unitcontroller

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/unitrepository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(ctx *gin.Context) {
	log.Println("Membuat unit baru...")

	req := request.UnitRequest{}
	if success := helper.DecodeAndValidate(&req, ctx); !success {
		return
	}

	// map req to model
	model := models.Unit{
		Name:      req.Name,
		BaseValue: req.BaseValue,
	}
	if err := unitrepository.Create(&model); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.OK(response.UnitResponse{
		Id:        model.ID,
		Name:      model.Name,
		BaseValue: model.BaseValue,
	}))
}

func List(ctx *gin.Context) {
	log.Println("Mengambil list unit...")

	var units []response.UnitResponse
	if err := unitrepository.List(&units); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	helper.JSON200(ctx, units)
}
