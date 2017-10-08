package app

import (
	"strconv"

	"github.com/duhruh/blog/config"
	"github.com/duhruh/tackle"
	tackleconfig "github.com/duhruh/tackle/config"
	"github.com/go-kit/kit/log/level"
)

type appConfig struct {
	tackleconfig.Config
	environment tackle.Environment
}

func NewConfig(env tackle.Environment, raw tackleconfig.Config) config.ApplicationConfig {
	return appConfig{
		environment: env,
		Config:      raw,
	}
}

func (c appConfig) HttpBindAddress() string {
	port := c.Get("http").(tackleconfig.OptionMap).Get("port").(int)
	return ":" + strconv.Itoa(port)
}
func (c appConfig) GrpcBindAddress() string {
	port := c.Get("grpc").(tackleconfig.OptionMap).Get("port").(int)
	return ":" + strconv.Itoa(port)
}

func (c appConfig) Environment() tackle.Environment {
	return c.environment
}
func (c appConfig) LogOption() level.Option {
	switch c.environment {
	case tackle.Development:
		return level.AllowAll()
	case tackle.Production:
		return level.AllowInfo()
	default:
		return level.AllowAll()
	}
}

func (c appConfig) DatabaseConnection() tackleconfig.OptionMap {
	return c.Get("database").(tackleconfig.OptionMap).Get(string(c.environment)).(tackleconfig.OptionMap)
}
