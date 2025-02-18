package dtos

type CreateReviewDTO struct {
	ReviewMessage string `json:"review_message" binding:"required"`
	Rating        int    `json:"rating" binding:"required,min=1,max=5"`
	PropertyID    uint   `json:"property_id" binding:"required"`
}

type ReviewResponseDTO struct {
	ReviewID      uint   `json:"review_id"`
	ReviewMessage string `json:"review_message"`
	Rating        int    `json:"rating"`
	LesseeID      uint   `json:"lessee_id"`
	PropertyID    uint   `json:"property_id"`
}
