package models

type Session struct {
	SessionID  int
	UserID     int
	Enabled    bool
	Token      string
	PhoneModel string
	IMEI       string
}
