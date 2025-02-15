package dtos

// @Description PropertyDTO represents a property.
type PropertyDTO struct {
	Name               string  `json:"name" example:"Sunset Villa"`    // name of the property
	Location           string  `json:"location" example:"California"`  // property location
	Size               string  `json:"size" example:"3500 sqft"`       // property size
	Price              float64 `json:"price" example:"1200000.50"`       // property price
	AvailabilityStatus string  `json:"status" example:"available"`     // availability status
}

// @Description GetPropertyDTO represents the details of a property along with identifier information.
type GetPropertyDTO struct {
	PropertyID         uint    `json:"id" example:"1"`                  // Property ID
	LessorID           uint    `json:"lessor_id" example:"10"`           // Lessor ID
	Name               string  `json:"name" example:"Sunset Villa"`      // name of the property
	Location           string  `json:"location" example:"California"`    // property's location
	Size               string  `json:"size" example:"3500 sqft"`         // property's size
	Price              float64 `json:"price" example:"1200000.50"`         // property's price
	AvailabilityStatus string  `json:"availability_status" example:"available"` // property's availability status
}
