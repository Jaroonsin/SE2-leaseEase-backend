package services

import (
	"LeaseEase/internal/dtos"
)

type AuthService interface {
	Register(registerDTO *dtos.RegisterDTO) error
	Login(loginDTO *dtos.LoginDTO) (string, error)
	AuthCheck(token string) (*dtos.AuthCheckDTO, error)
	RequestOTP(requestOTPDTO *dtos.RequestOTPDTO) error
	VerifyOTP(verifyOTPDTO *dtos.VerifyOTPDTO) error
	RequestPasswordReset(resetPassRequest *dtos.ResetPassRequestDTO) (string, error)
	ResetPassword(resetPassDTO *dtos.ResetPassDTO) error
}
