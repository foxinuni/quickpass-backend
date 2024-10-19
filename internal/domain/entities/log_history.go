package entities

import "time"

type LogHistory struct {
	LogID    int       `json:"log_id"`
	Email 	 string 	`json: "email"`
	IsInside bool      `json:"is_inside"`
	Time     time.Time `json:"time"`
}

func NewLogHistory(logID int, email string, isInside bool, time time.Time) *LogHistory {
	return &LogHistory{
		LogID:    logID,
		Email:  email,
		IsInside: isInside,
		Time:     time,
	}
}
