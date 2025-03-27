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
