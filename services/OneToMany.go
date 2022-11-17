package services

// import (
// 	"test/config"
// 	onetomany "test/models/OneToMany"
// )

// func CreditCard() {
// 	T1 := onetomany.CreditCard{
// 		UserNumber: "tag1",
// 	}
// 	T2 := onetomany.CreditCard{
// 		UserNumber: "tag2",
// 	}
// 	T3 := onetomany.CreditCard{
// 		UserNumber: "tag3",
// 	}
// 	T4 := onetomany.CreditCard{
// 		UserNumber: "tag4",
// 	}
// 	config.DB.Create(&T1)
// 	config.DB.Create(&T2)
// 	config.DB.Create(&T3)
// 	config.DB.Create(&T4)
// }

// func CreateUsers() {
// 	// var card onetomany.CreditCard
// 	// config.DB.First(&card, "user_number = ?", "tag1")

// 	// fmt.Println(card)

// 	T1 := onetomany.CreditCard{
// 		UserNumber: "tag1",
// 	}
// 	pt1 := onetomany.User{
// 		CreditCards: []onetomany.CreditCard{T1},
// 	}
// 	config.DB.Create(&pt1)
// }
