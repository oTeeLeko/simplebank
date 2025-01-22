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
	dbSource = "postgresql://root:password@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testStore Store

func TestMain(m *testing.M) {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	testStore = *NewStore(conn)

	os.Exit(m.Run())
}
