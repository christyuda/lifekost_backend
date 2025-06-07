package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Struct untuk permintaan login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct untuk respons login (berisi token JWT)
type LoginResponse struct {
	Token string `json:"token"`
}

// Struct untuk permintaan register user baru
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct respons register (opsional jika ingin menampilkan kembali data)
type RegisterResponse struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Struct untuk representasi user di DB
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// JWTClaims adalah payload yang dikandung di dalam token JWT
type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
