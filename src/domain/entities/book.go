package entities

type Book struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	AuthorId    int     `json:"authorId" gorm:"foreignKey:Id"`
}
