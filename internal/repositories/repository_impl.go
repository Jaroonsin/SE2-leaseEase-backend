package repositories

import (
	"LeaseEase/config"

	"gorm.io/gorm"
)

type repository struct {
	UserRepository     UserRepository
	PropertyRepository PropertyRepository
	AuthRepository     AuthRepository
	ReservationRepository  ReservationRepository
	ReviewRepository   ReviewRepository
}

func NewRepository(cfg *config.Config, db *gorm.DB) Repository {
	return &repository{
		UserRepository:     NewUserRepository(db),
		PropertyRepository: NewPropertyRepository(db),
		AuthRepository:     NewAuthRepository(db),
		ReservationRepository:  NewReservationRepository(db),
		ReviewRepository:   NewReviewRepository(db),
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

func (r *repository) Reservation() ReservationRepository {
	return r.ReservationRepository
}

func (r *repository) Review() ReviewRepository {
	return r.ReviewRepository
}
