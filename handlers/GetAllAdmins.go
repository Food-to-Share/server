package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllAdmins(c *gin.Context) {
	db := database.GetDatabase()

	var admin []models.Admin
	err := db.Find(&admin).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list user: " + err.Error(),
		})
		return
	}
	c.JSON(200, admin)
}
