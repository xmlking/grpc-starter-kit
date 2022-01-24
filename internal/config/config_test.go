package config_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/toolkit/confy"
	_ "github.com/xmlking/toolkit/logger/auto"
)

// CONFY_DEBUG_MODE=true go test -v ./internal/config/... -count=1

func setup() {
	// HINT: CertFile etc., Our schema has `file` path validation, which is relative to project root.
	if err := os.Chdir("../.."); err != nil {
		log.Fatal().Err(err).Send()
	}
	wd, _ := os.Getwd()
	log.Debug().Msgf("Setup: changing working directory to: %s", wd)
	log.Debug().Msg("Setup completed")
}

func teardown() {
	// Do something here.
	//confy.DefaultConfy = nil
	fmt.Println("Teardown completed")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestNestedConfig(t *testing.T) {
	t.Log(config.GetConfig().Database)
	t.Logf("Environment: %s", confy.GetEnvironment())
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func TestDefaultValues(t *testing.T) {
	t.Log(config.GetConfig().Database)
	t.Logf("Environment: %s", confy.GetEnvironment())
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func ExampleGetConfig() {
	fmt.Println(config.GetConfig().Email)
	// fmt.Println(config.GetConfig().Services.Account.Authority)

	// Output:
	// &{yourGmailUsername yourGmailAppPassword smtp.gmail.com 587 from-test@gmail.com}
}

func ExampleGetConfig_check_defaults() {
	fmt.Println(config.GetConfig().Services.Account.Endpoint)
	fmt.Println(config.GetConfig().Services.Account.Version)
	fmt.Println(config.GetConfig().Services.Account.Authority)

	// Output:
	// dns:///account.test:8080
	// v0.1.0
	// aaa.bbb.ccc
}

func TestOverwriteConfigurationWithEnvironmentWithDefaultPrefix(t *testing.T) {
	t.Setenv("CONFY_SERVICES_ACCOUNT_ENDPOINT", "dns:///localhost:8088")

	var cfg config.Configuration
	confy.DefaultConfy = confy.NewConfy(confy.WithFS(config.GetFileSystem()), confy.WithDebugMode())
	err := confy.Load(&cfg, "config/config.yml")
	assert.NoError(t, err)

	t.Logf("Environment: %s", confy.GetEnvironment())
	t.Log(cfg)
	t.Log(cfg.Services.Account)
	if cfg.Services.Account.Endpoint != "dns:///localhost:8088" {
		t.Errorf("Account Endpoint is %s, want %s", cfg.Services.Account.Endpoint, "dns:///localhost:8088")
	}
}
