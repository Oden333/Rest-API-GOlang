package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_ULR")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:1234@localhost:5433/restapi_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
