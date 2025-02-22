package services

type Service interface {
	Property() PropertyService
	Auth() AuthService
	Reservation() ReservationService
	Review() ReviewService
	Payment() PaymentService
}
