package middleware

import (
	"lifekost/libs/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware memverifikasi token JWT sebelum melanjutkan ke handler
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil Authorization header: "Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token salah"})
			c.Abort()
			return
		}

		// Verifikasi token
		claims, err := auth.VerifyToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Simpan data user ke context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}
