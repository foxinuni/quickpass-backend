package dtos

type PatchMyOccasionDTO struct {
	Confirming bool `json:"confirming" validate:"required"`
}
