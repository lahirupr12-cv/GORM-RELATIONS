package models

type Toy struct {
	ID        uint `gorm:"primaryKey" json:"toy_id"`
	Name      string
	OwnerID   int
	OwnerType string
}
