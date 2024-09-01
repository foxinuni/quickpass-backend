package services

import (
	"errors"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or number")
	ErrInvalidToken       = errors.New("invalid token")
	ErrSessionDisabled    = errors.New("session is disabled")
)

type AuthService interface {
	Login(email, number, model, emei string) (*entities.Session, error)
	Logout(session *entities.Session) error
	ValidateSession(token string) (*entities.Session, error)
	EnableSession(session *entities.Session) error
}

type JwtAuthServiceOptions interface {
	GetJwtSecret() string
}

type JwtAuthService struct {
	userRepo    repo.UserRepository
	sessionRepo repo.SessionRepository
	options     JwtAuthServiceOptions
}

func NewJwtAuthService(options JwtAuthServiceOptions, userRepo repo.UserRepository, sessionRepo repo.SessionRepository) AuthService {
	return &JwtAuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		options:     options,
	}
}

func (a *JwtAuthService) Login(email, number, model, emei string) (*entities.Session, error) {
	// Get the user by email
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		if err == stores.ErrUserNotFound {
			return nil, ErrInvalidCredentials
		}

		return nil, err
	}

	// Check if the number matches
	if user.Number != number {
		return nil, ErrInvalidCredentials
	}

	// Create a new JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(a.options.GetJwtSecret()))
	if err != nil {
		return nil, err
	}

	// Register the session
	session := entities.NewSession(0, user, false, tokenString, model, emei)
	if err := a.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	// Return the session
	return session, nil
}

func (a *JwtAuthService) Logout(session *entities.Session) error {
	return a.sessionRepo.Delete(session)
}

func (a *JwtAuthService) ValidateSession(token string) (*entities.Session, error) {
	// Parse the token
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.options.GetJwtSecret()), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !t.Valid {
		return nil, ErrInvalidToken
	}

	// Get the session by token
	session, err := a.sessionRepo.GetByToken(token)
	if err != nil {
		return nil, err
	}

	// Check if the session is valid
	if !session.Enabled {
		return nil, ErrSessionDisabled
	}

	// Return the session
	return session, nil
}

func (a *JwtAuthService) EnableSession(session *entities.Session) error {
	session.SetEnabled(true)
	return a.sessionRepo.Update(session)
}
