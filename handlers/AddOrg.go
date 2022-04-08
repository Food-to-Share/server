package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func AddOrg(c *gin.Context) {

	db := database.GetDatabase()

	var org models.Organization

	err := c.ShouldBindJSON(&org)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&org).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create org: " + err.Error(),
		})
		return
	}

	c.JSON(200, org)

}
