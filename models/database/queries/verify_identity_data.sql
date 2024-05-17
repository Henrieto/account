-- name: CreateVerifyIdentityData :one
INSERT INTO verify_identity_datas (
    random_string,
    identification_type,
    identification_value,
    otp,
    expiry,
    operation_type
)
VALUES($1 , $2 , $3 , $4 , $5 , $6 )
RETURNING *;


-- name: CountVerifyIdentityDatas :one
SELECT count(id) FROM verify_identity_datas;

-- name: GetVerifyIdentityData :one
SELECT * FROM verify_identity_datas
WHERE random_string = $1;

-- name: DeleteVerifyIdentiyData :exec
DELETE FROM groups
WHERE id = $1;

-- name: PaginateVerifyIdentityData :many
SELECT * FROM verify_identity_datas
ORDER BY $1
OFFSET $2
LIMIT $3;
