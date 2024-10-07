-- name: GetRepostsByAccountId :many
SELECT * FROM Repost WHERE account_id = $1;

-- name: GetRepostsBySongId :many
SELECT * FROM Repost WHERE song_id = $1;

-- name: RepostSong :exec
INSERT INTO Repost (account_id, song_id)
VALUES ($1, $2);

-- name: DeleteRepost :exec
DELETE FROM Repost WHERE repost_id = $1;
