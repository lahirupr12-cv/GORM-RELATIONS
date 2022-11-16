package services

import (
	"test/config"
	"test/models"
)

func CreateCategory() {
	t1 := models.Category{
		Name:   "Medical",
		Colour: "#5B1865",
		Icon:   "https://cowrie-dev-public-s3-bucket-temp.s3.amazonaws.com/category-icons/medical2x.png",
	}
	t2 := models.Category{
		Name:   "Employment",
		Colour: "#5D675c",
		Icon:   "https://cowrie-dev-public-s3-bucket-temp.s3.amazonaws.com/category-icons/employment2x.png",
	}
	t3 := models.Category{
		Name:   "Personal Information",
		Colour: "#C52233",
		Icon:   "https://cowrie-dev-public-s3-bucket-temp.s3.amazonaws.com/category-icons/personal2x.png",
	}
	t4 := models.Category{
		Name:   "Education",
		Colour: "#BCD8C1",
		Icon:   "https://cowrie-dev-public-s3-bucket-temp.s3.amazonaws.com/category-icons/education2x.png",
	}
	t5 := models.Category{
		Name:   "Financial",
		Colour: "#5688C7",
		Icon:   "https://cowrie-dev-public-s3-bucket-temp.s3.amazonaws.com/category-icons/financial2x.png",
	}
	config.DB.Create(&t1)
	config.DB.Create(&t2)
	config.DB.Create(&t3)
	config.DB.Create(&t4)
	config.DB.Create(&t5)
}
