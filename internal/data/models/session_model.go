package models

type Session struct {
	SessionID  int
	UserID     int
	Token      string
	PhoneModel string
	IMEI       string
}
