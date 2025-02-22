package handlers

import (
	"LeaseEase/internal/services"
)

type handler struct {
	PropertyHandler    *propertyHandler
	AuthHandler        *authHandler
	ReservationHandler *reservationHandler
	ReviewHandler      *reviewHandler
	PaymentHandler     *paymentHandler
}

func NewHandler(service services.Service) Handler {
	return &handler{
		PropertyHandler:    NewPropertyHandler(service.Property()),
		AuthHandler:        NewAuthHandler(service.Auth()),
		ReservationHandler: NewReservationHandler(service.Reservation()),
		ReviewHandler:      NewReviewHandler(service.Review()),
		PaymentHandler:     NewPaymentHandler(service.Payment()),
	}
}

func (h *handler) Auth() *authHandler {
	return h.AuthHandler
}

func (h *handler) Property() *propertyHandler {
	return h.PropertyHandler
}

func (h *handler) Reservation() *reservationHandler {
	return h.ReservationHandler
}

func (h *handler) Review() *reviewHandler {
	return h.ReviewHandler
}

func (h *handler) Payment() *paymentHandler {
	return h.PaymentHandler
}
