package services

type PaymentService interface {
	ProcessPayment(userID uint, amount int64, currency, token string) error
}
