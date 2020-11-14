package database_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/facebook/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/database"
	_ "github.com/xmlking/grpc-starter-kit/internal/logger"
)

/********************
Unit Tests
********************/
var dbClient *ent.Client
var err error

func TestMain(m *testing.M) {
	dur := time.Hour
	dbClient, err = database.InitDatabase(config.DatabaseConfiguration{
		Dialect:         dialect.SQLite,
		Host:            "127.0.0.1",
		Port:            3306,
		Username:        "root",
		Password:        "123456",
		Database:        "file:account?mode=memory&cache=shared&_fk=1",
		MaxOpenConns:    1,
		MaxIdleConns:    1,
		ConnMaxLifetime: &dur,
	})

	code := m.Run()
	if dbClient != nil {
		_ = dbClient.Close() // cleanup
	}
	os.Exit(code)
}

func TestDatabaseClient(t *testing.T) {
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}

	total, err := dbClient.User.
		Query().
		Count(context.TODO())

	if err != nil {
		t.Fatalf("Database Query failed, %v!", err)
	}
	if total != 0 {
		t.Fatalf("Database Count failed, expected: %d, actual: %d", 0, total)
	}

	t.Logf("Total Users: %d", total)
}
