package app

import (
	"strconv"

	"github.com/duhruh/blog/config"
	"github.com/duhruh/tackle"
	tackleconfig "github.com/duhruh/tackle/config"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/olivere/elastic.v5"
	"os"
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

func (c appConfig) Name() string {
	return c.Get("name").(string)
}

func (c appConfig) ElasticSearch() tackleconfig.OptionMap {
	return c.Get("elasticsearch").(tackleconfig.OptionMap)
}

func (c appConfig) DatabaseConnection() tackleconfig.OptionMap {
	return c.Get("database").(tackleconfig.OptionMap).Get(string(c.environment)).(tackleconfig.OptionMap)
}

func (c appConfig) Host() string {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return host
}

func (c appConfig) GenerateElasticSearchClient() *elastic.Client {

	client, err := elastic.NewClient(
		elastic.SetURL(c.ElasticSearch().Get("host").(string)),
		elastic.SetBasicAuth(c.ElasticSearch().Get("username").(string), c.ElasticSearch().Get("password").(string)),
	)

	if err != nil {
		panic(err)
	}

	return client
}
