package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type UserRepository interface {
	GetById(userID int) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(userID int) error
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
	user, err := r.userStore.GetById(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	// Convert the result to a User entity
	return ModelToUser(user), nil
}

func (r *StoreUserRepository) GetByEmail(email string) (*entities.User, error) {
	// Get the user from the store
	user, err := r.userStore.GetByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	// Convert the result to a User entity
	return ModelToUser(user), nil
}

func (r *StoreUserRepository) Create(user *entities.User) error {
	return r.userStore.Create(context.Background(), UserToModel(user))
}

func (r *StoreUserRepository) Update(user *entities.User) error {
	panic("not implemented")
}

func (r *StoreUserRepository) Delete(userID int) error {
	panic("not implemented")
}

func UserToModel(user *entities.User) *models.User {
	return models.NewUser(user.GetUserID(), user.GetEmail(), user.GetNumber())
}

func ModelToUser(model *models.User) *entities.User {
	return entities.NewUser(model.UserID, model.Email, model.Number)
}
