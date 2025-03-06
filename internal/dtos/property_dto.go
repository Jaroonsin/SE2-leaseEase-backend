package dtos

// @Description PropertyDTO represents a property.
type PropertyDTO struct {
	Name               string  `json:"name" example:"Sunset Villa"`                      // name of the property
	Location           string  `json:"location" example:"California"`                    // property location
	Size               float64 `json:"size" example:"3500.0"`                            // property size
	Price              float64 `json:"price" example:"1200000.50"`                       // property price
	AvailabilityStatus string  `json:"status" example:"available"`                       // availability status
	Details            string  `json:"details" example:"Spacious villa with a sea view"` // Property details
}

// @Description GetPropertyDTO represents the details of a property along with identifier information.
type GetPropertyDTO struct {
	PropertyID         uint    `json:"id" example:"1"`                          // Property ID
	LessorID           uint    `json:"lessor_id" example:"10"`                  // Lessor ID
	Name               string  `json:"name" example:"Sunset Villa"`             // name of the property
	Location           string  `json:"location" example:"California"`           // property's location
	Size               float64 `json:"size" example:"3500.0"`                   // property's size
	Price              float64 `json:"price" example:"1200000.50"`              // property's price
	AvailabilityStatus string  `json:"availability_status" example:"available"` // property's availability status
	Date               string  `json:"date" example:"2024-02-15T10:00:00Z"`     // Property creation date
	Rating             float64 `json:"rating" example:"4.5"`                    // Average rating
	ReviewCount        int     `json:"review_count" example:"12"`               // Number of reviews
	ReviewIDs          []uint  `json:"review_ids"`
	Details            string  `json:"details" example:"Spacious villa with a sea view"` // Property details
}

// @Description GetPropertyPaginatedDTO represents the response structure for retrieving a list of properties with pagination details.
type GetPropertyPaginatedDTO struct {
	Properties   []GetPropertyDTO `json:"properties"`    // List of properties retrieved from the database
	TotalRecords int              `json:"total_records"` // Total number of properties available
	TotalPages   int              `json:"total_pages"`   // Total number of pages based on total records and page size
	CurrentPage  int              `json:"current_page"`  // The current page number being retrieved
	PageSize     int              `json:"page_size"`     // Number of properties displayed per page
}

// @Description SearchPropertyDTO represents the details of a property along with identifier information.
type SearchPropertyDTO struct {
	PropertyID  uint    `json:"id" example:"1"`                // Property ID
	Name        string  `json:"name" example:"Sunset Villa"`   // name of the property
	Location    string  `json:"location" example:"California"` // property's location
	Size        float64 `json:"size" example:"3500.0"`         // property's size
	Price       float64 `json:"price" example:"1200000.50"`    // property's price
	Rating      float64 `json:"rating" example:"4.5"`          // Average rating
	ReviewCount int     `json:"review_count" example:"12"`     // Number of reviews
}
