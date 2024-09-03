package dtos

type ActionDTO struct {
	OccasionID int `json:"occasion_id" validate:"required"`
}
