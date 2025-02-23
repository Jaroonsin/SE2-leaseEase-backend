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
}

func NewHandler(service services.Service) Handler {
	return &handler{
		PropertyHandler: NewPropertyHandler(service.Property()),
		AuthHandler:     NewAuthHandler(service.Auth()),
		LesseeHandler:   NewLesseeHandler(service.Lessee()),
		ReviewHandler:   NewReviewHandler(service.Review()),
		PaymentHandler:  NewPaymentHandler(service.Payment()),
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
