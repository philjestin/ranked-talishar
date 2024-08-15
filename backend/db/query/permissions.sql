-- name: GetAllPermissionsForUser :many
SELECT permissions.code
FROM permissions
INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
INNER JOIN users on users_permissions.user_id = users.user_id
WHERE users.user_id = $1;

-- name: AddPermissionForUser :exec
INSERT INTO users_permissions (user_id, permission_id)
SELECT $1, permissions.id FROM permissions where permissions.code = ANY($2);