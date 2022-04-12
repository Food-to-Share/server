package handlers

import (
	"strconv"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetHelpById(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var help models.Help
	err = db.First(&help, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find help: " + err.Error(),
		})
		return
	}

	c.JSON(200, help)

}
