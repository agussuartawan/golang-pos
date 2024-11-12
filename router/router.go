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
	"github.com/agussuartawan/golang-pos/router/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter(baseRouter *gin.Engine) {
	// define route here
	// grouping route into /api
	router := baseRouter.Group("/api/v1")

	company := router.Group("/company")
	{
		company.Use(middleware.Authorized("view_company")).GET("", companycontroller.List)
		company.Use(middleware.Authorized("view_company")).GET(":id", companycontroller.FindById)
		company.Use(middleware.Authorized("create_company")).POST("", companycontroller.Create)
		company.Use(middleware.Authorized("update_company")).PATCH(":id", companycontroller.Update)
		company.Use(middleware.Authorized("delete_company")).DELETE(":id", companycontroller.Delete)
	}

	// route for waehoyse
	warehouse := router.Group("/warehouse")
	{
		warehouse.Use(middleware.Authorized("view_warehouse")).GET("", warehousecontroller.List)
		warehouse.Use(middleware.Authorized("view_warehouse")).GET(":id", warehousecontroller.FindById)
		warehouse.Use(middleware.Authorized("create_warehouse")).POST("", warehousecontroller.Create)
		warehouse.Use(middleware.Authorized("update_warehouse")).PATCH(":id", warehousecontroller.Update)
		warehouse.Use(middleware.Authorized("delete_warehouse")).DELETE(":id", warehousecontroller.Delete)
	}

	// route for role
	role := router.Group("/role")
	{
		role.Use(middleware.Authorized("view_role")).GET("", rolecontroller.List)
		role.Use(middleware.Authorized("view_role")).GET(":id", rolecontroller.FindById)
		role.Use(middleware.Authorized("create_role")).POST("", rolecontroller.Create)
		role.Use(middleware.Authorized("update_role")).PATCH(":id", rolecontroller.Update)
		role.Use(middleware.Authorized("delete_role")).DELETE(":id", rolecontroller.Delete)
		role.Use(middleware.Authorized("update_role")).POST("append-permissions", rolecontroller.AppendPermissions)
	}

	// route for permission
	permission := router.Group("/permission")
	{
		permission.Use(middleware.Authorized("view_permission")).GET("", permissioncontroller.List)
		// permission.GET("/:id", permissioncontroller.FindById)
		permission.Use(middleware.Authorized("create_permission")).POST("", permissioncontroller.Create)
		// permission.PATCH("/:id", permissioncontroller.Update)
		// permission.DELETE("/:id", permissioncontroller.Delete)
	}

	// route for user
	user := router.Group("/user")
	{
		user.Use(middleware.Authorized("view_user")).GET("", usercontroller.List)
		// user.GET("/:id", usercontroller.FindById)
		user.Use(middleware.Authorized("create_user")).POST("", usercontroller.Create)
		user.Use(middleware.Authorized("update_user")).POST("append-roles", usercontroller.AppendRoles)
		// user.PATCH("/:id", usercontroller.Update)
		// user.DELETE("/:id", usercontroller.Delete)
	}

	// route for auth
	auth := router.Group("/auth")
	{
		auth.POST("/login", authcontroller.Login)
		auth.Use(middleware.Authenticated()).GET("profile", authcontroller.Profile)
		auth.Use(middleware.Authenticated()).DELETE("logout", authcontroller.Logout)
	}

	// route for unit
	unit := router.Group("/unit")
	{
		unit.Use(middleware.Authorized("create_unit")).POST("", unitcontroller.Create)
		unit.Use(middleware.Authorized("view_unit")).GET("", unitcontroller.List)
	}

	// outlet
	outlet := router.Group("/outlet")
	{
		outlet.Use(middleware.Authorized("view_outlet")).GET("", outletcontroller.List)
		outlet.Use(middleware.Authorized("create_outlet")).POST("", outletcontroller.Create)
	}

	// product
	product := router.Group("/product")
	{
		product.Use(middleware.Authorized("view_product")).POST("", productcontroller.Create)
		product.Use(middleware.Authorized("create_product")).GET("", productcontroller.List)
	}
}
