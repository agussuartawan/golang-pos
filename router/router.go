package router

import (
	"github.com/agussuartawan/golang-pos/controllers/authController"
	"github.com/agussuartawan/golang-pos/controllers/companyController"
	"github.com/agussuartawan/golang-pos/controllers/permissionController"
	"github.com/agussuartawan/golang-pos/controllers/roleController"
	"github.com/agussuartawan/golang-pos/controllers/userController"
	"github.com/agussuartawan/golang-pos/controllers/warehouseController"
	"github.com/agussuartawan/golang-pos/core/middleware"
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
	// route for role
	roleRouterV1()
	// route for permission
	permissionRouterV1()
	// route for user
	userRouterV1()
	// route for auth
	authRouterV1()

	return router
}

func companyRouterV1() {
	router := apiRouterV1.Group("/company")
	router.Use(middleware.Authorized("view_company")).GET("/", companyController.List)
	router.Use(middleware.Authorized("view_company")).GET("/:id", companyController.FindById)
	router.Use(middleware.Authorized("create_company")).POST("/", companyController.Create)
	router.Use(middleware.Authorized("update_company")).PATCH("/:id", companyController.Update)
	router.Use(middleware.Authorized("delete_company")).DELETE("/:id", companyController.Delete)
}

func warehouseRouterV1() {
	router := apiRouterV1.Group("/warehouse")
	router.Use(middleware.Authorized("view_warehouse")).GET("/", warehouseController.List)
	router.Use(middleware.Authorized("view_warehouse")).GET("/:id", warehouseController.FindById)
	router.Use(middleware.Authorized("create_warehouse")).POST("/", warehouseController.Create)
	router.Use(middleware.Authorized("update_warehouse")).PATCH("/:id", warehouseController.Update)
	router.Use(middleware.Authorized("delete_warehouse")).DELETE("/:id", warehouseController.Delete)
}

func roleRouterV1() {
	router := apiRouterV1.Group("/role")
	router.Use(middleware.Authorized("view_role")).GET("/", roleController.List)
	router.Use(middleware.Authorized("view_role")).GET("/:id", roleController.FindById)
	router.Use(middleware.Authorized("create_role")).POST("/", roleController.Create)
	router.Use(middleware.Authorized("update_role")).PATCH("/:id", roleController.Update)
	router.Use(middleware.Authorized("delete_role")).DELETE("/:id", roleController.Delete)
	router.Use(middleware.Authorized("update_role")).POST("/append-permissions", roleController.AppendPermissions)
}

func permissionRouterV1() {
	router := apiRouterV1.Group("/permission")
	router.Use(middleware.Authorized("view_permission")).GET("/", permissionController.List)
	// router.GET("/:id", permissionController.FindById)
	router.Use(middleware.Authorized("create_permission")).POST("/", permissionController.Create)
	// router.PATCH("/:id", permissionController.Update)
	// router.DELETE("/:id", permissionController.Delete)
}

func userRouterV1() {
	router := apiRouterV1.Group("/user")
	router.Use(middleware.Authorized("view_user")).GET("/", userController.List)
	// router.GET("/:id", userController.FindById)
	router.Use(middleware.Authorized("create_user")).POST("/", userController.Create)
	router.Use(middleware.Authorized("update_user")).POST("/append-roles", userController.AppendRoles)
	// router.PATCH("/:id", userController.Update)
	// router.DELETE("/:id", userController.Delete)
}

func authRouterV1() {
	router := apiRouterV1.Group("/auth")
	router.POST("/login", authController.Login)
	router.Use(middleware.Authenticated()).GET("/profile", authController.Profile)
}