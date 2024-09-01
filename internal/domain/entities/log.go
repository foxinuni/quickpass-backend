package entities

import "time"

type Log struct {
	LogID    int
	Occasion *Occasion
	IsInside bool
	Time     time.Time
}

func NewLog(logID int, occasion *Occasion, isInside bool, time time.Time) *Log {
	return &Log{
		LogID:    logID,
		Occasion: occasion,
		IsInside: isInside,
		Time:     time,
	}
}

func (l *Log) GetLogID() int {
	return l.LogID
}

func (l *Log) GetOccasion() *Occasion {
	return l.Occasion
}

func (l *Log) GetIsInside() bool {
	return l.IsInside
}

func (l *Log) GetTime() time.Time {
	return l.Time
}

func (l *Log) SetLogID(logID int) {
	l.LogID = logID
}

func (l *Log) SetOccasion(occasion *Occasion) {
	l.Occasion = occasion
}

func (l *Log) SetIsInside(isInside bool) {
	l.IsInside = isInside
}

func (l *Log) SetTime(time time.Time) {
	l.Time = time
}
