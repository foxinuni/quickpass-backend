package models

import "time"

type Log struct {
	LogID      int
	OccasionID int
	IsInside   bool
	Time       time.Time
}
