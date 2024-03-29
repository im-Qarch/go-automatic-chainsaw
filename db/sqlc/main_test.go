package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"db.sqlc.dev/app/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load the config")
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
