package servers

import (
	"log"

	"github.com/Food-to-Share/server/handlers"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/organizations", handlers.GetAllOrgs)
	router.GET("/organizations/:name", handlers.GetOrg)
	router.POST("/organizations", handlers.AddOrg)

	log.Println("API is running!")
	router.Run(":4000")
}
