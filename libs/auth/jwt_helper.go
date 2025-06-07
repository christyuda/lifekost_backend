package auth

import (
	"os"
	"time"

	"lifekost/auth-services/pkg/domain"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken membuat JWT berdasarkan JWTClaims
func GenerateToken(claims domain.JWTClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyToken memverifikasi token JWT dan mengembalikan JWTClaims
func VerifyToken(tokenStr string) (*domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
