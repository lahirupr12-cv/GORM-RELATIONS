package services

import (
	"test/config"
	manytomany "test/models/ManyToMany"
)

func CreateLanguage() {
	l1 := manytomany.Language{
		Name: "Lahiru",
	}
	l2 := manytomany.Language{
		Name: "Bandara",
	}
	config.DB.Save(&l1)
	config.DB.Save(&l2)

	u1 := manytomany.User{
		Languages: []*manytomany.Language{&l1, &l2},
	}
	config.DB.Save(&u1)
}
