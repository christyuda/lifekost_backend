package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"` // Opsional, jika ingin menyertakan username
	jwt.RegisteredClaims
}
