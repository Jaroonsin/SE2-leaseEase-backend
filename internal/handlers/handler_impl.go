package handlers

import (
	"LeaseEase/internal/services"
)

type handler struct {
	PropertyHandler *propertyHandler
	AuthHandler     *authHandler
	LesseeHandler   *lesseeHandler
	ReviewHandler   *reviewHandler
	PaymentHandler  *paymentHandler
	lessorHandler   *lessorHandler
	userHandler     *userHandler
}

// LEssor implements Handler.

func NewHandler(service services.Service) Handler {
	return &handler{
		PropertyHandler: NewPropertyHandler(service.Property()),
		AuthHandler:     NewAuthHandler(service.Auth()),
		LesseeHandler:   NewLesseeHandler(service.Lessee()),
		ReviewHandler:   NewReviewHandler(service.Review()),
		PaymentHandler:  NewPaymentHandler(service.Payment()),
		lessorHandler:   NewLessorHandler(service.Lessor()),
		userHandler:     NewUserHandler(service.User()),
	}
}

func (h *handler) Auth() *authHandler {
	return h.AuthHandler
}

func (h *handler) Property() *propertyHandler {
	return h.PropertyHandler
}

func (h *handler) Lessee() *lesseeHandler {
	return h.LesseeHandler
}

func (h *handler) Review() *reviewHandler {
	return h.ReviewHandler
}

func (h *handler) Payment() *paymentHandler {
	return h.PaymentHandler
}
func (h *handler) Lessor() *lessorHandler {
	return h.lessorHandler
}

func (h *handler) User() *userHandler {
	return h.userHandler
}
