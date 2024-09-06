package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type SessionService interface {
	GetAllSessions() ([]*entities.Session, error)
}

type RepoSessionService struct {
	sessionRepo repo.SessionRepository
}

func NewRepoSessionService(sessionRepo repo.SessionRepository) SessionService {
	return &RepoSessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *RepoSessionService) GetAllSessions() ([]*entities.Session, error) {
	return s.sessionRepo.GetAll()
}
