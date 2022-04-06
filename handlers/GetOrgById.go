package handlers

// import (
// 	"net/http"

// 	"github.com/Food-to-Share/server/models"
// 	"github.com/gin-gonic/gin"
// )

// func GetOrgById(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range models.Organization {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "org not found"})
// }
