package dtos

import "time"

type CreateReviewDTO struct {
	ReviewMessage string `json:"review_message" binding:"required"`
	Rating        int    `json:"rating" binding:"required,min=0,max=5"`
	PropertyID    uint   `json:"property_id" binding:"required"`
}

type UpdateReviewDTO struct {
	ReviewMessage string `json:"review_message" binding:"required"`
	Rating        int    `json:"rating" binding:"required,min=0,max=5"`
}

type GetReviewDTO struct {
	ReviewID      uint      `json:"review_id"`
	ReviewMessage string    `json:"review_message"`
	Rating        int       `json:"rating"`
	TimeStamp     time.Time `json:"time_stamp"`
	LesseeName    string    `json:"lessee_name"`
}

type GetReviewPaginatedDTO struct {
	Reviews      []GetReviewDTO `json:"reviews"`
	TotalRecords int            `json:"total_records"`
	TotalPages   int            `json:"total_pages"`
	CurrentPage  int            `json:"current_page"`
	PageSize     int            `json:"page_size"`
}
