package mock_storage

import (
	"context"

	"github.com/henrieto/account/models"
)

type UserStorage struct {
	storage map[string]*models.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[string]*models.User)}
}

func (strg *UserStorage) Create(_ context.Context, user *models.User) (*models.User, error) {
	strg.storage[user.Email] = user
	return user, nil
}
func (strg *UserStorage) CreateStaff(_ context.Context, user *models.User) (*models.User, error) {
	strg.storage[user.Email] = user
	return user, nil
}
func (strg *UserStorage) CreateSuperUser(_ context.Context, user *models.User) (*models.User, error) {
	strg.storage[user.Email] = user
	return user, nil
}
func (strg *UserStorage) Update(_ context.Context, user *models.User) (*models.User, error) {
	strg.storage[user.Email] = user
	return user, nil
}
func (strg *UserStorage) List(context.Context, string) ([]*models.User, error) {
	users := []*models.User{}
	for _, user := range strg.storage {
		users = append(users, user)
	}
	return users, nil
}
func (strg *UserStorage) Get(_ context.Context, id uint8) (*models.User, error) {
	return nil, nil
}
func (strg *UserStorage) GetByEmail(_ context.Context, email string) (*models.User, error) {
	return strg.storage[email], nil
}
func (strg *UserStorage) Filter(_ context.Context, _ string, _ any, order string, offset, limit uint) ([]*models.User, error) {
	users := []*models.User{}
	for _, user := range strg.storage {
		users = append(users, user)
	}
	return users[int(offset):int(limit)], nil
}
func (strg *UserStorage) Paginate(_ context.Context, order string, offset, limit uint) ([]*models.User, error) {
	users := []*models.User{}
	for _, user := range strg.storage {
		users = append(users, user)
	}
	return users[int(offset):int(limit)], nil
}
func (strg *UserStorage) Delete(_ context.Context, id uint8) error {
	return nil
}
func (strg *UserStorage) EmailExists(_ context.Context, email string) bool {
	_, ok := strg.storage[email]
	return ok
}
