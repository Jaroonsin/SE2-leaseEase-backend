package repositories

type Repository interface {
	User() UserRepository
	Property() PropertyRepository
	Auth() AuthRepository
	Lessee() LesseeRepository
	Review() ReviewRepository
	Payment() PaymentRepository
	Lessor() LessorRepository
	Chat() ChatRepository
	Image() ImageRepository
}
