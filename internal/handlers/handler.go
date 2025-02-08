package handlers

type Handler interface {
	Auth() *userHandler
	Property() *propertyHandler
}
