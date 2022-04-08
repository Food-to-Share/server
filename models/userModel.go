package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  int    `json:"contact"`
	Password string `json:"password"`
	Address  string `json:"adress"`
}
