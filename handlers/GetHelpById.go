package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetHelpById(c *gin.Context) {

	id := c.Param("id")

	db := database.GetDatabase()

	var help models.Help

	err := db.First(&help, "id = ?", id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find help: " + err.Error(),
		})
		return
	}

	c.JSON(200, help)

}
