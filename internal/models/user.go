package models

type User struct {
	Email    string `json:"email" gorm:"unqie"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
