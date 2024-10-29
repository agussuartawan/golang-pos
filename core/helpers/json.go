package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/agussuartawan/golang-pos/core/errors"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JSON(ctx *gin.Context, response response.Response) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func ThrowError(ctx *gin.Context, err error) {
	switch {
		case err == errors.ErrUnauthorized || strings.Contains(err.Error(), "jwt"):
			JSON401(ctx, err.Error())
		case err == gorm.ErrRecordNotFound || strings.Contains(err.Error(), "not found"):
			JSON404(ctx, err)
		case err == errors.ErrFormatInvalid || strings.Contains(err.Error(), "invalid"):
			JSON400(ctx, err)
		default:
			JSON500(ctx, err)
	}
}

func ThrowFormatInvalid(ctx *gin.Context, errors []response.ValidationFailsResponse) {
	res := response.Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Errors:    errors,
	}
	ctx.JSON(res.Status, res)
}

func JSON500(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, response.Response{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Errors:  err.Error(),
	})
}

func JSON404(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, response.Response{
		Status:  http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Errors:  err.Error(),
	})
}

func JSON400(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, response.Response{
		Status:  http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Errors:  err.Error(),
	})
}

func JSON401(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, response.Response{
		Status:  http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
		Errors:  message,
	})
}

func JSON403(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden, response.Response{
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