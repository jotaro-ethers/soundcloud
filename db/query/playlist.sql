-- name: GetPlaylistById :one
SELECT * FROM Playlist WHERE playlist_id = $1;

-- name: CreatePlaylist :exec
INSERT INTO Playlist (name, account_id)
VALUES ($1, $2);

-- name: UpdatePlaylist :exec
UPDATE Playlist
SET name = $1
WHERE playlist_id = $2;

-- name: DeletePlaylist :exec
DELETE FROM Playlist WHERE playlist_id = $1;
