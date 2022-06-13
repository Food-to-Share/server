package routes

import (
	"github.com/Food-to-Share/server/handlers"
	"github.com/Food-to-Share/server/server/middlewares"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	main := router.Group("api/v1")
	{
		User := main.Group("users")
		{
			User.GET("/:id", handlers.GetUserById)
			User.GET("/", handlers.GetAllUsers)
			User.POST("/", handlers.AddUser)
			User.DELETE("/:id", handlers.DelUser)
			User.PUT("/", handlers.UpdateUser)
			User.POST("/verifyUserNiss", handlers.VerifyUserNiss)
		}
		Organization := main.Group("orgs", middlewares.Auth())
		{
			Organization.GET("/:id", handlers.GetOrgById)
			Organization.GET("/", handlers.GetAllOrgs)
			Organization.POST("/", handlers.AddOrg)
			Organization.DELETE("/:id", handlers.DelOrg)
			Organization.PUT("/", handlers.UpdateOrg)
		}
		Admin := main.Group("admins")
		{
			Admin.GET("/:id", handlers.GetAdminById)
			Admin.GET("/", handlers.GetAllAdmins)
			Admin.POST("/", handlers.AddAdmin)
			Admin.DELETE("/:id", handlers.DelAdmin)
			Admin.PUT("/", handlers.UpdateAdmin)
		}
		Help := main.Group("helps")
		{
			Help.GET("/:id", handlers.GetHelpById)
			Help.GET("/", handlers.GetAllHelps)
			Help.POST("/", handlers.AddHelp)
			Help.DELETE("/:id", handlers.DelHelp)
			Help.PUT("/", handlers.UpdateHelp)
		}
		main.POST("login", handlers.Login)
	}
	return router
}