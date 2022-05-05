package models

type Admin struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Organization string `json:"organization"`
}
