package services

import "LeaseEase/internal/dtos"

type UserService interface {
	Register(registerDTO *dtos.RegisterDTO) error
	Login(loginDTO *dtos.LoginDTO) (string, error)
}
