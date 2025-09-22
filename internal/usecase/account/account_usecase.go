package account

import (
	"ave_project/internal/domain"
	"errors"
)

type AccountUsecase struct {
	Repo domain.PasswordRepository
}

func (u *AccountUsecase) CreateAccount(p domain.Account) error {
	if p.Name == "" || p.Login == "" || p.Password == "" {
		return errors.New("all fields must be filled")
	}
	return u.Repo.CreatePassword(p)
}

func (u *AccountUsecase) GetAccounts(userID int) ([]domain.Account, error) {
	return u.Repo.GetPassword(userID)
}

func (u *AccountUsecase) UpdateAccount(p domain.Account) error {
	return u.Repo.UpdatePassword(p)
}

func (u *AccountUsecase) DeleteAccount(userID, id int) error {
	return u.Repo.DeletePassword(userID, id)
}
