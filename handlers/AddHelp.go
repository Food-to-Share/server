package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func AddHelp(c *gin.Context) {

	db := database.GetDatabase()

	var help models.Help

	err := c.ShouldBindJSON(&help)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
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
