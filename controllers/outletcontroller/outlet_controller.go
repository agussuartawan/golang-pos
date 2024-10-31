package outletcontroller

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/repositories/outletrepository"
	"github.com/agussuartawan/golang-pos/services/outletservice"
	"github.com/gin-gonic/gin"
	"log"
)

func Create(ctx *gin.Context) {
	log.Println("Membuat outlet baru...")

	var req request.OutletRequest
	if success := helper.DecodeAndValidate(&req, ctx); !success {
		return
	}

	res, err := outletservice.Create(req)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	helper.JSON201(ctx, res)
}

func List(ctx *gin.Context) {
	log.Println("Mengambil data outlet...")

	var param request.OutletParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	var data []response.OutletResponse
	total, err := outletrepository.List(&data, param)
	if err != nil {
		helper.ThrowError(ctx, err)
		return
	}

	helper.JSONPaginate(ctx, param.GetPage(), param.GetLimit(), total, data)
}
