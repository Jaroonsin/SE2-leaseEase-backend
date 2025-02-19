package handlers

type Handler interface {
	Property() *propertyHandler
	Auth() *authHandler
	Reservation() *reservationHandler
	Review() *reviewHandler
}
