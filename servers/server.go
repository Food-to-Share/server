package servers

import (
	"log"

	"github.com/Food-to-Share/server/handlers"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/organizations", handlers.GetAllOrgs)

	log.Println("API is running! Teste")
	router.Run(":4000")
}
