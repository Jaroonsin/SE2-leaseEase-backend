package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type requestService struct {
	requestRepo repositories.RequestRepository
	logger      *zap.Logger
}

func NewRequestService(requestRepo repositories.RequestRepository, logger *zap.Logger) RequestService {
	return &requestService{
		requestRepo: requestRepo,
		logger:      logger,
	}
}

func (r *requestService) CreateRequest(requestDTO *dtos.CreateRequest, lesseeId uint) error {
	request := &models.Request{
		LesseeID:           lesseeId,
		Purpose:            requestDTO.Purpose,
		ProposedMessage:    requestDTO.ProposedMessage,
		Question:           requestDTO.Question,
		InterestedProperty: requestDTO.InterestedProperty,
	}

	return r.requestRepo.CreateRequest(request)
}

func (r *requestService) UpdateRequest(requestDTO *dtos.UpdateRequest, requestID uint) error {
	request := &models.Request{
		ID:              requestID,
		Purpose:         requestDTO.Purpose,
		ProposedMessage: requestDTO.ProposedMessage,
		Question:        requestDTO.Question,
	}
	return r.requestRepo.UpdateRequest(request)
}

func (r *requestService) DeleteRequest(requestID uint) error {
	return r.requestRepo.DeleteRequest(requestID)
}
