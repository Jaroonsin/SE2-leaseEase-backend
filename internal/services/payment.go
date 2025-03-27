package services

type PaymentService interface {
	ProcessPayment(userID uint, currency string, token string, reservationID uint) error
}
