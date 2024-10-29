package main

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/router"
)

func init() {
	config.ConnectToDatabase()
}

func main() {
	r := router.LoadRouter()
	r.Run(":8080")
}