package handlers

import (
	"math/rand"
	"net/http"

	"github.com/Food-to-Share/server/mocks"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	newUser.ID = rand.Intn(1000)
	mocks.Users = append(mocks.Users, newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{"Message": "User criado com sucesso!"})
}
