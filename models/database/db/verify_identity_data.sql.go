// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: verify_identity_data.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countVerifyIdentityDatas = `-- name: CountVerifyIdentityDatas :one
SELECT count(id) FROM verify_identity_datas
`

func (q *Queries) CountVerifyIdentityDatas(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countVerifyIdentityDatas)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createVerifyIdentityData = `-- name: CreateVerifyIdentityData :one
INSERT INTO verify_identity_datas (
    random_string,
    identification_type,
    identification_value,
    otp,
    expiry,
    operation_type
)
VALUES($1 , $2 , $3 , $4 , $5 , $6 )
RETURNING id, random_string, identification_type, identification_value, otp, expiry, operation_type
`

type CreateVerifyIdentityDataParams struct {
	RandomString        string             `json:"random_string"`
	IdentificationType  string             `json:"identification_type"`
	IdentificationValue string             `json:"identification_value"`
	Otp                 string             `json:"otp"`
	Expiry              pgtype.Timestamptz `json:"expiry"`
	OperationType       string             `json:"operation_type"`
}

func (q *Queries) CreateVerifyIdentityData(ctx context.Context, arg CreateVerifyIdentityDataParams) (VerifyIdentityData, error) {
	row := q.db.QueryRow(ctx, createVerifyIdentityData,
		arg.RandomString,
		arg.IdentificationType,
		arg.IdentificationValue,
		arg.Otp,
		arg.Expiry,
		arg.OperationType,
	)
	var i VerifyIdentityData
	err := row.Scan(
		&i.ID,
		&i.RandomString,
		&i.IdentificationType,
		&i.IdentificationValue,
		&i.Otp,
		&i.Expiry,
		&i.OperationType,
	)
	return i, err
}

const deleteVerifyIdentiyData = `-- name: DeleteVerifyIdentiyData :exec
DELETE FROM groups
WHERE id = $1
`

func (q *Queries) DeleteVerifyIdentiyData(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteVerifyIdentiyData, id)
	return err
}

const getVerifyIdentityData = `-- name: GetVerifyIdentityData :one
SELECT id, random_string, identification_type, identification_value, otp, expiry, operation_type FROM verify_identity_datas
WHERE random_string = $1
`

func (q *Queries) GetVerifyIdentityData(ctx context.Context, randomString string) (VerifyIdentityData, error) {
	row := q.db.QueryRow(ctx, getVerifyIdentityData, randomString)
	var i VerifyIdentityData
	err := row.Scan(
		&i.ID,
		&i.RandomString,
		&i.IdentificationType,
		&i.IdentificationValue,
		&i.Otp,
		&i.Expiry,
		&i.OperationType,
	)
	return i, err
}

const paginateVerifyIdentityData = `-- name: PaginateVerifyIdentityData :many
SELECT id, random_string, identification_type, identification_value, otp, expiry, operation_type FROM verify_identity_datas
ORDER BY $1
OFFSET $2
LIMIT $3
`

type PaginateVerifyIdentityDataParams struct {
	Column1 interface{} `json:"column_1"`
	Offset  int32       `json:"offset"`
	Limit   int32       `json:"limit"`
}

func (q *Queries) PaginateVerifyIdentityData(ctx context.Context, arg PaginateVerifyIdentityDataParams) ([]VerifyIdentityData, error) {
	rows, err := q.db.Query(ctx, paginateVerifyIdentityData, arg.Column1, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VerifyIdentityData{}
	for rows.Next() {
		var i VerifyIdentityData
		if err := rows.Scan(
			&i.ID,
			&i.RandomString,
			&i.IdentificationType,
			&i.IdentificationValue,
			&i.Otp,
			&i.Expiry,
			&i.OperationType,
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