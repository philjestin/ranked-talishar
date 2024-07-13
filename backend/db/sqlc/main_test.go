package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// Special import, tell go formatter to keep it, use blank identifier _
	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable"
)

// By convention TestMain function is the main entry point
// of all unit tests inside 1 specific golang package
// in this case package db
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
