package db

import (
	"context"
	"database/sql"
	"testing"
    "github.com/jotaro-ethers/soundcloud/utils"

)

func CreateRandomAccount(t *testing.T, q *Queries) Account {
	t.Helper()
	accountParams := CreateAccountParams{
		Username:    utils.RandomString(10),
		DisplayName: sql.NullString{String: utils.RandomString(10), Valid: true},
		Email:       utils.RandomString(10) + "@example.com",
		Password:    utils.RandomString(20),
		Bio:         sql.NullString{String: utils.RandomString(50), Valid: true},
		AvatarUrl:   sql.NullString{String: "https://example.com/avatar.jpg", Valid: true},
	}

	account, err := q.CreateAccount(context.TODO(), accountParams)
	if err != nil {
		t.Fatalf("Không thể tạo tài khoản ngẫu nhiên: %v", err)
	}

	return account
}
