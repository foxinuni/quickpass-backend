package repo

import (
	"context"

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
	// Convert the User entity to a User model
	model := UserToModel(user)

	// Create the user in the store
	if err := r.userStore.Create(context.Background(), model); err != nil {
		return err
	}

	// Update the User entity with the new ID
	*user = *ModelToUser(model)
	return nil
}

func (r *StoreUserRepository) Update(user *entities.User) error {
	model := UserToModel(user)
	if err := r.userStore.Update(context.Background(), model); err != nil {
		return err
	}

	*user = *ModelToUser(model)
	return nil
}

func (r *StoreUserRepository) Delete(userID int) error {
	return r.userStore.Delete(context.Background(), userID)
}
