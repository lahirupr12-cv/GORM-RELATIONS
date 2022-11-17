package manytomany

type User struct {
	Id        uint        `gorm:"primaryKey" json:"id"`
	Languages []*Language `gorm:"many2many:user_languages;"`
}
