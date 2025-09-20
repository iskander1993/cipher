package repositories

import (
	"database/sql"
)

// UserRepository — работа с пользователями
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository создаёт новый репозиторий для пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser создаёт нового пользователя
func (r *UserRepository) CreateUser(username, passwordHash string) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, username, passwordHash)
	return err
}

// GetPasswordHash возвращает хеш пароля пользователя
func (r *UserRepository) GetPasswordHash(username string) (string, error) {
	var password string
	query := "SELECT password FROM users WHERE username=$1"
	err := r.DB.QueryRow(query, username).Scan(&password)
	return password, err
}
