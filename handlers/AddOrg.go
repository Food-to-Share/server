package handlers

import (
	"math/rand"

	"github.com/Food-to-Share/server/mocks"
	"github.com/Food-to-Share/server/models"
	"github.com/gin-gonic/gin"
)

func AddOrg(c *gin.Context)  {
	var newOrg models.Organization

	if err := c.BindJSON(&newOrg); err != nil {
		return
	}

	newOrg.ID = rand.Intn(1000)
	mocks.Organizations = append(mocks.Organizations, newOrg)
	
}