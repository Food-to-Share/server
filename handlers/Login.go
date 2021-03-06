package handlers

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/Food-to-Share/server/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	var admin models.Admin
	dbError := db.Where("email = ?", p.Email).First(&admin).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find admin",
		})
		return
	}

	if admin.Password != services.SHA256ENCODER(p.Password) {
		c.JSON(401, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(admin.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}
