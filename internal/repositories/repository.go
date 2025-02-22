package repositories

type Repository interface {
	User() UserRepository
	Property() PropertyRepository
	Auth() AuthRepository
	Reservation() ReservationRepository
	Review() ReviewRepository
	Payment() PaymentRepository
}
