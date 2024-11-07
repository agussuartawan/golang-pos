package router

import (
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Init() *gin.Engine {
	// load router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Next()
	})

	router.Use(func(c *gin.Context) {
		log.Println("CORS middleware executed")
		c.Next()
	})

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

	return router
}
