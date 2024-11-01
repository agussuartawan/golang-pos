package productcontroller

import (
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/repositories/productrepository"
	"github.com/agussuartawan/golang-pos/services/productservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req request.ProductRequest
	if success := helper.DecodeAndValidate(&req, c); !success {
		return
	}

	id, err := productservice.Create(req)
	if err != nil {
		helper.ThrowError(c, err)
		return
	}

	helper.JSON201(c, response.IDResponse{ID: *id})
}

func List(c *gin.Context) {
	var param request.ProductParam
	if err := c.ShouldBindQuery(&param); err != nil {
		helper.ThrowError(c, err)
		return
	}

	var data []response.ProductResponse
	if err := productrepository.List(&data, &param); err != nil {
		helper.ThrowError(c, err)
		return
	}

	helper.JSONPaginate(c, param.PaginationParam.SetData(data))
}
