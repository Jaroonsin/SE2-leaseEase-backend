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
	UserID uint   `json:"user_id" example:"1"`			// user's ID
	Role   string `json:"role" example:"lessee" `		// user's role
}
type AuthCheckDTO struct {
	UserID uint   `json:"user_id" example:"1"`		// user's ID
	Role   string `json:"role" example:"lessee" `	// user's role
}

type RequestOTPDTO struct {
	Email string `json:"email" example:"john.doe@example.com" binding:"required"`	// user's email
}

type VerifyOTPDTO struct {
	Email string `json:"email" example:"john.doe@example.com" binding:"required"`	// user's email
	OTP   string `json:"otp" example:"123456" binding:"required"`					// user's OTP
}
