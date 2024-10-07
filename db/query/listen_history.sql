-- name: GetListenHistoryByAccountId :many
SELECT * FROM Listen_History WHERE account_id = $1;

-- name: AddListenHistory :exec
INSERT INTO Listen_History (account_id, song_id)
VALUES ($1, $2);
