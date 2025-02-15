package services

type Service interface {
	Property() PropertyService
	Auth() AuthService
	Request() RequestService
}
