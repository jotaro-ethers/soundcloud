// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: playlist_songs.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const addSongToPlaylist = `-- name: AddSongToPlaylist :exec
INSERT INTO Playlist_Songs (playlist_id, song_id)
VALUES ($1, $2)
`

type AddSongToPlaylistParams struct {
	PlaylistID int32 `json:"playlist_id"`
	SongID     int32 `json:"song_id"`
}

func (q *Queries) AddSongToPlaylist(ctx context.Context, arg AddSongToPlaylistParams) error {
	_, err := q.exec(ctx, q.addSongToPlaylistStmt, addSongToPlaylist, arg.PlaylistID, arg.SongID)
	return err
}

const getSongsByPlaylistId = `-- name: GetSongsByPlaylistId :many
SELECT s.song_id, s.title, s.duration, s.file_url, s.account_id, s.album_id, s.genre, s.tags, s.upload_date, s.play_count, s.like_count, s.repost_count, s.is_public
FROM Playlist_Songs ps
JOIN Song s ON ps.song_id = s.song_id
WHERE ps.playlist_id = $1
`

func (q *Queries) GetSongsByPlaylistId(ctx context.Context, playlistID int32) ([]Song, error) {
	rows, err := q.query(ctx, q.getSongsByPlaylistIdStmt, getSongsByPlaylistId, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Song{}
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.SongID,
			&i.Title,
			&i.Duration,
			&i.FileUrl,
			&i.AccountID,
			&i.AlbumID,
			&i.Genre,
			pq.Array(&i.Tags),
			&i.UploadDate,
			&i.PlayCount,
			&i.LikeCount,
			&i.RepostCount,
			&i.IsPublic,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeSongFromPlaylist = `-- name: RemoveSongFromPlaylist :exec
DELETE FROM Playlist_Songs
WHERE playlist_id = $1 AND song_id = $2
`

type RemoveSongFromPlaylistParams struct {
	PlaylistID int32 `json:"playlist_id"`
	SongID     int32 `json:"song_id"`
}

func (q *Queries) RemoveSongFromPlaylist(ctx context.Context, arg RemoveSongFromPlaylistParams) error {
	_, err := q.exec(ctx, q.removeSongFromPlaylistStmt, removeSongFromPlaylist, arg.PlaylistID, arg.SongID)
	return err
}
