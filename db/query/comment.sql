-- name: GetCommentsBySongId :many
SELECT * FROM Comment WHERE song_id = $1;

-- name: CreateComment :exec
INSERT INTO Comment (account_id, song_id, content)
VALUES ($1, $2, $3);

-- name: DeleteComment :exec
DELETE FROM Comment WHERE comment_id = $1;
