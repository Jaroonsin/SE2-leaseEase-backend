package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"time"
)

type propertyService struct {
	propertyRepo repositories.PropertyRepository
}

func NewPropertyService(propertyRepo repositories.PropertyRepository) PropertyService {
	return &propertyService{
		propertyRepo: propertyRepo,
	}
}

func (s *propertyService) CreateProperty(propertyDTO *dtos.PropertyDTO, lessorID uint) error {

	property := &models.Property{
		LessorID:           lessorID,
		Name:               propertyDTO.Name,
		Location:           propertyDTO.Location,
		Size:               propertyDTO.Size,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.CreateProperty(property)
}

func (s *propertyService) UpdateProperty(propertyDTO *dtos.PropertyDTO, propertyID uint) error {
	property := &models.Property{
		ID:                 propertyID,
		Name:               propertyDTO.Name,
		Location:           propertyDTO.Location,
		Size:               propertyDTO.Size,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
	}

	return s.propertyRepo.UpdateProperty(property)
}

func (s *propertyService) DeleteProperty(propertyID uint) error {
	return s.propertyRepo.DeleteProperty(propertyID)
}
func (s *propertyService) GetAllProperty(lessorID uint, page, pageSize int) (*dtos.GetPropertyPaginatedDTO, error) {
	var properties []models.Property
	var totalRecords int64
	var err error

	// Case 1: Fetch all properties (when page and pageSize are 0)
	if page == 0 || pageSize == 0 {
		properties, err = s.propertyRepo.GetAllProperty(lessorID)
		if err != nil {
			return nil, err
		}
		totalRecords = int64(len(properties))
	} else {
		// Case 2: Apply pagination
		err = s.propertyRepo.CountPropertiesByLessor(lessorID, &totalRecords)
		if err != nil {
			return nil, err
		}

		offset := (page - 1) * pageSize
		properties, err = s.propertyRepo.GetPaginatedProperty(lessorID, pageSize, offset)
		if err != nil {
			return nil, err
		}
	}

	ratings, reviewCounts, err := s.propertyRepo.GetPropertyReviewsData(properties)
	if err != nil {
		return nil, err
	}

	totalPages := 1
	if pageSize > 0 {
		totalPages = int((totalRecords + int64(pageSize) - 1) / int64(pageSize)) // ปัดขึ้นเสมอ
	}

	// Convert to DTO
	var propertyDTOs []dtos.GetPropertyDTO
	for i, property := range properties {
		propertyDTO := dtos.GetPropertyDTO{
			PropertyID:         property.ID,
			LessorID:           property.LessorID,
			Name:               property.Name,
			Location:           property.Location,
			Size:               property.Size,
			Price:              property.Price,
			AvailabilityStatus: property.AvailabilityStatus,
			Date:               property.CreatedAt.Format(time.RFC3339),
			Rating:             ratings[i],
			ReviewCount:        reviewCounts[i],
		}
		propertyDTOs = append(propertyDTOs, propertyDTO)
	}

	return &dtos.GetPropertyPaginatedDTO{
		Properties:   propertyDTOs,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
	}, nil
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

func (s *propertyService) SearchProperty(query map[string]string) ([]dtos.GetPropertyDTO, error) {
	properties, err := s.propertyRepo.SearchProperty(query)
	if err != nil {
		return nil, err
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

func (s *propertyService) AutoComplete(query string) ([]string, error) {
	return s.propertyRepo.AutoComplete(query)
}
