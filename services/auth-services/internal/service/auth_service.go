package service

import (
	"errors"
	"lifekost/auth-services/internal/repository"
	"lifekost/auth-services/pkg/domain"
	"lifekost/libs/auth"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthService adalah interface layanan otentikasi
type AuthService interface {
	Register(req domain.RegisterRequest) (*domain.User, error)
	Login(req domain.LoginRequest) (*domain.LoginResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(req domain.RegisterRequest) (*domain.User, error) {
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Role:      "user",
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authService) Login(req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	// Buat klaim token JWT
	claims := auth.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}

	// Generate access token
	accessToken, err := auth.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := auth.GenerateRefreshToken(claims)
	if err != nil {
		return nil, err
	}

	// Return response
	return &domain.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour), // 1 hari
		CreatedAt:    time.Now(),
	}, nil

}
