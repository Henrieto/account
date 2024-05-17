-- name: CreatePermission :one
INSERT INTO permissions (
    model,
    name,
    codename,
    created_at,
    updated_at
)
VALUES ($1 ,$2 ,$3 ,$4 ,$5 )
RETURNING *;


-- name: CountPermissions :one
SELECT count(id) FROM permissions;

-- name: GetAllPermssions :many
SELECT * FROM permissions
ORDER BY $1
OFFSET $2
LIMIT $3;

-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: GetPermissionByName :one
SELECT * FROM permissions
WHERE name = $1 LIMIT 1;

-- name: UpdatePermission :one
UPDATE permissions
  SET
    model = $1,
    name = $2,
    codename = $3,
    updated_at = $4
WHERE id = $5
RETURNING *;


-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1;


-- name: AddPermissionToUser :one
INSERT INTO user_permissions (
    user_id,
    permission_id,
    created_at,
    updated_at
)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: CountUserPermissions :one
Select count(id) from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1;

-- name: DeleteUserPermission :exec
DELETE FROM user_permissions
WHERE user_permissions.user_id = $1 
AND user_permissions.permission_id = $2;


-- name: GetUserPermissions :many
Select * from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1
OFFSET $2
LIMIT $3;

-- name: GetAllUserPermissions :many
Select * from permissions
Join user_permissions on permissions.id = user_permissions.permission_id
WHERE user_permissions.user_id = $1;

-- name: AddPermissionToGroup :one
INSERT INTO group_permissions (
    group_id,
    permission_id,
    created_at,
    updated_at
)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: CountGroupPermissions :one
Select count(id) from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1;

-- name: GetGroupPermissions :many
Select * from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1
OFFSET $2
LIMIT $3;

-- name: GetAllGroupPermissions :many
Select * from permissions
Join group_permissions on permissions.id = group_permissions.permission_id
WHERE group_permissions.group_id = $1;

-- name: DeleteGroupPermission :exec
DELETE FROM group_permissions
WHERE group_permissions.group_id = $1 
AND group_permissions.permission_id = $2;
