package entities

type Book struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
}
