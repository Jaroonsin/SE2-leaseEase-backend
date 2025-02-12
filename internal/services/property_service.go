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


	property := &models.Property{
		Name:               propertyDTO.Name,
		LessorID:           propertyDTO.LessorID,
		Location:           propertyDTO.Location,
		Size:               propertyDTO.Size,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.CreateProperty(property)
}

func (s *propertyService) UpdateProperty(propertyDTO *dtos.UpdateDTO) error {
	property := &models.Property{
		ID: 				propertyDTO.PropertyID,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.UpdateProperty(property)
}

func (s *propertyService) DeleteProperty(propertyDTO *dtos.DeleteDTO) error {
	return s.propertyRepo.DeleteProperty(propertyDTO.PropertyID)
}
func (s *propertyService) GetAllProperty(page, pageSize int) ([]dtos.GetPropertyDTO, error) {
	var properties []models.Property
	var err error

	// Case 1: Fetch all properties (when page and pageSize are 0)
	if page == 0 || pageSize == 0 {
		properties, err = s.propertyRepo.GetAllProperty()
		if err != nil {
			return nil, err
		}
	} else {
		// Case 2: Apply pagination
		offset := (page - 1) * pageSize
		properties, err = s.propertyRepo.GetPaginatedProperty(pageSize, offset)
		if err != nil {
			return nil, err
		}
	}

	// Convert to DTO
	var propertyDTOs []dtos.GetPropertyDTO
	for _, property := range properties {
		propertyDTO := dtos.GetPropertyDTO{
			Name:               property.Name,
			PropertyID:         property.ID,
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

func (s *propertyService) GetPropertyByID(propertyID uint) (*dtos.GetPropertyDTO, error) {
	property, err := s.propertyRepo.GetPropertyById(propertyID)
	if err != nil {
		return nil, err
	}

	propertyDTO := &dtos.GetPropertyDTO{
		Name:               property.Name,
		PropertyID:         property.ID,
		LessorID:           property.LessorID,
		Location:           property.Location,
		Size:               property.Size,
		Price:              property.Price,
		AvailabilityStatus: property.AvailabilityStatus,
	}

	return propertyDTO, nil
}
