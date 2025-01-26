package services

import (
	"LeaseEase/internal/dtos"
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

func (s *propertyService) ListAllProperties() ([]dtos.GetPropertyDTO, error) {
	properties, err := s.propertyRepo.GetAllProperty()
	if err != nil {
		return nil, err
	}

	var propertyDTOs []dtos.GetPropertyDTO
	for _, property := range properties {
		propertyDTO := dtos.GetPropertyDTO{
			ID:                 property.MarketSlotID,
			LessorID:           property.LessorID,
			Location:           property.Location,
			Size:               property.Size,
			Price:              property.Price,
			AvailabilityStatus: property.AvailabilityStatus,
		}
		propertyDTOs = append(propertyDTOs, propertyDTO)
	}

	return propertyDTOs, nil
}

func (s *propertyService) FindPropertyByID(propertyID uint) (*dtos.GetPropertyDTO, error) {
	property, err := s.propertyRepo.GetPropertyById(propertyID)
	if err != nil {
		return nil, err
	}

	propertyDTO := &dtos.GetPropertyDTO{
		ID:                 property.MarketSlotID,
		LessorID:           property.LessorID,
		Location:           property.Location,
		Size:               property.Size,
		Price:              property.Price,
		AvailabilityStatus: property.AvailabilityStatus,
	}

	return propertyDTO, nil
}
