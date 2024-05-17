package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/henrieto/account/models"
)

func GetUser(object any) (*models.User, error) {
	switch user := object.(type) {
	case *models.User:
		return user, nil
	default:
		return nil, errors.New(" not a user object")
	}
}

type UserKey struct{}

func AddUserToRequest(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(context.Background(), UserKey{}, user)
	r = r.WithContext(ctx)
	return r
}

func GetUserFromRequest(r *http.Request) (*models.User, error) {
	object := r.Context().Value(UserKey{})
	switch user := object.(type) {
	case *models.User:
		return user, nil
	default:
		return nil, errors.New(" not user object ")
	}
}
