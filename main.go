package main

import (
	"test/config"
	"test/services"
)

func main() {
	config.ConnectionDB()
	services.Create()
}