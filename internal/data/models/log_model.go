package models

import "time"

type Log struct {
	LogID      int
	OccasionID int
	IsInside   bool
	Time       time.Time
}

func NewLog(logID int, occasionID int, isInside bool, time time.Time) *Log {
	return &Log{
		LogID:      logID,
		OccasionID: occasionID,
		IsInside:   isInside,
		Time:       time,
	}
}
