package models

type Organization struct {
	ID             int    `json:"id"`
	Name           string `json:"title"`
	Address        string `json:"address"`
	SocialSecurity string `json:"socialSec"`
}
