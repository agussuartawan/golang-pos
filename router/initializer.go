package router

import (
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// load router
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

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
			Error:   &response.Error{Message: "Resource Not Found"},
		}
		c.JSON(http.StatusNotFound, res)
	})

	LoadRouter(router)

	return router
}
