package handlers

import (
	"net/http"

	"github.com/Food-to-Share/server/mocks"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mocks.Users)
}
