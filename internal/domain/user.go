package domain

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepository interface {
	CreateUser(username, passwordHash string) error
	GetPasswordHash(username string) (string, error)
}
