package handlers

type Handler interface {
	Property() *propertyHandler
	Auth() *authHandler
	Request() *requestHandler
}
