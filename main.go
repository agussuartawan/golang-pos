package main

import (
	"github.com/agussuartawan/golang-pos/core/config"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"github.com/agussuartawan/golang-pos/router"
)

func init() {
	config.ConnectToDatabase()
}

func main() {
	r := router.LoadRouter()
	err := r.Run(":8080")
	if err != nil {
		helper.LogError(err)
	}
}
