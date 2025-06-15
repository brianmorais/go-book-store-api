package entities

import "time"

type Order struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id"`
	BookIDs   []int     `json:"book_ids"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
