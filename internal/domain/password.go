package domain

//Структура которая хранит информацию о пароле

type Account struct {
	ID       int
	UserID   int
	Name     string
	Url      string
	Login    string
	Password string
}

type PasswordRepository interface {
	CreatePassword(p Account) error
	GetPassword(userID int) ([]Account, error)
	GetPasswordByID(userID, id int) (Account, error)
	UpdatePassword(p Account) error
	DeletePassword(userID, id int) error
}
