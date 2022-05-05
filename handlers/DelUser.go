package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func DelUser(c *gin.Context) {

	id := c.Param("id")

	db := database.GetDatabase()

	err := db.Delete(&models.User{}, "id = ?", id).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete user: " + err.Error(),
		})
		return
	}
	c.Status(204)
}
