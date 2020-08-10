package database

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/facebookincubator/ent/dialect"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/logger/gormlog"
)

// GetDatabaseConnection return (gorm.DB or error)
func GetDatabaseConnection(dbConf config.DatabaseConfiguration) (db *gorm.DB, err error) {
	var timezoneCommand string

	switch dbConf.Dialect {
	case dialect.SQLite:
		db, err = connection(dbConf)
	case dialect.Postgres:
		timezoneCommand = "SET timezone = 'UTC'"
		db, err = connection(dbConf)
	case dialect.MySQL:
		timezoneCommand = "SET time_zone = '+00:00'"
		db, err = connection(dbConf)
	default:
		return nil, fmt.Errorf("database dialect %s not supported", dbConf.Dialect)
	}

	if err != nil {
		return
	}
	gLogger := log.With().
		Str("module", "gorm").
		Logger()

	db.SetLogger(gormlog.NewGormLogger(gLogger))

	if dbConf.Logging {
		db.Debug()
	}

	db.LogMode(dbConf.Logging)
	db.SingularTable(dbConf.Singularize)
	db.DB().SetMaxOpenConns(int(dbConf.MaxOpenConns))
	db.DB().SetMaxIdleConns(int(dbConf.MaxIdleConns))
	db.DB().SetConnMaxLifetime(*dbConf.ConnMaxLifetime)

	if dbConf.Utc {
		if _, err = db.DB().Exec(timezoneCommand); err != nil {
			return nil, errors.Wrapf(err, "error setting UTC timezone: %s", timezoneCommand)
		}
	}

	return
}

func connection(dbConf config.DatabaseConfiguration) (db *gorm.DB, err error) {
	url, err := dbConf.URL()
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(strings.ToLower(dbConf.Dialect), url)
	return
}
