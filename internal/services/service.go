package services

type Service interface {
	Property() PropertyService
	Auth() AuthService
	Lessee() LesseeService
	Review() ReviewService
	Payment() PaymentService
	Lessor() LessorService
}
