package repository

import (
	"database/sql"
	"errors"
	"lifekost/auth-services/pkg/domain"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, username, email, password, role, created_at FROM users WHERE email=$1"

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // email tidak ditemukan
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Create(user *domain.User) error {
	query := "INSERT INTO users (username, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Role, user.CreatedAt).
		Scan(&user.ID)

	return err
}
