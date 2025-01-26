package services

import "LeaseEase/internal/dtos"

type PropertyService interface {
	ListAllProperties() ([]dtos.GetPropertyDTO, error)
	FindPropertyByID(propertyID uint) (*dtos.GetPropertyDTO, error)
}
