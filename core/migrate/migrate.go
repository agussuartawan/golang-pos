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
	err := config.DB.AutoMigrate(
		&models.Company{},
		&models.Warehouse{},
		&models.User{},
		&models.Role{},
		&models.Session{},
		&models.Permission{},
	)
	if err != nil {
		helper.LogError(err)
	}
}
