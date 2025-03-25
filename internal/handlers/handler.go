package handlers

type Handler interface {
	Property() *propertyHandler
	Auth() *authHandler
	Lessee() *lesseeHandler
	Review() *reviewHandler
	Payment() *paymentHandler
	Lessor() *lessorHandler
	User() *userHandler
	Chat() *chatHandler
}
