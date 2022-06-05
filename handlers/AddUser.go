package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Jid         string `json:"jid"`
	DisplayName string `json:"displayName"`
}

func AddUser(c *gin.Context) {

	db := database.GetDatabase()
	httpposturl := "http://localhost:29318/_matrix/provision/v1/startUser"

	var user models.User
	var help models.Help

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	errHelp := c.ShouldBindJSON(&help)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + errHelp.Error(),
		})
		return
	}

	user.ID = uuid.NewString()
	help.ID = user.ID

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	errHelp = db.Create(&help).Error

	if errHelp != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create help: " + errHelp.Error(),
		})
		return
	}

	newUsers := NewUser{
		Jid:         user.ID,
		DisplayName: user.Name,
	}

	data, err := json.Marshal(newUsers)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)

	request, errorPost := http.NewRequest("POST", httpposturl, reader)
	if errorPost != nil {
		panic(errorPost)
	}
	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot do request: " + error.Error(),
		})
		return
	}
	defer response.Body.Close()

	c.Status(204)

}
