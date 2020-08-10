# Config

This will configures [configor](https://github.com/xmlking/configor) `configor` via environment variables and load `config.*.yaml` files into a package level struct variable. 
Once loaded, you can use packet level helper methods to retrieve config data.  

## Usage

Customize **Configor** at runtime with Environment Variables 

### Environment Variables

```bash
export CONFIGOR_FILES=/config/config.yaml
# (or) export CONFIGOR_FILES=/config/config.yaml,/config/config.pg.yaml
export CONFIGOR_DEBUG_MODE=true
export CONFIGOR_VERBOSE_MODE=true
export CONFIGOR_SILENT_MODE=true
export CONFIGOR_USE_PKGER=true
export CONFIGOR_ENV=prod

export CONFIGOR_ENV_PREFIX=APP
export APP_FEATURES_TLS_ENABLED=true

# for example
CONFIGOR_SERVICES_GREETER_ENDPOINT=dns:///localhost:8088 ./build/greeter-service
CONFIGOR_ENV_PREFIX=APP APP_SERVICES_GREETER_ENDPOINT=dns:///localhost:8088 ./build/greeter-service
CONFIGOR_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
CONFIGOR_ENV=prod ./build/greeter-service
```

### Examples

Import `shared/config` package. It will be *self-initialized*. 

```golang
import  "github.com/xmlking/grpc-starter-kit/shared/config"
```

Once `config` is initialized, then you can use `github.com/xmlking/grpc-starter-kit/shared/config` package's helper methods retrieve config items.

```go
import (
    _ "github.com/xmlking/grpc-starter-kit/shared/config"
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
	"github.com/xmlking/grpc-starter-kit/shared/config"
)

func TestOverwriteConfigurationWithEnvironmentWithDefaultPrefix(t *testing.T) {
	os.Setenv("CONFIGOR_SERVICES_ACCOUNT_ENDPOINT", "dns:///localhost:8088")
	defer os.Setenv("CONFIGOR_SERVICES_ACCOUNT_ENDPOINT", "")

	var cfg config.Configuration
	config.Configor.Load(&cfg, "/config/config.yaml")

	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(cfg.Services.Account)
	if cfg.Services.Account.Endpoint != "dns:///localhost:8088" {
		t.Errorf("Account Endpoint is %s, want %s", cfg.Services.Account.Endpoint, "dns:///localhost:8088")
	}
}
```

## Test
```
CONFIGOR_DEBUG_MODE=true go test -v ./shared/config/... -count=1
```


