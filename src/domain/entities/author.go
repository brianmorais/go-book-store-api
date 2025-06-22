package entities

type Author struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Books []Book `json:"books" gorm:"foreignKey:AuthorId;references:Id"`
}
