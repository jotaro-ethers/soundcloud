-- name: GetSongsByPlaylistId :many
SELECT s.*
FROM Playlist_Songs ps
JOIN Song s ON ps.song_id = s.song_id
WHERE ps.playlist_id = $1;

-- name: AddSongToPlaylist :exec
INSERT INTO Playlist_Songs (playlist_id, song_id)
VALUES ($1, $2);

-- name: RemoveSongFromPlaylist :exec
DELETE FROM Playlist_Songs
WHERE playlist_id = $1 AND song_id = $2;
