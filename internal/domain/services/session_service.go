package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type SessionService interface {
	GetAllSessions() ([]*entities.Session, error)
	EnableSession(sessionID int, state bool) error
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

func (s *RepoSessionService) EnableSession(sessionID int, state bool) error {
	session, err := s.sessionRepo.GetById(sessionID)
	if err != nil {
		return err
	}

	session.Enabled = state
	return s.sessionRepo.Update(session)
}
