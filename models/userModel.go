package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  int    `json:"contact"`
	Password string `json:"password"`
	Address  string `json:"adress"`
}
