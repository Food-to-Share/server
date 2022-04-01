package handlers

import (
	"net/http"

	"github.com/Food-to-Share/server/mocks"
	"github.com/gin-gonic/gin"
)

func GetOrg(c *gin.Context) {

	nome := c.Param("name")

	var orgs = mocks.Organizations
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range orgs {
		if a.Name == nome {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
