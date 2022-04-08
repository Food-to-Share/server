package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllOrgs(c *gin.Context) {
	db := database.GetDatabase()

	var org []models.Organization
	err := db.Find(&org).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list organization: " + err.Error(),
		})
		return
	}
	c.JSON(200, org)
}
