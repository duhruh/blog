package config

import (
	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/config"
	"github.com/go-kit/kit/log/level"
)

var (
	GitCommit string

	Version string

	BuildNumber string

	BuildTime string
)

type ApplicationConfig interface {
	HttpBindAddress() string
	GrpcBindAddress() string
	LogOption() level.Option
	Environment() tackle.Environment
	DatabaseConnection() config.OptionMap
	ElasticHost() string
	Name() string
}
