package models

type Place struct {
	ID     int `gorm:primary_key`
	Name   string
	Town   Town
	TownId int `gorm:"ForeignKey:id"` //this foreignKey tag didn't works
}

type Town struct {
	ID   int `gorm:"primary_key"`
	Name string
}
