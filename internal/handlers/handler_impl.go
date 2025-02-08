package handlers

import "LeaseEase/internal/services"

type handler struct {
	UserHandler     *userHandler
	PropertyHandler *propertyHandler
}

func NewHandler(service services.Service) Handler {
	return &handler{
		UserHandler:     NewUserHandler(service.User()),
		PropertyHandler: NewPropertyHandler(service.Property()),
	}
}

func (h *handler) Auth() *userHandler {
	return h.UserHandler
}

func (h *handler) Property() *propertyHandler {
	return h.PropertyHandler
}
