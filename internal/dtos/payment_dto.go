package dtos

type PaymentDTO struct {
	UserID   uint   `json:"user_id"`
	Amount   int64  `json:"amount"`   // Stored in satangs na ja
	Currency string `json:"currency"` // Example: "THB", "USD"
	Token    string `json:"token"`    // Example: "tokn_test_xxyy69btt9rnb5mir5b"
}
