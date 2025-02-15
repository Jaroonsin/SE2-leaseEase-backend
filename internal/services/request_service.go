package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
)

type requestService struct {
	requestRepo repositories.RequestRepository
}

func NewRequestService(requestRepo repositories.RequestRepository) RequestService {
	return &requestService{
		requestRepo: requestRepo,
	}
}

func (r *requestService) CreateRequest(requestDTO *dtos.CreateRequest , lesseeId uint) error {
	request := &models.Request{
		LesseeID: lesseeId,
		Purpose: requestDTO.Purpose,
		ProposedMessage: requestDTO.ProposedMessage,
		Question: requestDTO.Question,
		InterestedProperty: requestDTO.InterestedProperty,
	}
	
	return r.requestRepo.CreateRequest(request)
}

func (r *requestService) UpdateRequest(requestDTO *dtos.UpdateRequest, requestID uint) error {
	request := &models.Request{
		ID: requestID,
		Purpose: requestDTO.Purpose,
		ProposedMessage: requestDTO.ProposedMessage,
		Question: requestDTO.Question,
	}
	return r.requestRepo.UpdateRequest(request)
}

func (r *requestService) DeleteRequest(requestID uint) error {
	return r.requestRepo.DeleteRequest(requestID)
}