package services

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or number")
	ErrServerError		= errors.New("server error with verification")
	ErrInvalidToken       = errors.New("invalid token")
	ErrSessionDisabled    = errors.New("session is disabled")
	ErrExpiredCode 		= errors.New("code has expired")
	ErrIncorrectCode	= errors.New("code is incorrect")
)

type WaitSession struct{
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Model string `json:"model" validate:"required"`
	Code int `json:"code" validate:"required"`
}

type AuthService interface {
	Login(email, number, model string) (error)
	SubmitCode(phone string, code int) (*entities.Session, error)
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
	smsService 	SMSService
	redisClient *redis.Client
}

func NewJwtAuthService(options JwtAuthServiceOptions, userRepo repo.UserRepository, sessionRepo repo.SessionRepository, smsService SMSService) AuthService {
	redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	return &JwtAuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		options:     options,
		redisClient: redisClient,
		smsService: smsService,
	}
}

func (a *JwtAuthService) Login(email, number, model string) (error) {
	// Get the user by email
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		if err == stores.ErrUserNotFound {
			return ErrInvalidCredentials
		}

		return err
	}
	// Check if the number matches
	if user.Number != number {
		return ErrInvalidCredentials
	}
	// generate and send code
	code := rand.Intn(8999) + 1000
	err = a.smsService.SendVerificationSMS(number, strconv.Itoa(code))
	if err != nil {
		return ErrServerError
	}

	// create wait session and store it in redis
	session := WaitSession{
		Email: email,
		Phone: number,
		Model: model,
		Code : code,
	}
	err = a.redisClient.Set(context.Background(), number, session, 5*time.Minute).Err()
	if err != nil{
		return ErrServerError
	}
    return nil
}

func (a *JwtAuthService) SubmitCode(phone string, code int) (*entities.Session, error){
	var session WaitSession
	// get current session from redis
	value, err := a.redisClient.Get(context.Background(), phone).Result()
	if err != nil {
		return nil, ErrExpiredCode
	}

    err = json.Unmarshal([]byte(value), &session)
    if err != nil {
        return nil, ErrServerError
    }
	//check if code is correct
	if session.Code != code{
		return nil, ErrIncorrectCode
	}

	//delete session from redis
	a.redisClient.Del(context.Background(), phone)

	user, err := a.userRepo.GetByNumber(phone)
	if err != nil{
		return nil, ErrServerError
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
	entitySession := entities.NewSession(0, user, false, tokenString, session.Model)
	if err := a.sessionRepo.Create(entitySession); err != nil {
		return nil, err
	}
	return entitySession, nil
}

func (a *JwtAuthService) Logout(session *entities.Session) error {
	return a.sessionRepo.Delete(session)
}

func (a *JwtAuthService) ValidateSession(token string) (*entities.Session, error) {
	// Parse the token
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

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
		return nil, ErrInvalidToken
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
