package repositories

import (
	"LeaseEase/internal/models"
	"database/sql"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &propertyRepository{
		db: db,
	}
}

func (r *propertyRepository) CreateProperty(property *models.Property) error {
	return r.db.Create(property).Error
}

func (r *propertyRepository) UpdateProperty(property *models.Property) error {
	result := r.db.Model(&property).Updates(*property)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *propertyRepository) DeleteProperty(id uint) error {
	result := r.db.Delete(&models.Property{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *propertyRepository) GetAllProperty(lessorID uint) ([]models.Property, error) {
	var properties []models.Property
	err := r.db.Where("lessor_id = ?", lessorID).Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *propertyRepository) GetPaginatedProperty(lessorID uint, limit, offset int) ([]models.Property, error) {
	var properties []models.Property
	err := r.db.Where("lessor_id = ?", lessorID).Limit(limit).Offset(offset).Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *propertyRepository) CountPropertiesByLessor(lessorID uint, totalRecords *int64) error {
	return r.db.Model(&models.Property{}).Where("lessor_id = ?", lessorID).Count(totalRecords).Error
}

func (r *propertyRepository) GetPropertyReviewsData(properties []models.Property) ([]float64, []int, [][]uint, error) {
	var ratings []float64
	var reviewCounts []int
	var reviewIDsList [][]uint

	for _, property := range properties {
		var rating float64
		var reviewCount int
		var reviewIDsArray string

		err := r.db.Raw(`
			SELECT 
				COALESCE(AVG(r.rating), 0) AS avg_rating, 
				COUNT(pr.review_id) AS review_count,
				COALESCE(array_agg(pr.review_id), '{}') AS review_ids
			FROM property_reviews pr
			LEFT JOIN reviews r ON pr.review_id = r.id
			WHERE pr.property_id = ?
			GROUP BY pr.property_id
		`, property.ID).Row().Scan(&rating, &reviewCount, &reviewIDsArray)

		if err == sql.ErrNoRows {
			rating = 0
			reviewCount = 0
			reviewIDsArray = "{}"
		} else if err != nil {
			return nil, nil, nil, err
		}

		reviewIDs := parseArrayStringToUintSlice(reviewIDsArray)

		ratings = append(ratings, rating)
		reviewCounts = append(reviewCounts, reviewCount)
		reviewIDsList = append(reviewIDsList, reviewIDs)
	}

	return ratings, reviewCounts, reviewIDsList, nil
}

func (r *propertyRepository) GetPropertyReviewDataByID(propertyID uint) (float64, int, []uint, error) {
	var rating float64
	var reviewCount int
	var reviewIDsArray string

	err := r.db.Raw(`
		SELECT 
			COALESCE(AVG(r.rating), 0) AS avg_rating, 
			COUNT(pr.review_id) AS review_count,
			COALESCE(array_agg(pr.review_id), '{}') AS review_ids
		FROM property_reviews pr
		LEFT JOIN reviews r ON pr.review_id = r.id
		WHERE pr.property_id = ?
		GROUP BY pr.property_id
	`, propertyID).Row().Scan(&rating, &reviewCount, &reviewIDsArray)

	// กรณีไม่มีรีวิว
	if err == sql.ErrNoRows {
		return 0, 0, []uint{}, nil
	} else if err != nil {
		return 0, 0, nil, err
	}

	reviewIDs := parseArrayStringToUintSlice(reviewIDsArray)

	return rating, reviewCount, reviewIDs, nil
}

func (r *propertyRepository) GetPropertyById(propertyID uint) (*models.Property, error) {
	var property models.Property

	err := r.db.First(&property, propertyID).Error
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (r *propertyRepository) SearchProperty(query map[string]string) ([]models.Property, error) {
	var properties []models.Property
	dbQuery := r.db.Model(&models.Property{})

	// Filter by availability status.
	dbQuery = dbQuery.Where("availability_status = ?", "available")

	// Filter by minimum price if provided.
	if minStr, ok := query["minprice"]; ok && minStr != "" {
		if minVal, err := strconv.ParseFloat(minStr, 64); err == nil {
			dbQuery = dbQuery.Where("price >= ?", minVal)
		}
	}

	// Filter by maximum price if provided.
	if maxStr, ok := query["maxprice"]; ok && maxStr != "" {
		if maxVal, err := strconv.ParseFloat(maxStr, 64); err == nil {
			dbQuery = dbQuery.Where("price <= ?", maxVal)
		}
	}

	// Filter by minimum size if provided.
	if minStr, ok := query["minsize"]; ok && minStr != "" {
		if minVal, err := strconv.ParseFloat(minStr, 64); err == nil {
			dbQuery = dbQuery.Where("size >= ?", minVal)
		}
	}

	// Filter by maximum size if provided.
	if maxStr, ok := query["maxsize"]; ok && maxStr != "" {
		if maxVal, err := strconv.ParseFloat(maxStr, 64); err == nil {
			dbQuery = dbQuery.Where("size <= ?", maxVal)
		}
	}

	if minStr, ok := query["rating"]; ok && minStr != "" {
		if minVal, err := strconv.ParseFloat(minStr, 64); err == nil {
			dbQuery = dbQuery.Where("avg_rating >= ?", minVal)
		}
	}

	// Filter by location if provided.
	if location, ok := query["location"]; ok && location != "" {
		keyword := "%" + location + "%"
		dbQuery = dbQuery.Where("location ILIKE ?", keyword)
	}

	// Filter by name if provided.
	if name, ok := query["name"]; ok && name != "" {
		keyword := "%" + name + "%"
		dbQuery = dbQuery.Where("name ILIKE ?", keyword)
	}

	// Sort by price or size if provided.
	if sortBy, ok := query["sortby"]; ok && sortBy != "" {
		order := "ASC"
		if sortOrder, ok := query["order"]; ok && (sortOrder == "DESC" || sortOrder == "desc") {
			order = "DESC"
		}

		if sortBy == "price" || sortBy == "size" {
			dbQuery = dbQuery.Order(sortBy + " " + order)
		}
	}

	// Apply pagination if provided.
	if pageStr, ok := query["page"]; ok && pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		// Determine page size, default to 10 if not provided.
		pageSize := 10
		if pageSizeStr, ok := query["pagesize"]; ok && pageSizeStr != "" {
			if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
				pageSize = ps
			}
		}

		offset := (page - 1) * pageSize
		dbQuery = dbQuery.Limit(pageSize).Offset(offset)
	}

	// Execute the query.
	err := dbQuery.Find(&properties).Error
	return properties, err
}

func (r *propertyRepository) AutoComplete(query string) ([]string, error) {
	var properties []models.Property
	err := r.db.Where("name ILIKE ?", query+"%").Limit(10).Find(&properties).Error
	if err != nil {
		return nil, err
	}

	var names []string
	for _, property := range properties {
		names = append(names, property.Name)
	}

	return names, nil
}

func parseArrayStringToUintSlice(arrayStr string) []uint {
	arrayStr = strings.Trim(arrayStr, "{}")
	if arrayStr == "" {
		return []uint{}
	}
	strIDs := strings.Split(arrayStr, ",")
	var ids []uint
	for _, strID := range strIDs {
		id, err := strconv.ParseUint(strID, 10, 32)
		if err == nil {
			ids = append(ids, uint(id))
		}
	}
	return ids
}
