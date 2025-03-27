package services

import "LeaseEase/internal/dtos"

type UserService interface {
	UpdateUser(userID uint, User dtos.UpdateUserDTO) error
	UpdateImage(userID uint, Image dtos.UpdateImageDTO) error
	CheckUser(token string) (*dtos.CheckUserDTO, error)
	GetUser(userID uint) (*dtos.GetUserDTO, error)
}
