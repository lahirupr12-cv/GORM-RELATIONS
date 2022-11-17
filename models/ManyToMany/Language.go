package manytomany

type Language struct {
	Id    uint `gorm:"primaryKey" json:"id"`
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}
