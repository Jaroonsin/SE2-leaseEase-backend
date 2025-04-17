package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"errors"

	"go.uber.org/zap"
)

type adminService struct {
	adminRepo repositories.AdminRepository
	logger    *zap.Logger
}

func NewAdminService(adminRepo repositories.AdminRepository, logger *zap.Logger) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		logger:    logger,
	}
}

func (s *adminService) GetAllReviewsForAdmin(page int, pageSize int, queryString string, sortParam string, direction string) (*dtos.GetReviewPaginatedDTO, error) {
	logger := s.logger.Named("GetAllReviewsForAdmin")
	var propertyReviews []models.PropertyReview
	var totalRecords int64
	var err error

	if direction == "" {
		direction = "asc"
	}

	if sortParam == "" {
		sortParam = "name"
	}

	if direction != "ASC" && direction != "DESC" {
		logger.Error("Invalid sort direction", zap.String("direction", direction))
		return nil, errors.New("invalid sort direction")
	}

	if page == 0 || pageSize == 0 {
		switch sortParam {
		case "name":
			propertyReviews, err = s.adminRepo.GetAllReviewsSortedByPropertyName(queryString, direction)
		case "reviewer":
			propertyReviews, err = s.adminRepo.GetAllReviewsSortedByReviewer(queryString, direction)
		case "time":
			propertyReviews, err = s.adminRepo.GetAllReviewsSortedByTime(queryString, direction)
		default:
			logger.Error("Invalid sort parameter", zap.String("sortParam", sortParam))
			return nil, errors.New("invalid sort parameter")
		}

		if err != nil {
			logger.Error("Failed to fetch all reviews", zap.Error(err))
			return nil, err
		}

		totalRecords = int64(len(propertyReviews))

	} else {
		// With pagination
		err = s.adminRepo.CountReviewsByPropertyForAdmin(queryString, &totalRecords)
		if err != nil {
			return nil, err
		}

		offset := (page - 1) * pageSize

		switch sortParam {
		case "name":
			propertyReviews, err = s.adminRepo.GetPaginatedReviewsSortedByPropertyName(pageSize, offset, queryString, direction)
		case "reviewer":
			propertyReviews, err = s.adminRepo.GetPaginatedReviewsSortedByReviewer(pageSize, offset, queryString, direction)
		case "time":
			propertyReviews, err = s.adminRepo.GetPaginatedReviewsSortedByTime(pageSize, offset, queryString, direction)
		default:
			logger.Error("Invalid sort parameter", zap.String("sortParam", sortParam))
			return nil, errors.New("invalid sort parameter")
		}

		if err != nil {
			logger.Error("Failed to fetch paginated reviews", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.Error(err))
			return nil, err
		}
	}

	// Calculate total pages
	totalPages := 1
	if pageSize > 0 {
		totalPages = int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
	}

	// Convert to DTO
	var reviewDTOs []dtos.GetReviewDTO
	for _, pr := range propertyReviews {
		reviewDTOs = append(reviewDTOs, dtos.GetReviewDTO{
			ReviewID:      pr.Review.ID,
			ReviewMessage: pr.Review.ReviewMessage,
			Rating:        pr.Review.Rating,
			TimeStamp:     pr.Review.TimeStamp,
			LesseeName:    pr.Lessee.Name,
		})
	}

	// Prepare final response
	responseForAdmin := dtos.GetReviewPaginatedDTO{
		Reviews:      reviewDTOs,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
	}

	logger.Info("Success fetching reviews for admin", zap.Int("count", len(reviewDTOs)))
	return &responseForAdmin, nil
}

func (s *adminService) GetAllUsers(pageInt, pageSizeInt int) ([]dtos.UserForAdminDTO, error) {
	offset := (pageInt - 1) * pageSizeInt
	limit := pageSizeInt
	if pageInt == 0 || pageSizeInt == 0 {
		offset = 0
		limit = 10 // Default limit if no pagination is provided
	}
	users, err := s.adminRepo.GetAllUsers(limit, offset)
	if err != nil {
		return nil, err
	}

	var userDTOs []dtos.UserForAdminDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dtos.UserForAdminDTO{
			ID:       user.ID,
			Name:     user.Name,
			Role:     user.UserType,
			Status:   user.Status,
			Address:  user.Address,
			ImageURL: user.ImageURL,
		})
	}

	return userDTOs, nil
}

func (s *adminService) ManageUserStatus(id int, status string) error {
	logger := s.logger.Named("ManageUserStatus")
	if status != "active" && status != "warned" && status != "banned" {
		logger.Error("Invalid status provided", zap.String("status", status))
		return errors.New("invalid status provided")
	}

	_, err := s.adminRepo.GetUserByID(id)
	if err != nil {
		logger.Error("Failed to fetch user by ID", zap.Int("id", id), zap.Error(err))
		return err
	}

	err = s.adminRepo.UpdateUserStatus(id, status)
	if err != nil {
		logger.Error("Failed to update user status", zap.Int("id", id), zap.Error(err))
		return err
	}

	logger.Info("User status updated successfully", zap.Int("id", id), zap.String("status", status))
	return nil
}

func (s *adminService) DeleteReview(id int) error {
	logger := s.logger.Named("DeleteReview")
	_, err := s.adminRepo.GetUserByID(id)
	if err != nil {
		logger.Error("Failed to fetch review by ID", zap.Int("id", id), zap.Error(err))
		return err
	}

	err = s.adminRepo.DeleteReviewByID(id)
	if err != nil {
		logger.Error("Failed to delete review", zap.Int("id", id), zap.Error(err))
		return err
	}

	logger.Info("Review deleted successfully", zap.Int("id", id))
	return nil
}
