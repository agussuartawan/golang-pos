package main

import (
	"github.com/agussuartawan/golang-pos/config"
	"github.com/agussuartawan/golang-pos/router"
)

func init() {
	config.ConnectToDatabase()
}

func main() {
	r := router.LoadRouter()
	r.Run()
}