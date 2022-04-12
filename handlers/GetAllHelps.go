package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllHelps(c *gin.Context) {
	db := database.GetDatabase()

	var help []models.Help
	err := db.Find(&help).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list help: " + err.Error(),
		})
		return
	}
	c.JSON(200, help)
}
