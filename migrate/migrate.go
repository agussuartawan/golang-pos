package main

import (
	"github.com/agussuartawan/golang-pos/config"
	"github.com/agussuartawan/golang-pos/models"
)

func init() {
	config.ConnectToDatabase()
}

func main() {
	config.DB.AutoMigrate(
		&models.Company{},
		&models.Warehouse{},
		&models.User{},
		&models.Role{},
		&models.Session{},
		&models.Permission{},
	)
}