package config_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/xmlking/toolkit/confy"

	"github.com/stretchr/testify/assert"

	"github.com/xmlking/grpc-starter-kit/internal/config"
)

// CONFY_DEBUG_MODE=true go test -v ./internal/config/... -count=1

func TestNestedConfig(t *testing.T) {
	t.Logf("Environment: %s", confy.GetEnvironment())
	t.Log(config.GetConfig().Database)
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func TestDefaultValues(t *testing.T) {
	t.Logf("Environment: %s", confy.GetEnvironment())
	t.Log(config.GetConfig().Database)
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
	os.Setenv("CONFY_SERVICES_ACCOUNT_ENDPOINT", "dns:///localhost:8088")
	defer os.Clearenv()

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
