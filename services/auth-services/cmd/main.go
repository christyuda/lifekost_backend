package main

import (
	"log"
	"net/http"

	"github.com/christyuda/lifekost_backend/services/auth-service/configs"
	"github.com/joho/godotenv"
)

func main() {
	// Load konfigurasi dari file .env
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("❌ Gagal load file .env:", err)
	}

	// Inisialisasi koneksi ke database PostgreSQL
	db, err := configs.InitDB()
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database:", err)
	}
	defer db.Close()

	log.Println("✅ Berhasil konek ke database lifekost_auth")

	// Jalankan HTTP server default (sementara)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LifeKost Auth Service berjalan 🚀"))
	})

	log.Println("🚀 Server berjalan di port :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
