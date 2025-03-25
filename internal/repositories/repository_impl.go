package repositories

import (
	"LeaseEase/config"

	"gorm.io/gorm"
)

type repository struct {
	UserRepository     UserRepository
	PropertyRepository PropertyRepository
	AuthRepository     AuthRepository
	LesseeRepository   LesseeRepository
	ReviewRepository   ReviewRepository
	PaymentRepository  PaymentRepository
	LessorRepository   LessorRepository
	ChatRepository     ChatRepository
}

// Lessor implements Repository.

func NewRepository(cfg *config.Config, db *gorm.DB) Repository {
	return &repository{
		UserRepository:     NewUserRepository(db),
		PropertyRepository: NewPropertyRepository(db),
		AuthRepository:     NewAuthRepository(db),
		LesseeRepository:   NewLesseeRepository(db),
		ReviewRepository:   NewReviewRepository(db),
		PaymentRepository:  NewPaymentRepository(db),
		LessorRepository:   NewLessorRepository(db),
		ChatRepository:     NewChatRepository(db),
	}
}

func (r *repository) User() UserRepository {
	return r.UserRepository
}

func (r *repository) Property() PropertyRepository {
	return r.PropertyRepository
}

func (r *repository) Auth() AuthRepository {
	return r.AuthRepository
}

func (r *repository) Lessee() LesseeRepository {
	return r.LesseeRepository
}

func (r *repository) Lessor() LessorRepository {
	return r.LessorRepository
}

func (r *repository) Review() ReviewRepository {
	return r.ReviewRepository
}

func (r *repository) Payment() PaymentRepository {
	return r.PaymentRepository
}

func (r *repository) Chat() ChatRepository {
	return r.ChatRepository
}
