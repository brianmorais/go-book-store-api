package entities

import "time"

type Order struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"userId"`
	Books     []Book    `json:"books" gorm:"one2many:Book;foreignKey:Id;references:Id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
