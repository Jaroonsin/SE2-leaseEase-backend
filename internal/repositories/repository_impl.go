package repositories

import (
	"LeaseEase/config"

	"gorm.io/gorm"
)

type repository struct {
	UserRepository     UserRepository
	PropertyRepository PropertyRepository
}

func NewRepository(cfg *config.DBConfig, db *gorm.DB) Repository {
	return &repository{
		UserRepository:     NewUserRepository(db),
		PropertyRepository: NewPropertyRepository(db),
	}
}

func (r *repository) User() UserRepository {
	return r.UserRepository
}

func (r *repository) Property() PropertyRepository {
	return r.PropertyRepository
}
