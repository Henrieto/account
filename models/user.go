package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           pgtype.UUID        `json:"id"`
	Username     string             `json:"username"`
	Email        string             `json:"email"`
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
	GroupID      pgtype.UUID        `json:"group_id"`
}

func (user *User) HasPermission(permission string) bool {
	return user.Superuser.Bool
}

func (user *User) HasPermissions(permissions ...string) (has bool) {
	for _, permsission := range permissions {
		has = user.HasPermission(permsission)
	}
	return
}
