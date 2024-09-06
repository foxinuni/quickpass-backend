package dtos

type UserXEvent struct {
	OccasionsID []int `json:"user_x_event" validate:"required"`
}
