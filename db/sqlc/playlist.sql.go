// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: playlist.sql

package db

import (
	"context"
	"database/sql"
)

const createPlaylist = `-- name: CreatePlaylist :exec
INSERT INTO Playlist (name, account_id, is_public)
VALUES ($1, $2, $3)
`

type CreatePlaylistParams struct {
	Name      string       `json:"name"`
	AccountID int32        `json:"account_id"`
	IsPublic  sql.NullBool `json:"is_public"`
}

func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) error {
	_, err := q.exec(ctx, q.createPlaylistStmt, createPlaylist, arg.Name, arg.AccountID, arg.IsPublic)
	return err
}

const deletePlaylist = `-- name: DeletePlaylist :exec
DELETE FROM Playlist WHERE playlist_id = $1
`

func (q *Queries) DeletePlaylist(ctx context.Context, playlistID int32) error {
	_, err := q.exec(ctx, q.deletePlaylistStmt, deletePlaylist, playlistID)
	return err
}

const getAllPlaylists = `-- name: GetAllPlaylists :many
SELECT playlist_id, name, account_id, is_public, created_at FROM Playlist
`

func (q *Queries) GetAllPlaylists(ctx context.Context) ([]Playlist, error) {
	rows, err := q.query(ctx, q.getAllPlaylistsStmt, getAllPlaylists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.PlaylistID,
			&i.Name,
			&i.AccountID,
			&i.IsPublic,
			&i.CreatedAt,
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

const getPlaylistById = `-- name: GetPlaylistById :one
SELECT playlist_id, name, account_id, is_public, created_at FROM Playlist WHERE playlist_id = $1
`

func (q *Queries) GetPlaylistById(ctx context.Context, playlistID int32) (Playlist, error) {
	row := q.queryRow(ctx, q.getPlaylistByIdStmt, getPlaylistById, playlistID)
	var i Playlist
	err := row.Scan(
		&i.PlaylistID,
		&i.Name,
		&i.AccountID,
		&i.IsPublic,
		&i.CreatedAt,
	)
	return i, err
}

const updatePlaylist = `-- name: UpdatePlaylist :exec
UPDATE Playlist
SET name = $1, is_public = $2
WHERE playlist_id = $3
`

type UpdatePlaylistParams struct {
	Name       string       `json:"name"`
	IsPublic   sql.NullBool `json:"is_public"`
	PlaylistID int32        `json:"playlist_id"`
}

func (q *Queries) UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) error {
	_, err := q.exec(ctx, q.updatePlaylistStmt, updatePlaylist, arg.Name, arg.IsPublic, arg.PlaylistID)
	return err
}
