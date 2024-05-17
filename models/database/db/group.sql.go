// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: group.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUserToGroup = `-- name: AddUserToGroup :exec
UPDATE users
  SET
    group_id = $1,
    updated_at = $2
WHERE id = $3
RETURNING id, username, email, phone, first_name, last_name, gender, password_hash, verified, birthday, staff, superuser, auth_id, created_at, updated_at, group_id
`

type AddUserToGroupParams struct {
	GroupID   pgtype.Int4        `json:"group_id"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	ID        int32              `json:"id"`
}

func (q *Queries) AddUserToGroup(ctx context.Context, arg AddUserToGroupParams) error {
	_, err := q.db.Exec(ctx, addUserToGroup, arg.GroupID, arg.UpdatedAt, arg.ID)
	return err
}

const countGroups = `-- name: CountGroups :one
SELECT count(id) FROM groups
`

func (q *Queries) CountGroups(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countGroups)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUserGroup = `-- name: CreateUserGroup :one
INSERT INTO groups (
    name,
    created_at,
    updated_at
  )
VALUES($1 , $2 , $3)
RETURNING id, name, created_at, updated_at
`

type CreateUserGroupParams struct {
	Name      string             `json:"name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateUserGroup(ctx context.Context, arg CreateUserGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, createUserGroup, arg.Name, arg.CreatedAt, arg.UpdatedAt)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGroup, id)
	return err
}

const getAllGroups = `-- name: GetAllGroups :many
SELECT id, name, created_at, updated_at FROM groups
ORDER BY $1
`

func (q *Queries) GetAllGroups(ctx context.Context, dollar_1 interface{}) ([]Group, error) {
	rows, err := q.db.Query(ctx, getAllGroups, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Group{}
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const getGroup = `-- name: GetGroup :one
SELECT id, name, created_at, updated_at FROM groups
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetGroup(ctx context.Context, name string) (Group, error) {
	row := q.db.QueryRow(ctx, getGroup, name)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const paginateGroups = `-- name: PaginateGroups :many
SELECT id, name, created_at, updated_at FROM groups
ORDER BY $1
OFFSET $2
LIMIT $3
`

type PaginateGroupsParams struct {
	Column1 interface{} `json:"column_1"`
	Offset  int32       `json:"offset"`
	Limit   int32       `json:"limit"`
}

func (q *Queries) PaginateGroups(ctx context.Context, arg PaginateGroupsParams) ([]Group, error) {
	rows, err := q.db.Query(ctx, paginateGroups, arg.Column1, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Group{}
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateGroup = `-- name: UpdateGroup :one
UPDATE groups
  SET
    name = $1,
    updated_at = $2
WHERE id = $3
RETURNING id, name, created_at, updated_at
`

type UpdateGroupParams struct {
	Name      string             `json:"name"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	ID        int32              `json:"id"`
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, updateGroup, arg.Name, arg.UpdatedAt, arg.ID)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
