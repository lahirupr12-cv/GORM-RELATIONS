package services

import (
	"fmt"
	"test/config"
	"test/models"
)

func CreatePrimitives() {
	var cat models.Category
	config.DB.First(&cat, "name = ?", "Medical")
	fmt.Println(cat)
	p1 := models.Primitive{
		Name: "p1",
		Category: models.Category{
			Name: "Medical",
		},
	}
	config.DB.Create(&p1)
}
