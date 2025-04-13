package services

import "LeaseEase/internal/dtos"

type AdminService interface {
	GetAllReviewsForAdmin(page int, pageSize int, queryString string, sortParam string, direction string) (*dtos.GetReviewPaginatedDTO, error)
	GetAllUsers(pageInt, pageSizeInt int) ([]dtos.UserForAdminDTO, error)
	ManageUserStatus(id int, status string) error
	DeleteReview(id int) error
}
