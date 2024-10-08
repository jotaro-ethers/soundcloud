// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO Account (username, display_name, email, password, bio, avatar_url, is_verified, role)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING account_id, username, display_name, email, password, bio, avatar_url, is_verified, role, created_at
`

type CreateAccountParams struct {
	Username    string         `json:"username"`
	DisplayName sql.NullString `json:"display_name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Bio         sql.NullString `json:"bio"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	IsVerified  bool           `json:"is_verified"`
	Role        string         `json:"role"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.queryRow(ctx, q.createAccountStmt, createAccount,
		arg.Username,
		arg.DisplayName,
		arg.Email,
		arg.Password,
		arg.Bio,
		arg.AvatarUrl,
		arg.IsVerified,
		arg.Role,
	)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsVerified,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM Account WHERE account_id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID int32) error {
	_, err := q.exec(ctx, q.deleteAccountStmt, deleteAccount, accountID)
	return err
}

const getAccountById = `-- name: GetAccountById :one
SELECT account_id, username, display_name, email, password, bio, avatar_url, is_verified, role, created_at FROM Account WHERE account_id = $1
`

func (q *Queries) GetAccountById(ctx context.Context, accountID int32) (Account, error) {
	row := q.queryRow(ctx, q.getAccountByIdStmt, getAccountById, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsVerified,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getAllAccounts = `-- name: GetAllAccounts :many
SELECT account_id, username, display_name, email, password, bio, avatar_url, is_verified, role, created_at FROM Account
ORDER BY account_id
LIMIT $1
OFFSET $2
`

type GetAllAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllAccounts(ctx context.Context, arg GetAllAccountsParams) ([]Account, error) {
	rows, err := q.query(ctx, q.getAllAccountsStmt, getAllAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.AccountID,
			&i.Username,
			&i.DisplayName,
			&i.Email,
			&i.Password,
			&i.Bio,
			&i.AvatarUrl,
			&i.IsVerified,
			&i.Role,
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

const updateAccount = `-- name: UpdateAccount :exec
UPDATE Account
SET username = $1, display_name = $2,email = $3, password = $4, bio = $5, avatar_url = $6, is_verified = $7
WHERE account_id = $8
`

type UpdateAccountParams struct {
	Username    string         `json:"username"`
	DisplayName sql.NullString `json:"display_name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Bio         sql.NullString `json:"bio"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	IsVerified  bool           `json:"is_verified"`
	AccountID   int32          `json:"account_id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.exec(ctx, q.updateAccountStmt, updateAccount,
		arg.Username,
		arg.DisplayName,
		arg.Email,
		arg.Password,
		arg.Bio,
		arg.AvatarUrl,
		arg.IsVerified,
		arg.AccountID,
	)
	return err
}
