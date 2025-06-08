package main

import (
	"lifekost/gateway/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables (AUTH_SERVICE_URL, etc.)
	if err := godotenv.Load("internal/configs/.env"); err != nil {
		log.Fatal("❌ Gagal load file .env:", err)
	}

	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	// Init Gin
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		routes.AuthRoutes(api) // setup /api/auth/*
	}

	// Run server
	log.Printf("🚀 Gateway berjalan di port :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Gagal menjalankan Gateway:", err)
	}
}
