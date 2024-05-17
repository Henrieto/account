-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    first_name,
    last_name,
    gender,
    password_hash,
    verified,
    birthday,
    group_id,
    created_at,
    updated_at
)
VALUES( $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 , $9 , $10 , $11)
RETURNING *;


-- name: CountUsers :one
SELECT count(id) FROM users;

-- name: CreateStaff :one
INSERT INTO users (
    username,
    email,
    first_name,
    last_name,
    gender,
    password_hash,
    verified,
    birthday,
    group_id,
    staff,
    created_at,
    updated_at
)
VALUES( $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 , $9 , $10 , $11 , $12)
RETURNING *;

-- name: CreateSuperUser :one
INSERT INTO users (
    username,
    email,
    first_name,
    last_name,
    gender,
    password_hash,
    verified,
    birthday,
    group_id,
    superuser,
    created_at,
    updated_at
)
VALUES( $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 , $9 , $10 , $11 , $12)
RETURNING *;


-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY $1;


-- name: PaginateUsers :many
SELECT * FROM users
ORDER BY $1
OFFSET $2
LIMIT $3;


-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users 
WHERE phone = $1 LIMIT 1;

-- name: GetUserByAuthId :one
SELECT * FROM users 
WHERE auth_id = $1 LIMIT 1;

-- name: FilterUsers :many
SELECT * FROM users 
WHERE $1 = $2 
ORDER BY $3
OFFSET $4 
LIMIT $5;


-- name: UpdateUser :one
UPDATE users
  SET
    username = $1,
    email = $2,
    first_name = $3,
    last_name = $4,
    gender = $5,
    password_hash = $6,
    verified = $7,
    birthday = $8,
    staff = $9,
    superuser = $10,
    auth_id = $11,
    group_id = $12,
    updated_at = $13
WHERE id = $14
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;