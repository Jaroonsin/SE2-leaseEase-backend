package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
	LesseeService   LesseeService
	ReviewService   ReviewService
	PaymentService  PaymentService
	LessorService   LessorService
	UserService     UserService
	ChatService     ChatService
}

func NewService(repo repositories.Repository, logger *zap.Logger) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property(), logger),
		AuthService:     NewAuthService(repo.Auth(), logger),
		LesseeService:   NewLesseeService(repo.Lessee(), logger),
		ReviewService:   NewReviewService(repo.Review(), logger),
		PaymentService:  NewPaymentService(repo.Payment(), logger),
		LessorService:   NewLessorService(repo.Lessor(), logger),
		UserService:     NewUserService(repo.User(), logger),
		ChatService:     NewChatService(repo.Chat(), logger),
	}
}

func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}

func (s *service) Lessee() LesseeService {
	return s.LesseeService
}

func (s *service) Review() ReviewService {
	return s.ReviewService
}

func (s *service) Payment() PaymentService {
	return s.PaymentService
}

func (s *service) Lessor() LessorService {
	return s.LessorService
}

func (s *service) User() UserService {
	return s.UserService
}

func (s *service) Chat() ChatService {
	return s.ChatService
}
