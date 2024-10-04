-- name: GetAccountById :one
SELECT * FROM Account WHERE account_id = $1;

-- name: CreateAccount :exec
INSERT INTO Account (username, email, password, bio, avatar_url)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateAccount :exec
UPDATE Account
SET username = $1, email = $2, bio = $3, avatar_url = $4
WHERE account_id = $5;

-- name: DeleteAccount :exec
DELETE FROM Account WHERE account_id = $1;
