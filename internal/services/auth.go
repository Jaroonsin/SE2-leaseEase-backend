package services

import "LeaseEase/internal/dtos"

type AuthService interface {
	Register(registerDTO *dtos.RegisterDTO) error
	Login(loginDTO *dtos.LoginDTO) (string, error)
}
