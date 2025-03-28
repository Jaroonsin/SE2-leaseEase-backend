package dtos

type PaymentDTO struct {
	ReservationID uint   `json:"reservation_id"`
	Currency      string `json:"currency"` // Example: "THB", "USD"
	Token         string `json:"token"`    // Example: "tokn_test_xxyy69btt9rnb5mir5b"
}
