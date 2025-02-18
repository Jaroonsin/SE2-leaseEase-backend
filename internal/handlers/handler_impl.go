package handlers

import (
	"LeaseEase/internal/services"
)

type handler struct {
	PropertyHandler *propertyHandler
	AuthHandler     *authHandler
	RequestHandler  *requestHandler
	ReviewHandler   *reviewHandler
}

func NewHandler(service services.Service) Handler {
	return &handler{
		PropertyHandler: NewPropertyHandler(service.Property()),
		AuthHandler:     NewAuthHandler(service.Auth()),
		RequestHandler:  NewRequestHandler(service.Request()),
		ReviewHandler:   NewReviewHandler(service.Review()),
	}
}

func (h *handler) Auth() *authHandler {
	return h.AuthHandler
}

func (h *handler) Property() *propertyHandler {
	return h.PropertyHandler
}

func (h *handler) Request() *requestHandler {
	return h.RequestHandler
}

func (h *handler) Review() *reviewHandler {
	return h.ReviewHandler
}
