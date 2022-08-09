package database

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/samber/do"
	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/internal/config"
)

type Client interface {
	do.Shutdownable
	do.Healthcheckable
	GetEntClient() *ent.Client
}

type entClient struct {
	db  *sql.DB
	ent *ent.Client
}

func NewClient(i *do.Injector) (Client, error) {
	cfg, err := do.Invoke[config.Configuration](i)
	if err != nil {
		return nil, err
	}
	entC, db, err := InitDatabase(*cfg.Database)
	if err != nil {
		return nil, err
	}
	return &entClient{db, entC}, nil
}

func (c *entClient) HealthCheck() error {
	log.Trace().Str("component", "database.Client").Msg("Checking health...")
	return c.db.Ping()
}

func (c *entClient) Shutdown() error {
	log.Info().Str("component", "database.Client").Msg("Shutting down...")
	return c.ent.Close()
}

func (c *entClient) GetEntClient() *ent.Client {
	return c.ent
}
