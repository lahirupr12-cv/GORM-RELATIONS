package onetomany

type User struct {
	Id           string `gorm:"primaryKey"`
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignKey:Number;references:MemberNumber"`
}
