package services

type Service interface {
	Property() PropertyService
	Auth() AuthService
}
