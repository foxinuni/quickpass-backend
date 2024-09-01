package models

type Session struct {
	SessionID  int
	UserID     int
	Enabled    bool
	Token      string
	PhoneModel string
	IMEI       string
}

func NewSession(sessionID int, userID int, enabled bool, token string, phoneModel string, imei string) *Session {
	return &Session{
		SessionID:  sessionID,
		UserID:     userID,
		Enabled:    enabled,
		Token:      token,
		PhoneModel: phoneModel,
		IMEI:       imei,
	}
}
