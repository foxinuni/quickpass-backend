package auth

import (
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
)

type AuthServiceStrategy struct {
	authService services.AuthService
}

func NewAuthServiceStrategy(authService services.AuthService) *AuthServiceStrategy {
	return &AuthServiceStrategy{
		authService: authService,
	}
}

func (as *AuthServiceStrategy) Authenticate(token string) (*entities.Session, error) {
	return as.authService.ValidateSession(token)
}
