-- name: GetLikesBySongId :many
SELECT * FROM "Like" WHERE song_id = $1;

-- name: LikeSong :exec
INSERT INTO "Like" (account_id, song_id)
VALUES ($1, $2);

-- name: UnlikeSong :exec
DELETE FROM "Like" WHERE account_id = $1 AND song_id = $2;
