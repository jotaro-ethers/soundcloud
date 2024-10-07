package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
	
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@172.17.0.2:5432/soundcloud?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB
var testAccount1 Account
var testAccount2 Account


func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
