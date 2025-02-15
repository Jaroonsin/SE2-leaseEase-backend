package repositories

import "LeaseEase/internal/models"

type RequestRepository interface {
	CreateRequest(request *models.Request) error
	UpdateRequest(request *models.Request) error
	DeleteRequest(requestID uint) error
}