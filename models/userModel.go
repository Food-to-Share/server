package models

type User struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Contact      int    `json:"contact"`
	Address      string `json:"address"`
	Organization string `json:"organization"`
}
