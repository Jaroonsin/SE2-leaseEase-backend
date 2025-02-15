package services

import "LeaseEase/internal/dtos"

type RequestService interface {
	CreateRequest(requestDTO *dtos.CreateRequest , lesseeId uint) error
	UpdateRequest(requestDTO *dtos.UpdateRequest, requestID uint) error
	DeleteRequest(requestID uint) error
}