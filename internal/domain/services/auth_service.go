package services

import "github.com/foxinuni/quickpass-backend/internal/domain/entities"

type AuthService interface {
	Login(email, number string) (*entities.Session, error)
	Logout(session *entities.Session) error
	ValidateSession(token string) (*entities.Session, error)
	EnableSession(session *entities.Session) error
}

type JwtAuthServiceOptions interface {
	GetJwtSecret() string
}

type JwtAuthService struct {
	options JwtAuthServiceOptions
}

func NewJwtAuthService(options JwtAuthServiceOptions) AuthService {
	return &JwtAuthService{
		options: options,
	}
}

func (a *JwtAuthService) Login(email, number string) (*entities.Session, error) {
	panic("not implemented") // Implement me
}

func (a *JwtAuthService) Logout(session *entities.Session) error {
	panic("not implemented") // Implement me
}

func (a *JwtAuthService) ValidateSession(token string) (*entities.Session, error) {
	panic("not implemented") // Implement me
}

func (a *JwtAuthService) EnableSession(session *entities.Session) error {
	panic("not implemented") // Implement me
}
