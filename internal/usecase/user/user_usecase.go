package user

import (
	"errors"

	"ave_project/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo domain.UserRepository
}

func (u *UserUsecase) Register(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = u.Repo.CreateUser(username, string(hash))
	if err != nil {
		return errors.New("username already exists")
	}
	return nil
}

func (u *UserUsecase) Login(username, password string) error {
	hash, err := u.Repo.GetPasswordHash(username)
	if err != nil {
		return errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("invalid username or password")
	}
	return nil
}
