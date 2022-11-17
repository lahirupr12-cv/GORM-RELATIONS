package onetomany

type CreditCard struct {
	Number     uint `gorm:"primaryKey"`
	UserNumber string
}
