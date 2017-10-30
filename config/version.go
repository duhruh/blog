package config

import (
	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/config"
	"github.com/go-kit/kit/log/level"

	"gopkg.in/olivere/elastic.v5"
)

var (
	GitCommit string

	BuildNumber string

	BuildTime string
)

type ApplicationConfig interface {
	HttpBindAddress() string
	GrpcBindAddress() string
	LogOption() level.Option
	Environment() tackle.Environment
	DatabaseConnection() config.OptionMap
	ElasticSearch() config.OptionMap
	Name() string
	Host() string
	GenerateElasticSearchClient() *elastic.Client
	Version() string
	Description() string
	LogFile() string
}
