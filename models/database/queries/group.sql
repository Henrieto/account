-- name: CreateUserGroup :one
INSERT INTO groups (
    name,
    created_at,
    updated_at
  )
VALUES($1 , $2 , $3)
RETURNING *;

-- name: CountGroups :one
SELECT count(id) FROM groups;

-- name: GetAllGroups :many
SELECT * FROM groups
ORDER BY $1;

-- name: PaginateGroups :many
SELECT * FROM groups
ORDER BY $1
OFFSET $2
LIMIT $3;


-- name: GetGroup :one
SELECT * FROM groups
WHERE name = $1 LIMIT 1;

-- name: UpdateGroup :one
UPDATE groups
  SET
    name = $1,
    updated_at = $2
WHERE id = $3
RETURNING *;


-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;

-- name: AddUserToGroup :exec
UPDATE users
  SET
    group_id = $1,
    updated_at = $2
WHERE id = $3
RETURNING *;