package models

type Dog struct {
	ID   uint `gorm:"primaryKey" json:"dog_id"`
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}
