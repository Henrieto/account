package repository

import (
	"context"

	"github.com/henrieto/account/models"
)

type IUserRepository interface {
	Create(context.Context, *models.User) (*models.User, error)
	Update(context.Context, *models.User) (*models.User, error)
	List(context.Context, uint, uint) ([]*models.User, error)
	Get(context.Context, string) (*models.User, error)
	Filter(context.Context, string, any) (*models.User, error)
	Delete(context.Context, string) error
	EmailExists(context.Context, string) bool
}

var UserRepository IUserRepository
