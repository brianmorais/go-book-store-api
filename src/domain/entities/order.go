package entities

import "time"

type Order struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"userId"`
	Books     []Book    `json:"books" gorm:"many2many:order_books;"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
