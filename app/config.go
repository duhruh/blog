package app

import (
	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/log/level"
)

type config struct {
	httpBindAddress    string
	grpcBindAddress    string
	environment        tackle.Environment
	databaseConnection map[string]string
}

func NewConfig(env tackle.Environment, httpAddr string, grpcAddr string, connection map[string]string) tackle.Config {
	return config{
		environment:        env,
		httpBindAddress:    httpAddr,
		grpcBindAddress:    grpcAddr,
		databaseConnection: connection,
	}
}

func (c config) HttpBindAddress() string {
	return c.httpBindAddress
}
func (c config) GrpcBindAddress() string {
	return c.grpcBindAddress
}

func (c config) Environment() tackle.Environment {
	return c.environment
}
func (c config) LogOption() level.Option {
	switch c.environment {
	case tackle.Development:
		return level.AllowAll()
	case tackle.Production:
		return level.AllowInfo()
	default:
		return level.AllowAll()
	}
}

func (c config) DatabaseConnection() map[string]string {
	return c.databaseConnection
}
