package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xmlking/grpc-starter-kit/internal/events"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/hook"
	"github.com/xmlking/grpc-starter-kit/internal/config"
)

// InitDatabase database
// https://github.com/reddydodda/magmatest/blob/master/orc8r/cloud/go/blobstore/ent/client.go
func InitDatabase(dbConf config.DatabaseConfiguration) (client *ent.Client, err error) {

	var url string
	url, err = dbConf.URL()
	if err != nil {
		return
	}

	var db *sql.DB
	db, err = sql.Open(dbConf.Dialect, url)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(dbConf.MaxIdleConns)
	db.SetMaxOpenConns(dbConf.MaxOpenConns)
	db.SetConnMaxLifetime(*dbConf.ConnMaxLifetime)

	// ping test
	err = db.Ping()
	if err != nil {
		return
	}

	// set timezone
	if dbConf.Utc {
		var timezoneCommand string
		switch dbConf.Dialect {
		case dialect.SQLite:
		case dialect.Postgres:
			timezoneCommand = "SET timezone = 'UTC'"
		case dialect.MySQL:
			timezoneCommand = "SET time_zone = '+00:00'"
		default:
			return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
		}
		if _, err = db.Exec(timezoneCommand); err != nil {
			return nil, errors.Wrapf(err, "error setting UTC timezone: %s", timezoneCommand)
		}
	}

	drv := entsql.OpenDB(dbConf.Dialect, db)

	opts := []ent.Option{ent.Driver(drv)}
	if dbConf.Logging {
		opts = append(opts, ent.Debug())
		opts = append(opts, ent.Log(log.Print))
	}
	client = ent.NewClient(opts...)

	// Run Database Setup/Migrations
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msgf("failed creating schema resources")
		return nil, err
	}

	// Add a hook only on update operations for TZPolicy
	// client.TZPolicy.Use(hook.On(events.PolicyChangeEvents(), ent.OpUpdate|ent.OpUpdateOne))
	client.Profile.Use(events.ProfileChangeEvents())
	// Reject delete operations.
	client.Use(hook.Reject(ent.OpDelete | ent.OpDeleteOne))

	return
}
