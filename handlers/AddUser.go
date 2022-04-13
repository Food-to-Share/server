package handlers

import (
	"github.com/google/uuid"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	db := database.GetDatabase()

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

	err = db.Create(&help).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create help: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
