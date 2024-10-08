// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: listen_history.sql

package db

import (
	"context"
)

const addListenHistory = `-- name: AddListenHistory :exec
INSERT INTO Listen_History (account_id, song_id)
VALUES ($1, $2)
`

type AddListenHistoryParams struct {
	AccountID int32 `json:"account_id"`
	SongID    int32 `json:"song_id"`
}

func (q *Queries) AddListenHistory(ctx context.Context, arg AddListenHistoryParams) error {
	_, err := q.exec(ctx, q.addListenHistoryStmt, addListenHistory, arg.AccountID, arg.SongID)
	return err
}

const getListenHistoryByAccountId = `-- name: GetListenHistoryByAccountId :many
SELECT account_id, song_id, listened_at FROM Listen_History WHERE account_id = $1
`

func (q *Queries) GetListenHistoryByAccountId(ctx context.Context, accountID int32) ([]ListenHistory, error) {
	rows, err := q.query(ctx, q.getListenHistoryByAccountIdStmt, getListenHistoryByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListenHistory{}
	for rows.Next() {
		var i ListenHistory
		if err := rows.Scan(&i.AccountID, &i.SongID, &i.ListenedAt); err != nil {
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
