package repositories

type Repository interface {
	User() UserRepository
	Property() PropertyRepository
	Auth() AuthRepository
	Request() RequestRepository
}
