package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func DelHelp(c *gin.Context) {

	id := c.Param("id")

	db := database.GetDatabase()

	err := db.Delete(&models.Help{}, "id = ?", id).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete help: " + err.Error(),
		})
		return
	}
	c.Status(204)
}
