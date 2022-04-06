package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/Food-to-Share/server/mocks"
// 	"github.com/gin-gonic/gin"
// )

// func DelOrg(c *gin.Context) {
// 	id := c.Param("id")
// 	strconv.Atoi(id)

// 	var orgs = mocks.Organizations

// 	for i, a := range orgs {
// 		if a.ID == id {
// 			orgs = append(orgs[:i], orgs[i+1:]...)
// 			break
// 		}
// 	}

// 	c.Status(http.StatusNoContent)
// }
