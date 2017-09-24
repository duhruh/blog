package app

import (
	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/log/level"
)

type Config interface {
	HttpBindAddress() string
	GrpcBindAddress() string
	LogOption() level.Option
	Environment() tackle.Environment
}
type config struct {
	httpBindAddress string
	grpcBindAddress string
	environment     tackle.Environment
}

func NewConfig(env tackle.Environment, httpAddr string, grpcAddr string) Config {
	return config{environment: env, httpBindAddress: httpAddr, grpcBindAddress: grpcAddr}
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
