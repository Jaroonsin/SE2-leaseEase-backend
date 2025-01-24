package handlers

import "LeaseEase/internal/services"

type handler struct {
	UserHandler *userHandler
}

func NewHandler(service services.Service) *handler {
	return &handler{
		UserHandler: NewUserHandler(service.User()),
	}
}

func (h *handler) User() *userHandler {
	return h.UserHandler
}
