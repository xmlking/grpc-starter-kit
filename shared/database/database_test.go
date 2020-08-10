package database

import (
	"testing"
	"time"

	"github.com/facebookincubator/ent/dialect"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func TestDatabase(t *testing.T) {
	dur := time.Hour
	_, err := GetDatabaseConnection(config.DatabaseConfiguration{
		Dialect:         dialect.SQLite,
		Host:            "127.0.0.1",
		Port:            3306,
		Username:        "root",
		Password:        "123456",
		Database:        ":memory:",
		MaxOpenConns:    1,
		MaxIdleConns:    1,
		ConnMaxLifetime: &dur,
	})
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}
}
