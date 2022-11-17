package manytomany

type Language struct {
	Name  string  `gorm:"primaryKey" json:"name"`
	Users []*User `gorm:"many2many:user_languages;"`
}
