package handlers

import (
	"strconv"

	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var user models.User
	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)

}
