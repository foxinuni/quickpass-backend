package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type SessionRepository interface {
	GetById(sessionID int) (*entities.Session, error)
	GetByToken(token string) (*entities.Session, error)
	Create(session *entities.Session) error
	Update(session *entities.Session) error
	Delete(session *entities.Session) error
}

type StoreSessionRepository struct {
	userStore    stores.UserStore
	sessionStore stores.SessionStore
}

func NewStoreSessionRepository(sessionStore stores.SessionStore, userStore stores.UserStore) SessionRepository {
	return &StoreSessionRepository{
		sessionStore: sessionStore,
		userStore:    userStore,
	}
}

func (r *StoreSessionRepository) GetById(sessionID int) (*entities.Session, error) {
	// Get the session from the store
	s, err := r.sessionStore.GetById(context.Background(), sessionID)
	if err != nil {
		return nil, err
	}

	u, err := r.userStore.GetById(context.Background(), s.UserID)
	if err != nil {
		return nil, err
	}

	// Convert the result to a Session entity
	user := ModelToUser(u)
	session := ModelToSession(s, user)
	return session, nil
}

func (r *StoreSessionRepository) GetByToken(token string) (*entities.Session, error) {
	// Get the session from the store
	s, err := r.sessionStore.GetByToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	u, err := r.userStore.GetById(context.Background(), s.UserID)
	if err != nil {
		return nil, err
	}

	// Convert the result to a Session entity
	user := ModelToUser(u)
	session := ModelToSession(s, user)
	return session, nil
}

func (r *StoreSessionRepository) Create(session *entities.Session) error {
	model := SessionToModel(session)
	if err := r.sessionStore.Create(context.Background(), model); err != nil {
		return err
	}

	*session = *ModelToSession(model, session.GetUser())
	return nil
}

func (r *StoreSessionRepository) Update(session *entities.Session) error {
	model := SessionToModel(session)
	if err := r.sessionStore.Update(context.Background(), model); err != nil {
		return err
	}

	*session = *ModelToSession(model, session.GetUser())
	return nil
}

func (r *StoreSessionRepository) Delete(session *entities.Session) error {
	model := SessionToModel(session)
	if err := r.sessionStore.Delete(context.Background(), model.SessionID); err != nil {
		return err
	}

	return nil
}
