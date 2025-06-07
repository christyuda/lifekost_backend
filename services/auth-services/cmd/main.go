package main

import (
	"log"
	"os"

	"lifekost/auth-services/configs"
	"lifekost/auth-services/internal/handler"
	"lifekost/auth-services/internal/repository"
	"lifekost/auth-services/internal/service"
	"lifekost/auth-services/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("❌ Gagal load file .env:", err)
	}

	// Init DB
	db, err := configs.InitDB()
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}
	defer db.Close()

	log.Println("✅ Berhasil konek ke database lifekost_auth")

	// Init repo, service, handler
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Setup Gin
	r := gin.Default()

	// Public Routes
	r.GET("/", func(c *gin.Context) {
		c.String(200, "LifeKost Auth Service berjalan 🚀")
	})
	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)

	// Protected Route (requires JWT)
	protected := r.Group("/api/auth")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			email := c.GetString("user_email")
			role := c.GetString("user_role")
			c.JSON(200, gin.H{
				"message": "Profil user berhasil diakses",
				"email":   email,
				"role":    role,
			})
		})
	}

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("🚀 Server berjalan di port :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
