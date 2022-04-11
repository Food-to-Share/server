package middlewares

import (
	"github.com/Food-to-Share/server/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}
		token := header[len(bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
