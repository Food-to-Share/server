package handlers

import (
	"net/http"

	"github.com/Food-to-Share/server/mocks"
	"github.com/gin-gonic/gin"
)

func GetAllOrgs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Organizations)
}
