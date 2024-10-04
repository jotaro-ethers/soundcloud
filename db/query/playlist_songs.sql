-- name: GetPlaylistSongs :many
SELECT * FROM Playlist_Songs WHERE playlist_id = $1;

-- name: AddSongToPlaylist :exec
INSERT INTO Playlist_Songs (playlist_id, song_id)
VALUES ($1, $2);

-- name: RemoveSongFromPlaylist :exec
DELETE FROM Playlist_Songs
WHERE playlist_id = $1 AND song_id = $2;
