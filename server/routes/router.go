package routes

import (
	"github.com/Food-to-Share/server/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		User := main.Group("users")
		{
			User.GET("/:id", handlers.GetUserById)
			User.GET("/", handlers.GetAllUsers)
			User.POST("/", handlers.AddUser)
		}
	}
	return router
}

// 	router.GET("/organizations", handlers.GetAllOrgs)
// 	router.GET("/organizations/:name", handlers.GetOrg)
// 	router.POST("/organizations", handlers.AddOrg)
// 	// router.DELETE("/organizations", handlers.DelOrg)
// 	router.POST("/users", handlers.AddUser)
// 	router.GET("/users", handlers.GetAllUsers)
// 	router.GET("/users/:id", handlers.GetUserById)
