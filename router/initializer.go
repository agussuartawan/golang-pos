package router

import (
	"net/http"

	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// load router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		res := response.Response{
			Status:  http.StatusOK,
			Message: "Welcome!",
			Data:    nil,
		}
		c.JSON(http.StatusOK, res)
	})

	router.NoRoute(func(c *gin.Context) {
		res := response.Response{
			Status:  http.StatusNotFound,
			Message: "Not Found",
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, res)
	})

	return router
}