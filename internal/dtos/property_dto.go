package dtos

type GetDTO struct {
	ID                 uint    `json:"id"`                  // Property/MarketSlot ID
	LessorID           uint    `json:"lessor_id"`           //
	Location           string  `json:"location"`            // Property's location
	Size               string  `json:"size"`                // Property's size
	Price              float64 `json:"price"`               // Property's price
	AvailabilityStatus string  `json:"availability_status"` // Property's status
}
