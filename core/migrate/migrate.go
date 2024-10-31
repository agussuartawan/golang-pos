package main

import (
	"github.com/agussuartawan/golang-pos/core/config"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/models"
)

func init() {
	config.ConnectToDatabase()
}

func main() {
	if err := config.DB.AutoMigrate(
		&models.Company{},
		&models.Category{},
		&models.Warehouse{},
		&models.Outlet{},
		&models.Supplier{},
		&models.User{},
		&models.Role{},
		&models.Session{},
		&models.Permission{},
		&models.Employee{},
		&models.Unit{},
		&models.Product{},
		&models.ProductPrice{},
		&models.ProductStock{},
	); err != nil {
		helper.LogError(err)
	}

	models.RunSeeders()
}
