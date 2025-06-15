package entities

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}
