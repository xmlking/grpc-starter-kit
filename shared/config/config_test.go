package config_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	configPB "github.com/xmlking/grpc-starter-kit/shared/proto/config"
)

// CONFIGOR_DEBUG_MODE=true go test -v ./shared/config/... -count=1

func TestNestedConfig(t *testing.T) {
	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(config.GetConfig().Database)
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func TestDefaultValues(t *testing.T) {
	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(config.GetConfig().Database)
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func ExampleGetConfig() {
	fmt.Println(config.GetConfig().Email)
	// fmt.Println(config.GetConfig().Services["account"].Deadline)

	// Output:
	// username:"yourGmailUsername" password:"yourGmailAppPassword" email_server:"smtp.gmail.com" port:587 from:"from-test@gmail.com"
}

func ExampleGetConfig_check_defaults() {
	fmt.Println(config.GetConfig().Services.Account.Endpoint)
	fmt.Println(config.GetConfig().Services.Account.Version)
	fmt.Println(config.GetConfig().Services.Account.Deadline)

	// Output:
	// dns:///account.test:8080
	// v0.1.0
	// 8888
}

func TestParseTargetString(t *testing.T) {
	for _, test := range []struct {
		targetStr string
		want      config.Target
	}{
		{targetStr: "", want: config.Target{Scheme: "", Host: "", Port: "", Path: ""}},
		{targetStr: "dns:///google.com:8080", want: config.Target{Scheme: "dns", Host: "google.com", Port: "8080", Path: ""}},
		{targetStr: "dns:///google.com", want: config.Target{Scheme: "dns", Host: "google.com", Port: "", Path: ""}},
		{targetStr: "dns:///google.com/?a=b", want: config.Target{Scheme: "dns", Host: "google.com", Port: "", Path: "/"}},
		{targetStr: "https://www.server.com:9999", want: config.Target{Scheme: "https", Host: "www.server.com", Port: "9999", Path: ""}},
		{targetStr: "/unix/socket/address", want: config.Target{Scheme: "", Host: "", Port: "", Path: "/unix/socket/address"}},
		{targetStr: "unix:///tmp/mysrv.sock", want: config.Target{Scheme: "unix", Host: "", Port: "", Path: "/tmp/mysrv.sock"}},
	} {
		got, err := config.ParseTarget(test.targetStr)
		if err != nil {
			t.Error(err)
		}
		if got != test.want {
			t.Errorf("ParseTarget(%q) = %+v, want %+v", test.targetStr, got, test.want)
		}
	}
}

func TestOverwriteConfigurationWithEnvironmentWithDefaultPrefix(t *testing.T) {
	os.Setenv("CONFIGOR_SERVICES_ACCOUNT_ENDPOINT", "dns:///localhost:8088")
	defer os.Setenv("CONFIGOR_SERVICES_ACCOUNT_ENDPOINT", "")

	var cfg configPB.Configuration
	config.Configor.Load(&cfg, "/config/config.yaml")

	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(cfg.Services.Account)
	if cfg.Services.Account.Endpoint != "dns:///localhost:8088" {
		t.Errorf("Account Endpoint is %s, want %s", cfg.Services.Account.Endpoint, "dns:///localhost:8088")
	}
}
