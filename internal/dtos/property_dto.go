package dtos

type CreateDTO struct {
	MarketSlotID       uint    `json:"id"`       // market slot id
	LessorID           uint    `json:"lessorid"` // lessor id
	Location           string  `json:"location"` // location of the property
	Size               string  `json:"size"`     // size of the property
	Price              float64 `json:"price"`    // price of the property
	AvailabilityStatus string  `json:"status"`   // availability status of the property
}

type UpdateDTO struct {
	MarketSlotID       uint    `json:"id"`     // market slot id
	Price              float64 `json:"price"`  // price of the property
	AvailabilityStatus string  `json:"status"` // availability status of the property
}

type DeleteDTO struct {
	PropertyID uint `json:"id"` // property id
}
