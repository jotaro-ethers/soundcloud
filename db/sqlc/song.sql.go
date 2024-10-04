// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: song.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSong = `-- name: CreateSong :exec
INSERT INTO Song (title, duration, file_url, account_id, genre, upload_date)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateSongParams struct {
	Title      string         `json:"title"`
	Duration   time.Time      `json:"duration"`
	FileUrl    string         `json:"file_url"`
	AccountID  sql.NullInt32  `json:"account_id"`
	Genre      sql.NullString `json:"genre"`
	UploadDate sql.NullTime   `json:"upload_date"`
}

func (q *Queries) CreateSong(ctx context.Context, arg CreateSongParams) error {
	_, err := q.exec(ctx, q.createSongStmt, createSong,
		arg.Title,
		arg.Duration,
		arg.FileUrl,
		arg.AccountID,
		arg.Genre,
		arg.UploadDate,
	)
	return err
}

const deleteSong = `-- name: DeleteSong :exec
DELETE FROM Song WHERE song_id = $1
`

func (q *Queries) DeleteSong(ctx context.Context, songID int32) error {
	_, err := q.exec(ctx, q.deleteSongStmt, deleteSong, songID)
	return err
}

const getSongById = `-- name: GetSongById :one
SELECT song_id, title, duration, file_url, account_id, genre, upload_date, play_count, like_count FROM Song WHERE song_id = $1
`

func (q *Queries) GetSongById(ctx context.Context, songID int32) (Song, error) {
	row := q.queryRow(ctx, q.getSongByIdStmt, getSongById, songID)
	var i Song
	err := row.Scan(
		&i.SongID,
		&i.Title,
		&i.Duration,
		&i.FileUrl,
		&i.AccountID,
		&i.Genre,
		&i.UploadDate,
		&i.PlayCount,
		&i.LikeCount,
	)
	return i, err
}

const updateSong = `-- name: UpdateSong :exec
UPDATE Song
SET title = $1, duration = $2, genre = $3, file_url = $4
WHERE song_id = $5
`

type UpdateSongParams struct {
	Title    string         `json:"title"`
	Duration time.Time      `json:"duration"`
	Genre    sql.NullString `json:"genre"`
	FileUrl  string         `json:"file_url"`
	SongID   int32          `json:"song_id"`
}

func (q *Queries) UpdateSong(ctx context.Context, arg UpdateSongParams) error {
	_, err := q.exec(ctx, q.updateSongStmt, updateSong,
		arg.Title,
		arg.Duration,
		arg.Genre,
		arg.FileUrl,
		arg.SongID,
	)
	return err
}
