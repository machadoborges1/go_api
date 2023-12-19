package main

import (
	"github.com/machadoborges1/go_api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBPort)
}
