-- name: GetAccountById :one
SELECT * FROM Account WHERE account_id = $1;

-- name: GetAllAccounts :many
SELECT * FROM Account
ORDER BY account_id
LIMIT $1
OFFSET $2
;

-- name: CreateAccount :one
INSERT INTO Account (username, display_name, email, password, bio, avatar_url, is_verified, role)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: UpdateAccount :exec
UPDATE Account
SET username = $1, display_name = $2,email = $3, password = $4, bio = $5, avatar_url = $6, is_verified = $7
WHERE account_id = $8;

-- name: DeleteAccount :exec
DELETE FROM Account WHERE account_id = $1;
