package mocks

import "github.com/Food-to-Share/server/models"

var Users = []models.User{
	{ID: 1, Name: "Bruno", Username: "bmiguel", Email: "bmiguel@gmail.com", Contact: 937765982, Password: ""},
	{ID: 2, Name: "Kevin", Username: "kevind", Email: "kevind@gmail.com", Contact: 937765982, Password: ""},
	{ID: 3, Name: "Goncalo", Username: "gonaf", Email: "gonaf@gmail.com", Contact: 937765982, Password: ""},
}


// type User struct {
// 	ID       int    `json:"id"`
// 	Username int    `json:"username"`
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Contact  string `json:"contact"`
// 	Password string `json:"password"`
// 	Address  string `json:"adress"`
// }
