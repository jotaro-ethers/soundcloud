-- name: GetPlaylistById :one
SELECT * FROM Playlist WHERE playlist_id = $1;

-- name: GetAllPlaylists :many
SELECT * FROM Playlist;

-- name: CreatePlaylist :exec
INSERT INTO Playlist (name, account_id, is_public)
VALUES ($1, $2, $3);

-- name: UpdatePlaylist :exec
UPDATE Playlist
SET name = $1, is_public = $2
WHERE playlist_id = $3;

-- name: DeletePlaylist :exec
DELETE FROM Playlist WHERE playlist_id = $1;
