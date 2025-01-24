package services

type UserService interface {
	Register(email, password, role string) error
	Login(email, password string) (string, error)
}
