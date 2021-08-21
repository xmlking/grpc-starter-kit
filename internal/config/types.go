package config

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect"

	"github.com/xmlking/toolkit/telemetry"
)

// Service struct
type Service struct {
	Endpoint      string            `yaml:"endpoint" required:"true"`
	Version       string            `yaml:",omitempty" default:"v0.1.0"`
	Metadata      map[string]string `yaml:",omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ServiceConfig string            `yaml:"service_config,omitempty"`
	Authority     string            `yaml:",omitempty"`
}

// EmailConfiguration struct
type EmailConfiguration struct {
	Username    string `yaml:"username"`
	Password    string `yaml:",omitempty"`
	EmailServer string `yaml:"email_server,omitempty"`
	Port        uint32 `yaml:",omitempty" default:"587"`
	From        string `yaml:",omitempty" validate:"email"`
}

// DatabaseConfiguration struct
type DatabaseConfiguration struct {
	Dialect         string         `yaml:",omitempty" validate:"oneof=mysql sqlite3 postgres gremlin" default:"sqlite3"`
	Host            string         `yaml:",omitempty" validate:"required"`
	Port            uint32         `yaml:",omitempty" default:"5432" validate:"number"`
	Username        string         `yaml:"username,omitempty" validate:"alphanum,required"`
	Password        string         `yaml:"password,omitempty" validate:"alphanum,required"`
	Database        string         `yaml:"database,omitempty" validate:"required"`
	Charset         string         `yaml:",omitempty" default:"utf8"`
	Utc             bool           `yaml:",omitempty" default:"true"`
	Logging         bool           `yaml:",omitempty" default:"false"`
	Singularize     bool           `yaml:",omitempty" default:"false"`
	MaxOpenConns    int            `yaml:"max_open_conns,omitempty" default:"100"`
	MaxIdleConns    int            `yaml:"max_idle_conns,omitempty" default:"10"`
	ConnMaxLifetime *time.Duration `yaml:"conn_max_lifetime,omitempty" default:"1h"`
}

// URL returns a connection string for the database.
func (d *DatabaseConfiguration) URL() (url string, err error) {
	switch d.Dialect {
	case dialect.SQLite:
		return d.Database, nil
	case dialect.Postgres:
		return fmt.Sprintf(
			"host=%s port=%v user=%s dbname=%s sslmode=disable password=%s",
			d.Host, d.Port, d.Username, d.Database, d.Password,
		), nil
	case dialect.MySQL:
		return fmt.Sprintf(
			"%s:%s@(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset,
		), nil
	default:
		return "", fmt.Errorf(" '%v' driver doesn't exist. ", d.Dialect)
	}
}

// Features struct
type Features struct {
	Metrics   *telemetry.MetricsConfig `yaml:"metrics,omitempty"`
	Tracing   *telemetry.TracingConfig `yaml:"tracing,omitempty"`
	TLS       *Features_TLS            `yaml:"tls,omitempty"`
	Validator *Features_Validator      `yaml:"validator,omitempty"`
	Rpclog    *Features_Rpclog         `yaml:"rpclog,omitempty"`
	Translog  *Features_Translog       `yaml:"translog,omitempty"`
}

// Features_TLS struct
type Features_TLS struct {
	Enabled    bool   `yaml:",omitempty" default:"false"`
	CertFile   string `yaml:"cert_file" validate:"file,required"`
	KeyFile    string `yaml:"key_file" validate:"file,required"`
	CaFile     string `yaml:"ca_file" validate:"file,required"`
	Password   string `yaml:"password,omitempty"`
	ServerName string `yaml:"server_name,omitempty" default:"'*'"`
}

// Features_Validator struct
type Features_Validator struct {
	Enabled bool `yaml:",omitempty" default:"false"`
}

// Features_Rpclog struct
type Features_Rpclog struct {
	Enabled bool `yaml:",omitempty" default:"false"`
}

// Features_Translog struct
type Features_Translog struct {
	Enabled bool   `yaml:",omitempty" default:"false"`
	Topic   string `yaml:",omitempty"`
}

// Services struct
type Services struct {
	Account  *Service `yaml:"account,omitempty"`
	Greeter  *Service `yaml:"greeter,omitempty"`
	Emailer  *Service `yaml:"emailer,omitempty"`
	Recorder *Service `yaml:"recorder,omitempty"`
	Play     *Service `yaml:"play,omitempty"`
}

// Configuration struct
type Configuration struct {
	Database *DatabaseConfiguration `yaml:"database,omitempty"`
	Email    *EmailConfiguration    `yaml:"email,omitempty"`
	Features *Features              `yaml:"features,omitempty"`
	Services *Services              `yaml:"services,omitempty"`
}
