package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/agussuartawan/golang-pos/data/request"
	"net/http"
	"strings"

	customErrors "github.com/agussuartawan/golang-pos/core/errors"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JSON200(ctx *gin.Context, res interface{}) {
	ctx.JSON(http.StatusOK, response.OK(res))
}

func JSON201(ctx *gin.Context, res interface{}) {
	ctx.JSON(http.StatusCreated, response.Created(res))
}

func JSONPaginate(ctx *gin.Context, param request.PaginationParam, data interface{}) {
	ctx.JSON(http.StatusOK, response.Response{
		Status:     http.StatusOK,
		Message:    "Success",
		Pagination: param.ToResponse(),
		Data:       data,
	})
}

func ThrowError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, customErrors.ErrUnauthorized) || strings.Contains(err.Error(), "jwt"):
		JSON401(ctx, err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(err.Error(), "not found"):
		JSON404(ctx, err)
	case errors.Is(err, customErrors.ErrFormatInvalid) || strings.Contains(err.Error(), "invalid"):
		JSON400(ctx, err)
	default:
		JSON500(ctx, err)
	}
}

func ThrowFormatInvalid(ctx *gin.Context, errors []response.ValidationFailsResponse) {
	res := response.Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Errors:  errors,
	}
	ctx.AbortWithStatusJSON(res.Status, res)
}

func JSON500(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Errors:  err.Error(),
	})
}

func JSON404(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, response.Response{
		Status:  http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Errors:  err.Error(),
	})
}

func JSON400(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Response{
		Status:  http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Errors:  err.Error(),
	})
}

func JSON401(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
		Status:  http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
		Errors:  message,
	})
}

func JSON403(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, response.Response{
		Status:  http.StatusForbidden,
		Message: http.StatusText(http.StatusForbidden),
		Errors:  message,
	})
}

func LogInfo(value interface{}) {
	jsonData, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
}

func DecodeAndValidate(req interface{}, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ThrowError(ctx, err)
		return false
	}

	if err := Validator(req); err != nil {
		ThrowFormatInvalid(ctx, err)
		return false
	}

	return true
}

func GetQueryParam(param interface{}, ctx *gin.Context) bool {
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ThrowError(ctx, err)
		return false
	}

	return true
}
