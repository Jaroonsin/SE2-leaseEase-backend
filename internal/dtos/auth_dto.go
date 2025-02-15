package dtos

// @Description RegisterDTO represents a request for user registration.
type RegisterDTO struct {
	Email    string `json:"email" example:"john.doe@example.com"`        // user's email
	Password string `json:"password" example:"strongPassword123"`        // user's password
	Name     string `json:"name" example:"John"`                         // user's first name
	Address  string `json:"address" example:"1234 Main St, Springfield"` // user's address
	Role     string `json:"role" example:"lessee"`                       // role: lessor, lessee
}

// @Description LoginDTO represents a request for user login.
type LoginDTO struct {
	Email    string `json:"email" example:"john.doe@example.com"` // user's email
	Password string `json:"password" example:"strongPassword123"` // user's password
}

type JWTDTO struct {
	UserID int    `json:"user_id" example:"1"`
	Role   string `json:"role" example:"lessee" `
}
