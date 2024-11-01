package router

import (
	"github.com/agussuartawan/golang-pos/controllers/authcontroller"
	"github.com/agussuartawan/golang-pos/controllers/companycontroller"
	"github.com/agussuartawan/golang-pos/controllers/outletcontroller"
	"github.com/agussuartawan/golang-pos/controllers/permissioncontroller"
	"github.com/agussuartawan/golang-pos/controllers/productcontroller"
	"github.com/agussuartawan/golang-pos/controllers/rolecontroller"
	"github.com/agussuartawan/golang-pos/controllers/unitcontroller"
	"github.com/agussuartawan/golang-pos/controllers/usercontroller"
	"github.com/agussuartawan/golang-pos/controllers/warehousecontroller"
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
	// route for unit
	unitRouterV1()
	// outlet
	outletRouterV1()
	// product
	productRouterV1()

	return router
}

func companyRouterV1() {
	router := apiRouterV1.Group("/company")
	router.Use(middleware.Authorized("view_company")).GET("/", companycontroller.List)
	router.Use(middleware.Authorized("view_company")).GET("/:id", companycontroller.FindById)
	router.Use(middleware.Authorized("create_company")).POST("/", companycontroller.Create)
	router.Use(middleware.Authorized("update_company")).PATCH("/:id", companycontroller.Update)
	router.Use(middleware.Authorized("delete_company")).DELETE("/:id", companycontroller.Delete)
}

func warehouseRouterV1() {
	router := apiRouterV1.Group("/warehouse")
	router.Use(middleware.Authorized("view_warehouse")).GET("/", warehousecontroller.List)
	router.Use(middleware.Authorized("view_warehouse")).GET("/:id", warehousecontroller.FindById)
	router.Use(middleware.Authorized("create_warehouse")).POST("/", warehousecontroller.Create)
	router.Use(middleware.Authorized("update_warehouse")).PATCH("/:id", warehousecontroller.Update)
	router.Use(middleware.Authorized("delete_warehouse")).DELETE("/:id", warehousecontroller.Delete)
}

func roleRouterV1() {
	router := apiRouterV1.Group("/role")
	router.Use(middleware.Authorized("view_role")).GET("/", rolecontroller.List)
	router.Use(middleware.Authorized("view_role")).GET("/:id", rolecontroller.FindById)
	router.Use(middleware.Authorized("create_role")).POST("/", rolecontroller.Create)
	router.Use(middleware.Authorized("update_role")).PATCH("/:id", rolecontroller.Update)
	router.Use(middleware.Authorized("delete_role")).DELETE("/:id", rolecontroller.Delete)
	router.Use(middleware.Authorized("update_role")).POST("/append-permissions", rolecontroller.AppendPermissions)
}

func permissionRouterV1() {
	router := apiRouterV1.Group("/permission")
	router.Use(middleware.Authorized("view_permission")).GET("/", permissioncontroller.List)
	// router.GET("/:id", permissioncontroller.FindById)
	router.Use(middleware.Authorized("create_permission")).POST("/", permissioncontroller.Create)
	// router.PATCH("/:id", permissioncontroller.Update)
	// router.DELETE("/:id", permissioncontroller.Delete)
}

func userRouterV1() {
	router := apiRouterV1.Group("/user")
	router.Use(middleware.Authorized("view_user")).GET("/", usercontroller.List)
	// router.GET("/:id", usercontroller.FindById)
	router.Use(middleware.Authorized("create_user")).POST("/", usercontroller.Create)
	router.Use(middleware.Authorized("update_user")).POST("/append-roles", usercontroller.AppendRoles)
	// router.PATCH("/:id", usercontroller.Update)
	// router.DELETE("/:id", usercontroller.Delete)
}

func authRouterV1() {
	router := apiRouterV1.Group("/auth")
	router.POST("/login", authcontroller.Login)
	router.Use(middleware.Authenticated()).GET("/profile", authcontroller.Profile)
}

func unitRouterV1() {
	router := apiRouterV1.Group("/unit")
	router.Use(middleware.Authorized("create_unit")).POST("/", unitcontroller.Create)
	router.Use(middleware.Authorized("view_unit")).GET("/", unitcontroller.List)
}

func outletRouterV1() {
	router := apiRouterV1.Group("/outlet")
	router.Use(middleware.Authorized("view_outlet")).GET("/", outletcontroller.List)
	router.Use(middleware.Authorized("create_outlet")).POST("/", outletcontroller.Create)
}

func productRouterV1() {
	router := apiRouterV1.Group("/product")
	router.Use(middleware.Authorized("view_product")).POST("/", productcontroller.Create)
	router.Use(middleware.Authorized("create_product")).GET("/", productcontroller.List)
}
