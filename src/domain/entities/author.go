package entities

type Author struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
