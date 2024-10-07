-- name: GetSongById :one
SELECT * FROM Song WHERE song_id = $1;

-- name: GetAllSongs :many
SELECT * FROM Song;

-- name: CreateSong :exec
INSERT INTO Song (title, duration, file_url, account_id, album_id, genre, tags, upload_date, play_count, like_count, repost_count, is_public)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);

-- name: UpdateSong :exec
UPDATE Song
SET title = $1, duration = $2, file_url = $3, genre = $4, tags = $5, is_public = $6
WHERE song_id = $7;

-- name: DeleteSong :exec
DELETE FROM Song WHERE song_id = $1;
