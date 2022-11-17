package config

import (
	"test/Utils"
	manytomany "test/models/ManyToMany"

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

	db.AutoMigrate(&manytomany.Language{})
	db.AutoMigrate(&manytomany.User{})
	db.AutoMigrate(&manytomany.UserLanguages{})

	// db.Table("worker_recipes").AddForeignKey("worker_id", "workers(id)", "RESTRICT", "CASCADE")

	DB = db
	// return db
}
