package routes

import (
	"lifekost/auth-services/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Root check
	r.GET("/", func(c *gin.Context) {
		c.String(200, "LifeKost Auth Service berjalan 🚀")
	})

	// Public endpoints
	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)
	r.POST("/api/auth/refresh", authHandler.RefreshToken)

	// Protected endpoint (assumes JWT already validated at Gateway)
	protected := r.Group("/api/auth")
	{
		protected.GET("/profile", func(c *gin.Context) {
			email := c.GetString("user_email")
			role := c.GetString("user_role")
			userID := c.GetInt64("user_id")

			c.JSON(200, gin.H{
				"message": "Profil user berhasil diakses",
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})
	}

	return r
}
