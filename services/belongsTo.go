package services

import (
	"test/config"
	"test/models"
)

func CreateTable() {
	t1 := models.Town{
		Name: "Pune",
	}
	// t2 := models.Town{
	// 	Name: "Mumbai",
	// }
	t3 := models.Town{
		Name: "Hyderabad",
	}

	p1 := models.Place{
		Name: "Katraj",
		Town: t1,
	}
	p2 := models.Place{
		Name: "Thane",
		Town: models.Town{
			Name: "aaaa",
		},
	}
	p3 := models.Place{
		Name: "Secundarabad",
		Town: t3,
	}
	config.DB.Create(&p1) //Saving one to one relationship
	config.DB.Create(&p2)
	config.DB.Create(&p3)
}

// func CreateUser() {
// 	// towns := models.Town{}
// 	// data := config.DB.Where("id=?", 0).Find(towns)
// 	// // p3 := models.Place{
// 	// // 	Name: "Secundarabad",
// 	// // 	Town: data.GenerateModel("users"),
// 	// // }
// 	// fmt.Println(data)
// }
