package models

type Session struct {
	SessionID  int
	UserID     int
	Enabled    bool
	Token      string
	PhoneModel string
}

func NewSession(sessionID int, userID int, enabled bool, token string, phoneModel string) *Session {
	return &Session{
		SessionID:  sessionID,
		UserID:     userID,
		Enabled:    enabled,
		Token:      token,
		PhoneModel: phoneModel,
	}
}
