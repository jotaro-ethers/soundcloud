// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO Account (username, email, password, bio, avatar_url)
VALUES ($1, $2, $3, $4, $5)
`

type CreateAccountParams struct {
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Bio       sql.NullString `json:"bio"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.exec(ctx, q.createAccountStmt, createAccount,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Bio,
		arg.AvatarUrl,
	)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM Account WHERE account_id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID int32) error {
	_, err := q.exec(ctx, q.deleteAccountStmt, deleteAccount, accountID)
	return err
}

const getAccountById = `-- name: GetAccountById :one
SELECT account_id, username, email, password, bio, avatar_url, created_at FROM Account WHERE account_id = $1
`

func (q *Queries) GetAccountById(ctx context.Context, accountID int32) (Account, error) {
	row := q.queryRow(ctx, q.getAccountByIdStmt, getAccountById, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE Account
SET username = $1, email = $2, bio = $3, avatar_url = $4
WHERE account_id = $5
`

type UpdateAccountParams struct {
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Bio       sql.NullString `json:"bio"`
	AvatarUrl sql.NullString `json:"avatar_url"`
	AccountID int32          `json:"account_id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.exec(ctx, q.updateAccountStmt, updateAccount,
		arg.Username,
		arg.Email,
		arg.Bio,
		arg.AvatarUrl,
		arg.AccountID,
	)
	return err
}
