package services

import "LeaseEase/internal/dtos"

type PropertyService interface {
	CreateProperty(propertyDTO *dtos.PropertyDTO, lessorID uint) error
	UpdateProperty(propertyDTO *dtos.PropertyDTO, propertyID uint) error
	DeleteProperty(propertyID uint) error
	GetAllProperty(lessorID uint, page, pageSize int) (*dtos.GetPropertyPaginatedDTO, error)
	GetPropertyByID(propertyID uint) (*dtos.GetPropertyDTO, error)
	SearchProperty(query map[string]string) ([]dtos.GetPropertyDTO, error)
	AutoComplete(query string) ([]string, error)
}
