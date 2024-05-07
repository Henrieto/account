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
    staff,
    created_at,
    updated_at
)
VALUES( $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 , $9 , $10 , $11)
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
    superuser,
    created_at,
    updated_at
)
VALUES( $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 , $9 , $10 , $11)
RETURNING *;


-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY $1;


-- name: PaginateUsers :many
SELECT * FROM users
ORDER BY $1
OFFSET $2
LIMIT $3;


-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;


-- name: UpdateUser :one
UPDATE users
  SET
    username = $1,
    first_name = $2,
    last_name = $3,
    gender = $4,
    birthday = $5,
    updated_at = $6
WHERE id = $7
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
  SET
    password_hash = $1,
    updated_at = $2
WHERE id = $3
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
  SET
    email = $1,
    updated_at = $2
WHERE id = $3
RETURNING *;


-- name: MakeUserVerified :one
UPDATE users
  SET
    verified = true,
    updated_at = $1
WHERE id = $2
RETURNING *;


-- name: SetUserAuthId :one
UPDATE users
  SET
    auth_id = $1,
    updated_at = $2
WHERE id = $3
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;