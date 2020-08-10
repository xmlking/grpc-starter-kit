package database

import (
	"database/sql"

	entsql "github.com/facebookincubator/ent/dialect/sql"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/shared/config"
)

// https://github.com/reddydodda/magmatest/blob/master/orc8r/cloud/go/blobstore/ent/client.go
func GetDatabaseConnection2(dbConf config.DatabaseConfiguration) (client *ent.Client, err error) {

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

	err = db.Ping()
	if err != nil {
		return
	}

	drv := entsql.OpenDB(dbConf.Dialect, db)

	opts := []ent.Option{ent.Driver(drv)}
	//opts = append(opts, ent.Debug())
	//opts = append(opts, ent.Log())
	client = ent.NewClient(opts...)
	return
}
