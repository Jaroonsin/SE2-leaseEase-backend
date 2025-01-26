package dtos

type GetPropertyDTO struct {
	ID                 uint    `json:"id"`                  // Property/MarketSlot ID
	LessorID           uint    `json:"lessor_id"`           //
	Location           string  `json:"location"`            // Property's location
	Size               string  `json:"size"`                // Property's size
	Price              float64 `json:"price"`               // Property's price
	AvailabilityStatus string  `json:"availability_status"` // Property's status
}

type PaginatedPropertiesDTO struct {
	Properties []GetPropertyDTO `json:"properties"`  // List of properties
	Total      int64            `json:"total"`       // Total number of records
	Page       int              `json:"page"`        // Current page number
	Limit      int              `json:"limit"`       // Records per page
	TotalPages int64            `json:"total_pages"` // Total number of pages
}
