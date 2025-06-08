package domain

import (
	"time"
)

// Struct untuk permintaan login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct untuk respons login (berisi token JWT)
type LoginResponse struct {
	Token        string    `json:"token"`         // Token JWT yang
	RefreshToken string    `json:"refresh_token"` // Token refresh untuk mendapatkan token baru
	ExpiresAt    time.Time `json:"expires_at"`    // Waktu kedaluwarsa token
	CreatedAt    time.Time `json:"created_at"`    // Waktu pembuatan token
}

// Struct untuk permintaan register user baru
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"` // Misalnya: "user", "admin"
	Password string `json:"password"`
}

// Struct respons register (opsional jika ingin menampilkan kembali data)
type RegisterResponse struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
