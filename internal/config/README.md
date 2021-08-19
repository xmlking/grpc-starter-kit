# Config

This will configures [confy](https://github.com/xmlking/toolkit/confy) via environment variables and load `config.*.yml` files into a package level struct variable. 
Once loaded, you can use packet level helper methods to retrieve config data.  

## Usage

Customize **Confy** at runtime with Environment Variables 

### Environment Variables

```bash
export CONFY_FILES=/config/config.yml
# (or) export CONFY_FILES=/config/config.yml,/config/config.pg.yml
export CONFY_DEBUG_MODE=true
export CONFY_VERBOSE_MODE=true
export CONFY_SILENT_MODE=true
export CONFY_USE_PKGER=true
export CONFY_ENV=production

export CONFY_ENV_PREFIX=APP
export APP_FEATURES_TLS_ENABLED=true

# for example
CONFY_SERVICES_GREETER_ENDPOINT=dns:///localhost:8088 ./build/greeter-service
CONFY_ENV_PREFIX=APP APP_SERVICES_GREETER_ENDPOINT=dns:///localhost:8088 ./build/greeter-service
CONFY_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
CONFY_ENV=production ./build/greeter-service
```

### Examples

Import `shared/config` package. It will be *self-initialized*. 

```golang
import  "github.com/xmlking/grpc-starter-kit/internal/config"
```

Once `config` is initialized, then you can use `github.com/xmlking/grpc-starter-kit/internal/config` package's helper methods retrieve config items.

```go
import (
    "github.com/xmlking/grpc-starter-kit/internal/config"
)

func ExampleGetConfig_check_defaults() {
	fmt.Println(config.GetConfig().Services.Account.Endpoint)
	fmt.Println(config.GetConfig().Services.Account.Version)
	fmt.Println(config.GetConfig().Services.Account.Deadline)

	// Output:
	// dns:///account.test:8080
	// v0.1.0
	// 8888
}
```

You can also use `Configor` to load any yaml files into your Struct.

```go
import (
	"github.com/xmlking/grpc-starter-kit/internal/config"
)

func TestOverwriteConfigurationWithEnvironmentWithDefaultPrefix(t *testing.T) {
	os.Setenv("CONFY_SERVICES_ACCOUNT_ENDPOINT", "dns:///localhost:8088")
	defer os.Setenv("CONFY_SERVICES_ACCOUNT_ENDPOINT", "")

	var cfg config.Configuration
	config.Load(&cfg, "/config/config.yml")

	t.Logf("Environment: %s", config.GetEnvironment())
	t.Log(cfg.Services.Account)
	if cfg.Services.Account.Endpoint != "dns:///localhost:8088" {
		t.Errorf("Account Endpoint is %s, want %s", cfg.Services.Account.Endpoint, "dns:///localhost:8088")
	}
}
```

## Test
```
CONFY_DEBUG_MODE=true go test -v ./internal/config/... -count=1
```


