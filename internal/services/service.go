package services

type Service interface {
	User() UserService
	Property() PropertyService
	Auth() AuthService
}
