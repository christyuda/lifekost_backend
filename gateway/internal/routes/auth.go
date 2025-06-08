package routes

import (
	"lifekost/gateway/middleware"
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authURL := os.Getenv("AUTH_SERVICE_URL")
	target, err := url.Parse(authURL)
	if err != nil {
		log.Fatalf("❌ Invalid AUTH_SERVICE_URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Public routes
	r.POST("/auth/register", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/auth/login", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/auth/refresh", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	// Protected routes with JWT middleware
	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}
}
