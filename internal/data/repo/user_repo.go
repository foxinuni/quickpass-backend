package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type UserRepository interface {
	GetById(userID int) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
}

type StoreUserRepository struct {
	userStore stores.UserStore
}

func NewStoreUserRepository(userStore stores.UserStore) UserRepository {
	return &StoreUserRepository{
		userStore: userStore,
	}
}

func (r *StoreUserRepository) GetById(userID int) (*entities.User, error) {
	// Get the user from the store
	result, err := r.userStore.GetById(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	// Convert the result to a User entity
	user := entities.NewUser(result.UserID, result.Email, result.Number)
	return user, nil
}

func (r *StoreUserRepository) GetByEmail(email string) (*entities.User, error) {
	// Get the user from the store
	result, err := r.userStore.GetByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	// Convert the result to a User entity
	user := entities.NewUser(result.UserID, result.Email, result.Number)
	return user, nil
}
