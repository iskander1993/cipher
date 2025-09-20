package repositories

import (
	"ave_project/internal/domain"
	"database/sql"
)

// PasswordRepository — работа с паролями (Account)
type PasswordRepository struct {
	DB *sql.DB
}

// NewPasswordRepository создаёт репозиторий для паролей
func NewPasswordRepository(db *sql.DB) *PasswordRepository {
	return &PasswordRepository{DB: db}
}

// CreatePassword сохраняет новый пароль
func (r *PasswordRepository) CreatePassword(p domain.Account) error {
	query := "INSERT INTO passwords (user_id, name, url, login, password) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.DB.Exec(query, p.UserID, p.Name, p.Url, p.Login, p.Password)
	return err
}

// GetPasswords возвращает все пароли пользователя
func (r *PasswordRepository) GetPasswords(userID int) ([]domain.Account, error) {
	query := "SELECT id, name, url, login, password FROM passwords WHERE user_id=$1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []domain.Account
	for rows.Next() {
		var p domain.Account
		p.UserID = userID
		err := rows.Scan(&p.ID, &p.Name, &p.Url, &p.Login, &p.Password)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, p)
	}
	return accounts, nil
}

// GetPasswordByID возвращает один пароль
func (r *PasswordRepository) GetPasswordByID(userID, id int) (domain.Account, error) {
	var p domain.Account
	query := "SELECT name, url, login, password FROM passwords WHERE user_id=$1 AND id=$2"
	err := r.DB.QueryRow(query, userID, id).Scan(&p.Name, &p.Url, &p.Login, &p.Password)
	if err != nil {
		return domain.Account{}, err
	}
	p.ID = id
	p.UserID = userID
	return p, nil
}

// UpdatePassword обновляет пароль
func (r *PasswordRepository) UpdatePassword(p domain.Account) error {
	query := "UPDATE passwords SET name=$1, url=$2, login=$3, password=$4 WHERE user_id=$5 AND id=$6"
	_, err := r.DB.Exec(query, p.Name, p.Url, p.Login, p.Password, p.UserID, p.ID)
	return err
}

// DeletePassword удаляет пароль
func (r *PasswordRepository) DeletePassword(userID, id int) error {
	query := "DELETE FROM passwords WHERE user_id=$1 AND id=$2"
	_, err := r.DB.Exec(query, userID, id)
	return err
}
