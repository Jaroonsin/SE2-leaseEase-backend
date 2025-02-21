package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils/constant"

	"time"

	"go.uber.org/zap"
)

type propertyService struct {
	propertyRepo repositories.PropertyRepository
	logger       *zap.Logger
}

func NewPropertyService(propertyRepo repositories.PropertyRepository, logger *zap.Logger) PropertyService {
	return &propertyService{
		propertyRepo: propertyRepo,
		logger:       logger,
	}
}

func (s *propertyService) CreateProperty(propertyDTO *dtos.PropertyDTO, lessorID uint) (uint, error) {
	logger := s.logger.Named("CreateProperty")

	property := &models.Property{
		LessorID:           lessorID,
		Name:               propertyDTO.Name,
		Location:           propertyDTO.Location,
		Size:               propertyDTO.Size,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
		Details:            propertyDTO.Details,
	}

	err := s.propertyRepo.CreateProperty(property)

	if err != nil {
		logger.Error("Error in create property", zap.Error(err))
		return 0, err
	}

	logger.Info(constant.SuccessCreateProp, zap.String("Name", property.Name))
	return property.ID, nil
}

func (s *propertyService) UpdateProperty(propertyDTO *dtos.PropertyDTO, propertyID uint) error {
	logger := s.logger.Named("UpdateProperty")

	property := &models.Property{
		ID:                 propertyID,
		Name:               propertyDTO.Name,
		Location:           propertyDTO.Location,
		Size:               propertyDTO.Size,
		Price:              propertyDTO.Price,
		AvailabilityStatus: propertyDTO.AvailabilityStatus,
		Details:            propertyDTO.Details,
	}

	err := s.propertyRepo.UpdateProperty(property)
	if err != nil {
		logger.Error("Error in update property", zap.Error(err))
		return err
	}

	logger.Info(constant.SuccessUpdateProp, zap.Uint("ID", propertyID), zap.String("Name", property.Name))
	return err
}

func (s *propertyService) DeleteProperty(propertyID uint) error {
	logger := s.logger.Named("DeleteProperty")

	err := s.propertyRepo.DeleteProperty(propertyID)
	if err != nil {
		logger.Error("Error in delete property", zap.Error(err))
		return err
	}

	logger.Info(constant.SuccessDeleteProp, zap.Uint("propertyID", propertyID))
	return err
}

func (s *propertyService) GetAllProperty(lessorID uint, page, pageSize int) (*dtos.GetPropertyPaginatedDTO, error) {
	logger := s.logger.Named("GetAllProperty")
	var properties []models.Property
	var totalRecords int64
	var err error

	// Case 1: Fetch all properties (when page and pageSize are 0)
	if page == 0 || pageSize == 0 {
		properties, err = s.propertyRepo.GetAllProperty(lessorID)
		if err != nil {
			logger.Error("Failed to fetch all properties", zap.Error(err))
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
			logger.Error("Failed to fetch paginated properties", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.Error(err))
			return nil, err
		}
	}

	ratings, reviewCounts, reviewIDsList, err := s.propertyRepo.GetPropertyReviewsData(properties)
	if err != nil {
		return nil, err
	}

	totalPages := 1
	if pageSize > 0 {
		totalPages = int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
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
			ReviewIDs:          reviewIDsList[i],
			Details:            property.Details,
		}
		propertyDTOs = append(propertyDTOs, propertyDTO)
	}

	logger.Info(constant.SuccessGetAllProp, zap.Int("count", len(propertyDTOs)))
	return &dtos.GetPropertyPaginatedDTO{
		Properties:   propertyDTOs,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
	}, nil
}

func (s *propertyService) GetPropertyByID(propertyID uint) (*dtos.GetPropertyDTO, error) {
	logger := s.logger.Named("GetPropertyByID")
	property, err := s.propertyRepo.GetPropertyById(propertyID)
	if err != nil {
		logger.Error("Failed to fetch property", zap.Uint("propertyID", propertyID), zap.Error(err))
		return nil, err
	}

	rating, reviewCount, reviewIDs, err := s.propertyRepo.GetPropertyReviewDataByID(propertyID)
	if err != nil {
		return nil, err
	}

	propertyDTO := &dtos.GetPropertyDTO{
		PropertyID:         property.ID,
		LessorID:           property.LessorID,
		Name:               property.Name,
		Location:           property.Location,
		Size:               property.Size,
		Price:              property.Price,
		AvailabilityStatus: property.AvailabilityStatus,
		Date:               property.CreatedAt.Format(time.RFC3339),
		Rating:             rating,
		ReviewCount:        reviewCount,
		ReviewIDs:          reviewIDs,
		Details:            property.Details,
	}

	logger.Info(constant.SuccessGetByIDProp, zap.Uint("propertyID", propertyID))
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
			Details:            property.Details,
		}
		propertyDTOs = append(propertyDTOs, propertyDTO)
	}

	return propertyDTOs, nil
}

func (s *propertyService) AutoComplete(query string) ([]string, error) {
	return s.propertyRepo.AutoComplete(query)
}
