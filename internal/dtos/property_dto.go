package dtos

type CreateDTO struct {
	Name               string  `json:"name"`     // name of the property
	LessorID           uint    `json:"lessorid"` // lessor id
	Location           string  `json:"location"` // location of the property
	Size               string  `json:"size"`     // size of the property
	Price              float64 `json:"price"`    // price of the property
	AvailabilityStatus string  `json:"status"`   // availability status of the property
}

type UpdateDTO struct {
	PropertyID         uint    `json:"id"`     // market slot id
	Price              float64 `json:"price"`  // price of the property
	AvailabilityStatus string  `json:"status"` // availability status of the property
}

type DeleteDTO struct {
	PropertyID uint `json:"id"` // property id
}

type GetPropertyDTO struct {
	PropertyID         uint    `json:"id"`                  // Property/MarketSlot ID
	LessorID           uint    `json:"lessor_id"`           //
	Location           string  `json:"location"`            // Property's location
	Size               string  `json:"size"`                // Property's size
	Price              float64 `json:"price"`               // Property's price
	AvailabilityStatus string  `json:"availability_status"` // Property's status
}
