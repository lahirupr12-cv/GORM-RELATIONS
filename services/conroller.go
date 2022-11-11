package services

import (
	"test/config"
	"test/models"
)

func Create() {
	config.DB.Create(&models.Dog{Name: "dog1", Toys: models.Toy{Name: "toy2"}})
}
