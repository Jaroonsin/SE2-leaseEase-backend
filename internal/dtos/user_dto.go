package dtos

type RegisterDTO struct {
	ID       uint   `json:"id"`       // student id
	Name     string `json:"name"`     // user's first name
	Address  string `json:"address"`  // user's address
	Email    string `json:"email"`    // user's email
	Password string `json:"password"` // user's password
	Role     string `json:"role"`     // role: lessor, lessee
}
