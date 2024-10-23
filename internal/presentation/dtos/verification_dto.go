package dtos

type VerificationDTO struct {
	Number string `json:"number" validate:"required"`
	Code   string `json:"code" validate:"required"`
}