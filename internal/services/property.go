package services

import "LeaseEase/internal/dtos"

type PropertyService interface {
	CreateProperty(propertyDTO *dtos.CreateDTO) error
	UpdateProperty(propertyDTO *dtos.UpdateDTO) error
	DeleteProperty(propertyID *dtos.DeleteDTO) error
}
