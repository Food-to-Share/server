package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)

}
