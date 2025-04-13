package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) GetAllUsers(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := r.db.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *adminRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *adminRepository) CountReviewsByPropertyForAdmin(queryString string, totalRecords *int64) error {
	return r.db.Model(&models.PropertyReview{}).
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Where("properties.name ILIKE ?", "%"+queryString+"%").
		Count(totalRecords).Error
}

func (r *adminRepository) GetAllReviewsSortedByPropertyName(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Preload("Review").
		Preload("Lessee").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("properties.name " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) GetAllReviewsSortedByReviewer(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN users ON users.id = property_reviews.lessee_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("users.name ILIKE ?", "%"+query+"%").
		Order("users.name " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) GetAllReviewsSortedByTime(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview

	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Joins("JOIN reviews ON reviews.id = property_reviews.review_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("reviews.time_stamp " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) GetPaginatedReviewsSortedByPropertyName(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Preload("Review").
		Preload("Lessee").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("properties.name " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) GetPaginatedReviewsSortedByReviewer(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN lessees ON users.id = property_reviews.lessee_id").
		Preload("Review").
		Preload("Lessee").
		Where("users.name ILIKE ?", "%"+query+"%").
		Order("users.name " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) GetPaginatedReviewsSortedByTime(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Joins("JOIN reviews ON reviews.id = property_reviews.review_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("reviews.time_stamp " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *adminRepository) UpdateUserStatus(id int, status string) error {
	err := r.db.Model(&models.User{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) DeleteReviewByID(id int) error {
	err := r.db.Model(&models.Review{}).Where("id = ?", id).Delete(&models.Review{}).Error
	if err != nil {
		return err
	}
	return nil
}
