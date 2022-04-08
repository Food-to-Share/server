package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/Food-to-Share/server/services"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256ENCODER(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
