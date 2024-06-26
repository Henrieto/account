// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: permisions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addPermissionToGroup = `-- name: AddPermissionToGroup :one
INSERT INTO group_permissions (
    group_id,
    permission_id,
    created_at,
    updated_at
)
VALUES ($1,$2,$3,$4)
RETURNING id, group_id, permission_id, created_at, updated_at
`

type AddPermissionToGroupParams struct {
	GroupID      pgtype.Int4        `json:"group_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) AddPermissionToGroup(ctx context.Context, arg AddPermissionToGroupParams) (GroupPermission, error) {
	row := q.db.QueryRow(ctx, addPermissionToGroup,
		arg.GroupID,
		arg.PermissionID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i GroupPermission
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.PermissionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addPermissionToUser = `-- name: AddPermissionToUser :one
INSERT INTO user_permissions (
    user_id,
    permission_id,
    created_at,
    updated_at
)
VALUES ($1,$2,$3,$4)
RETURNING id, user_id, permission_id, created_at, updated_at
`

type AddPermissionToUserParams struct {
	UserID       pgtype.Int4        `json:"user_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) AddPermissionToUser(ctx context.Context, arg AddPermissionToUserParams) (UserPermission, error) {
	row := q.db.QueryRow(ctx, addPermissionToUser,
		arg.UserID,
		arg.PermissionID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i UserPermission
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PermissionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const countGroupPermissions = `-- name: CountGroupPermissions :one
Select count(id) from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1
`

func (q *Queries) CountGroupPermissions(ctx context.Context, groupID pgtype.Int4) (int64, error) {
	row := q.db.QueryRow(ctx, countGroupPermissions, groupID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPermissions = `-- name: CountPermissions :one
SELECT count(id) FROM permissions
`

func (q *Queries) CountPermissions(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPermissions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUserPermissions = `-- name: CountUserPermissions :one
Select count(id) from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1
`

func (q *Queries) CountUserPermissions(ctx context.Context, userID pgtype.Int4) (int64, error) {
	row := q.db.QueryRow(ctx, countUserPermissions, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPermission = `-- name: CreatePermission :one
INSERT INTO permissions (
    model,
    name,
    codename,
    created_at,
    updated_at
)
VALUES ($1 ,$2 ,$3 ,$4 ,$5 )
RETURNING id, model, name, codename, created_at, updated_at
`

type CreatePermissionParams struct {
	Model     string             `json:"model"`
	Name      string             `json:"name"`
	Codename  string             `json:"codename"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error) {
	row := q.db.QueryRow(ctx, createPermission,
		arg.Model,
		arg.Name,
		arg.Codename,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Name,
		&i.Codename,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteGroupPermission = `-- name: DeleteGroupPermission :exec
DELETE FROM group_permissions
WHERE group_permissions.group_id = $1 
AND group_permissions.permission_id = $2
`

type DeleteGroupPermissionParams struct {
	GroupID      pgtype.Int4 `json:"group_id"`
	PermissionID pgtype.Int4 `json:"permission_id"`
}

func (q *Queries) DeleteGroupPermission(ctx context.Context, arg DeleteGroupPermissionParams) error {
	_, err := q.db.Exec(ctx, deleteGroupPermission, arg.GroupID, arg.PermissionID)
	return err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePermission, id)
	return err
}

const deleteUserPermission = `-- name: DeleteUserPermission :exec
DELETE FROM user_permissions
WHERE user_permissions.user_id = $1 
AND user_permissions.permission_id = $2
`

type DeleteUserPermissionParams struct {
	UserID       pgtype.Int4 `json:"user_id"`
	PermissionID pgtype.Int4 `json:"permission_id"`
}

func (q *Queries) DeleteUserPermission(ctx context.Context, arg DeleteUserPermissionParams) error {
	_, err := q.db.Exec(ctx, deleteUserPermission, arg.UserID, arg.PermissionID)
	return err
}

const getAllGroupPermissions = `-- name: GetAllGroupPermissions :many
Select permissions.id, model, name, codename, permissions.created_at, permissions.updated_at, group_permissions.id, group_id, permission_id, group_permissions.created_at, group_permissions.updated_at from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1
`

type GetAllGroupPermissionsRow struct {
	ID           int32              `json:"id"`
	Model        string             `json:"model"`
	Name         string             `json:"name"`
	Codename     string             `json:"codename"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	ID_2         int32              `json:"id_2"`
	GroupID      pgtype.Int4        `json:"group_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
	UpdatedAt_2  pgtype.Timestamptz `json:"updated_at_2"`
}

func (q *Queries) GetAllGroupPermissions(ctx context.Context, groupID pgtype.Int4) ([]GetAllGroupPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getAllGroupPermissions, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllGroupPermissionsRow{}
	for rows.Next() {
		var i GetAllGroupPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Name,
			&i.Codename,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.GroupID,
			&i.PermissionID,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPermssions = `-- name: GetAllPermssions :many
SELECT id, model, name, codename, created_at, updated_at FROM permissions
ORDER BY $1
OFFSET $2
LIMIT $3
`

type GetAllPermssionsParams struct {
	Column1 interface{} `json:"column_1"`
	Offset  int32       `json:"offset"`
	Limit   int32       `json:"limit"`
}

func (q *Queries) GetAllPermssions(ctx context.Context, arg GetAllPermssionsParams) ([]Permission, error) {
	rows, err := q.db.Query(ctx, getAllPermssions, arg.Column1, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Permission{}
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Name,
			&i.Codename,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUserPermissions = `-- name: GetAllUserPermissions :many
Select permissions.id, model, name, codename, permissions.created_at, permissions.updated_at, user_permissions.id, user_id, permission_id, user_permissions.created_at, user_permissions.updated_at from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1
`

type GetAllUserPermissionsRow struct {
	ID           int32              `json:"id"`
	Model        string             `json:"model"`
	Name         string             `json:"name"`
	Codename     string             `json:"codename"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	ID_2         int32              `json:"id_2"`
	UserID       pgtype.Int4        `json:"user_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
	UpdatedAt_2  pgtype.Timestamptz `json:"updated_at_2"`
}

func (q *Queries) GetAllUserPermissions(ctx context.Context, userID pgtype.Int4) ([]GetAllUserPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getAllUserPermissions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllUserPermissionsRow{}
	for rows.Next() {
		var i GetAllUserPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Name,
			&i.Codename,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.UserID,
			&i.PermissionID,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupPermissions = `-- name: GetGroupPermissions :many
Select permissions.id, model, name, codename, permissions.created_at, permissions.updated_at, group_permissions.id, group_id, permission_id, group_permissions.created_at, group_permissions.updated_at from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1
OFFSET $2
LIMIT $3
`

type GetGroupPermissionsParams struct {
	GroupID pgtype.Int4 `json:"group_id"`
	Offset  int32       `json:"offset"`
	Limit   int32       `json:"limit"`
}

type GetGroupPermissionsRow struct {
	ID           int32              `json:"id"`
	Model        string             `json:"model"`
	Name         string             `json:"name"`
	Codename     string             `json:"codename"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	ID_2         int32              `json:"id_2"`
	GroupID      pgtype.Int4        `json:"group_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
	UpdatedAt_2  pgtype.Timestamptz `json:"updated_at_2"`
}

func (q *Queries) GetGroupPermissions(ctx context.Context, arg GetGroupPermissionsParams) ([]GetGroupPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getGroupPermissions, arg.GroupID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetGroupPermissionsRow{}
	for rows.Next() {
		var i GetGroupPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Name,
			&i.Codename,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.GroupID,
			&i.PermissionID,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPermission = `-- name: GetPermission :one
SELECT id, model, name, codename, created_at, updated_at FROM permissions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPermission(ctx context.Context, id int32) (Permission, error) {
	row := q.db.QueryRow(ctx, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Name,
		&i.Codename,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPermissionByName = `-- name: GetPermissionByName :one
SELECT id, model, name, codename, created_at, updated_at FROM permissions
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPermissionByName(ctx context.Context, name string) (Permission, error) {
	row := q.db.QueryRow(ctx, getPermissionByName, name)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Name,
		&i.Codename,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserPermissions = `-- name: GetUserPermissions :many
Select permissions.id, model, name, codename, permissions.created_at, permissions.updated_at, user_permissions.id, user_id, permission_id, user_permissions.created_at, user_permissions.updated_at from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1
OFFSET $2
LIMIT $3
`

type GetUserPermissionsParams struct {
	UserID pgtype.Int4 `json:"user_id"`
	Offset int32       `json:"offset"`
	Limit  int32       `json:"limit"`
}

type GetUserPermissionsRow struct {
	ID           int32              `json:"id"`
	Model        string             `json:"model"`
	Name         string             `json:"name"`
	Codename     string             `json:"codename"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	ID_2         int32              `json:"id_2"`
	UserID       pgtype.Int4        `json:"user_id"`
	PermissionID pgtype.Int4        `json:"permission_id"`
	CreatedAt_2  pgtype.Timestamptz `json:"created_at_2"`
	UpdatedAt_2  pgtype.Timestamptz `json:"updated_at_2"`
}

func (q *Queries) GetUserPermissions(ctx context.Context, arg GetUserPermissionsParams) ([]GetUserPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getUserPermissions, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPermissionsRow{}
	for rows.Next() {
		var i GetUserPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Name,
			&i.Codename,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.UserID,
			&i.PermissionID,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePermission = `-- name: UpdatePermission :one
UPDATE permissions
  SET
    model = $1,
    name = $2,
    codename = $3,
    updated_at = $4
WHERE id = $5
RETURNING id, model, name, codename, created_at, updated_at
`

type UpdatePermissionParams struct {
	Model     string             `json:"model"`
	Name      string             `json:"name"`
	Codename  string             `json:"codename"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	ID        int32              `json:"id"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error) {
	row := q.db.QueryRow(ctx, updatePermission,
		arg.Model,
		arg.Name,
		arg.Codename,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Name,
		&i.Codename,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
