package dtos

type UpdateUserDTO struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateImageDTO struct {
	ImageURL string `json:"image_url"`
}

type CheckUserDTO struct {
	UserID   uint   `json:"id" example:"1"`                                    // user's ID
	Role     string `json:"role" example:"lessee" `                            // user's role
	Email    string `json:"email" example:"john.doe@example.com"`              // user's email
	Name     string `json:"name" example:"John"`                               // user's first name
	Address  string `json:"address" example:"1234 Main St, Springfield"`       // user's address
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"` // user's image URL
}

type GetUserDTO struct {
	Name     string `json:"name" example:"John"`                               // user's first name
	Address  string `json:"address" example:"1234 Main St, Springfield"`       // user's address
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"` // user's image URL
}

type ChangePassDTO struct {
	OldPassword string `json:"old_password" example:"strongPassword123"`
	NewPassword string `json:"new_password" example:"superstrongPassword123"`
}

type UserForAdminDTO struct {
	ID       uint   `json:"id" example:"1"`                                    // user's ID
	Role     string `json:"role" example:"lessee"`                             // user's role
	Name     string `json:"name" example:"John"`                               // user's first name
	Address  string `json:"address" example:"1234 Main St, Springfield"`       // user's address
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"` // user's image URL
}

type UpdateUserStatusDTO struct {
	Status string `json:"status" example:"active"` // user's status (active, warned, banned)
}
