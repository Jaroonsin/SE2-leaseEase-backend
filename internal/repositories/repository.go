package repositories

type Repository interface {
	User() UserRepository
	Property() PropertyRepository
}
