package models

type Organization struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	SocialSecurity string `json:"socialsecurity"`
}
