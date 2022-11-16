package config

import (
	"test/Utils"
	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// create DB connection
func ConnectionDB() {
	dsn := "host=localhost user=postgres password=lahiru12 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//call checkError method to check wether thers error
	Utils.CheckError(err)

	//drop table if table already exists
	// db.Migrator().DropTable(&models.Primitive{})
	// db.Migrator().DropTable(&models.Field{})

	// db.AutoMigrate(&models.Town{})
	// db.AutoMigrate(&models.Place{})

	// db.AutoMigrate(&models.Dog{})
	// db.AutoMigrate(&models.Toy{})
	db.AutoMigrate(&models.Category{})

	DB = db
	// return db
}
