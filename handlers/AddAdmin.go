package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/Food-to-Share/server/services"
	"github.com/gin-gonic/gin"
)

func AddAdmin(c *gin.Context) {

	db := database.GetDatabase()

	var admin models.Admin

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	admin.Password = services.SHA256ENCODER(admin.Password)

	err = db.Create(&admin).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
