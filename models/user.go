package models

import (
	"github.com/henrieto/account/models/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           int32              `json:"id"`
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	Phone        string             `json:"phone"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	Gender       string             `json:"gender"`
	PasswordHash string             `json:"-"`
	Verified     pgtype.Bool        `json:"verified"`
	Birthday     pgtype.Timestamptz `json:"birthday"`
	Staff        pgtype.Bool        `json:"staff"`
	Superuser    pgtype.Bool        `json:"superuser"`
	AuthID       pgtype.Text        `json:"auth_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	GroupID      pgtype.Int4        `json:"group_id"`
	Permissions  []*db.Permission
}

func (user *User) HasPermission(permission string) bool {
	// check if the user is a super user
	if user.Superuser.Bool {
		// if the user is a super user return true
		return true
	}
	// iterate through the user permissions , to see if the user has the required permission
	for _, user_permission := range user.Permissions {
		if permission == user_permission.Codename {
			// if the user has the required permission return true
			return true
		}
	}
	return false
}

func (user *User) HasPermissions(permissions ...string) (has_permission bool) {
	// iterate to check if the user has the permissions required
	for _, permsission := range permissions {
		has_permission = user.HasPermission(permsission)
		if has_permission {
			return true
		}
	}
	return
}

func (user *User) HasAllPermissions(permissions ...string) (has_permission bool) {
	// iterate to check if the user has all the permissions required
	for _, permsission := range permissions {
		has_permission = user.HasPermission(permsission)
		if !has_permission {
			return false
		}
	}
	return
}
