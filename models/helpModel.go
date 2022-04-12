package models

type Help struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	HelpType  string `json:"helpType"`
	CheckHelp string `json:"checkHelp"`
}
