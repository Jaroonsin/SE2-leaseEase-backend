package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
)

type propertyService struct {
	propertyRepo repositories.PropertyRepository
}

func NewPropertyService(propertyRepo repositories.PropertyRepository) PropertyService {
	return &propertyService{
		propertyRepo: propertyRepo,
	}
}

func (s *propertyService) CreateProperty(propertyDTO *dtos.CreateDTO) error {

	property := &models.MarketSlot{
		MarketSlotID : propertyDTO.MarketSlotID,
		LessorID : propertyDTO.LessorID,
		Location : propertyDTO.Location,
		Size : propertyDTO.Size,
		Price : propertyDTO.Price,
		AvailabilityStatus : propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.CreateProperty(property)
}

func (s *propertyService) UpdateProperty(propertyDTO *dtos.UpdateDTO) error {
	property := &models.MarketSlot{
		MarketSlotID : propertyDTO.MarketSlotID,
		Price : propertyDTO.Price,
		AvailabilityStatus : propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.UpdateProperty(property)
}

func (s *propertyService) DeleteProperty(propertyDTO *dtos.DeleteDTO) error {
	return s.propertyRepo.DeleteProperty(propertyDTO.PropertyID)
}
