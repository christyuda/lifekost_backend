package service

import (
	"errors"
	"time"

	"github.com/christyuda/lifekost_backend/libs/auth" // JWT helper dari libs/auth
	"github.com/christyuda/lifekost_backend/services/auth-service/internal/domain"
	"github.com/christyuda/lifekost_backend/services/auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// AuthService adalah interface layanan otentikasi
type AuthService interface {
	Register(req domain.RegisterRequest) (*domain.User, error)
	Login(req domain.LoginRequest) (*domain.LoginResponse, error)
}

// authService implementasi dari AuthService
type authService struct {
	repo repository.UserRepository
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

// Register menangani pendaftaran user baru
func (s *authService) Register(req domain.RegisterRequest) (*domain.User, error) {
	// Cek apakah email sudah digunakan
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email sudah digunakan")
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Buat user baru
	newUser := &domain.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	// Simpan ke database
	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Login menangani proses login user
func (s *authService) Login(req domain.LoginRequest) (*domain.LoginResponse, error) {
	// Cari user berdasarkan email
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email tidak ditemukan")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	// Buat klaim token JWT
	claims := domain.JWTClaims{
		UserID:   user.ID,
		Username: user.Email,
		Role:     "user", // bisa disesuaikan jika pakai role
	}

	// Generate token JWT
	token, err := auth.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	// Kembalikan response token
	return &domain.LoginResponse{Token: token}, nil
}
