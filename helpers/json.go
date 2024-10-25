package helper

import (
	"net/http"
	"strings"

	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JSON(ctx *gin.Context, response response.Response) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func ThrowError(ctx *gin.Context, err error) {
	switch {
		case err == gorm.ErrRecordNotFound || strings.Contains(err.Error(), "not found"):
			JSON404(ctx, err)
		case err == errors.ErrFormatInvalid:
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