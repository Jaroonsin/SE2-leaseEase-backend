package handlers

type Handler interface {
	User() *userHandler
	Property() *propertyHandler
	Auth() *authHandler
}
