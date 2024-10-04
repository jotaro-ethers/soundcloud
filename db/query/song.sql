-- name: GetSongById :one
SELECT * FROM Song WHERE song_id = $1;

-- name: CreateSong :exec
INSERT INTO Song (title, duration, file_url, account_id, genre, upload_date)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateSong :exec
UPDATE Song
SET title = $1, duration = $2, genre = $3, file_url = $4
WHERE song_id = $5;

-- name: DeleteSong :exec
DELETE FROM Song WHERE song_id = $1;
