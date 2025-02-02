package handlers

type Handler interface {
	User() *userHandler
}
