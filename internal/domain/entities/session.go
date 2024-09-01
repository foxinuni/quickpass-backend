package entities

type Session struct {
	SessionID  int
	User       *User
	Enabled    bool
	Token      string
	PhoneModel string
	IMEI       string
}

func NewSession(sessionID int, user *User, enabled bool, token string, phoneModel string, imei string) *Session {
	return &Session{
		SessionID:  sessionID,
		User:       user,
		Enabled:    enabled,
		Token:      token,
		PhoneModel: phoneModel,
		IMEI:       imei,
	}
}

func (s *Session) GetSessionID() int {
	return s.SessionID
}

func (s *Session) GetUser() *User {
	return s.User
}

func (s *Session) GetEnabled() bool {
	return s.Enabled
}

func (s *Session) GetToken() string {
	return s.Token
}

func (s *Session) GetPhoneModel() string {
	return s.PhoneModel
}

func (s *Session) GetIMEI() string {
	return s.IMEI
}

func (s *Session) SetSessionID(sessionID int) {
	s.SessionID = sessionID
}

func (s *Session) SetUser(user *User) {
	s.User = user
}

func (s *Session) SetToken(token string) {
	s.Token = token
}

func (s *Session) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

func (s *Session) SetPhoneModel(phoneModel string) {
	s.PhoneModel = phoneModel
}

func (s *Session) SetIMEI(imei string) {
	s.IMEI = imei
}
