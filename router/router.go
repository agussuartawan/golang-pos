package router

import (
	companyController "github.com/agussuartawan/golang-pos/controllers/company"
	warehouseController "github.com/agussuartawan/golang-pos/controllers/warehouse"
	"github.com/gin-gonic/gin"
)

var apiRouterV1 *gin.RouterGroup

func LoadRouter() *gin.Engine {
	// init router
	router := Init()

	// define route here
	// grouping route into /api
	apiRouterV1 = router.Group("/api/v1")

	// route for company
	companyRouterV1()
	// route for waehoyse
	warehouseRouterV1()

	return router
}

func companyRouterV1() {
	router := apiRouterV1.Group("/company")
	router.GET("/", companyController.List)
	router.GET("/:id", companyController.FindById)
	router.POST("/", companyController.Create)
	router.PATCH("/:id", companyController.Update)
	router.DELETE("/:id", companyController.Delete)
}

func warehouseRouterV1() {
	router := apiRouterV1.Group("/warehouse")
	router.GET("/", warehouseController.List)
	router.GET("/:id", warehouseController.FindById)
	router.POST("/", warehouseController.Create)
	router.PATCH("/:id", warehouseController.Update)
	router.DELETE("/:id", warehouseController.Delete)
}