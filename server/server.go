package server

import (
	"log"

	"github.com/Food-to-Share/server/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}

}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	log.Print("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
