package main

import (
	"log"
	"os"

	"lifekost/auth-services/configs"
	"lifekost/auth-services/internal/handler"
	"lifekost/auth-services/internal/repository"
	"lifekost/auth-services/internal/routes"
	"lifekost/auth-services/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("❌ Gagal load file .env:", err)
	}

	// Init DB
	db, err := configs.InitDB()
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}
	defer db.Close()

	log.Println("✅ Berhasil konek ke database lifekost_auth")

	// Init DI
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Init router
	router := routes.SetupRouter(authHandler)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("🚀 Server berjalan di port :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
