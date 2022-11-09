package internal

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func NewInitializedMariaDBConfigFromEnv() *MariadbConfig{
	config := &MariadbConfig{}
	envconfig.MustProcess("MARIADB",config)
	return config
}

type MariadbConfig struct {
	User string
	Password string
	Hostname string
	Database string
	Port int
}

func (c *MariadbConfig) DSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.Database,
	)
}