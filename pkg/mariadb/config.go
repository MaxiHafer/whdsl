package mariadb

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("MARIADB", cfg)

	return cfg, errors.Wrap(err, "failure while processing mariadb cfg from env")
}

type Config struct {
	User     string `required:"true"`
	Password string `required:"true"`
	Hostname string `required:"true"`
	Port     int    `default:"3306"`
	Database string `default:"whdsl_data"`
	Retries  uint   `default:"10"`
}

func (c *Config) DSN() string {
	dsn := url.URL{}

	if c.User != "" && c.Password != "" {
		dsn.User = url.UserPassword(c.User, c.Password)
	}

	// mysql driver requires this shitty tcp(host:port)-thing
	// otherwise: BOOM (default addr for network 'localhost:port' unknown)
	// https://github.com/go-sql-driver/mysql/blob/21f789cd2353b7ac81538f41426e9cfd2b1fcc87/dsn.go#L98
	dsn.Host = fmt.Sprintf("tcp(%s:%d)", c.Hostname, c.Port)
	dsn.Path = c.Database

	q := dsn.Query()
	q.Set("multiStatements", "true")
	q.Set("collation", "utf8mb4_general_ci")
	dsn.RawQuery = q.Encode()

	// return dsn without leading "//" added by url package due to missing scheme
	return strings.TrimLeft(dsn.String(), "/")
}
