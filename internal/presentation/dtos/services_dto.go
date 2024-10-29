package dtos

type SessionPatchDTO struct {
	Enabled *bool `json:"enabled" validate:"required"`
}
