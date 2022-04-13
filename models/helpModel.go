package models

type Help struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Lunch     string `json:"lunch"`
	Breakfast string `json:"breakfast"`
	Dinner    string `json:"dinner"`
}
