package models

type Primitive struct {
	Id         uint   `gorm:"primaryKey" json:"ptimitive_id"`
	Name       string `json:"name" gorm:"index:idx_name,unique"`
	Category   Category
	CategoryId int `gorm:"ForeignKey:id"` //this foreignKey tag didn't works
}
