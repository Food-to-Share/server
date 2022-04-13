package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetOrgById(c *gin.Context) {

	id := c.Param("id")

	db := database.GetDatabase()

	var org models.Organization
	err := db.First(&org, "id = ?", id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find organization: " + err.Error(),
		})
		return
	}

	c.JSON(200, org)

}
