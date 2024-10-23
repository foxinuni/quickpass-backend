package dtos

type LoginDTO struct {
	Email      string `json:"email" validate:"required,email"`
	Number     string `json:"number" validate:"required,len=10"`
	PhoneModel string `json:"phone_model" validate:"required"`
}
