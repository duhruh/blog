package app

import (
	"strconv"

	"github.com/duhruh/blog/config"
	"github.com/duhruh/tackle"
	tackleconfig "github.com/duhruh/tackle/config"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/olivere/elastic.v5"
	"io"
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

func NewConfigFromYamlFile(env tackle.Environment, file string) config.ApplicationConfig {
	var r io.Reader
	r, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	yaml := tackleconfig.NewYamlLoader()
	cfg, err := yaml.LoadFromFile(r)
	if err != nil {
		panic(err)
	}

	return NewConfig(env, cfg)
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

func (c appConfig) Version() string {
	return c.Get("version").(string)
}
func (c appConfig) Description() string {
	return c.Get("description").(string)
}
func (c appConfig) LogFile() string {
	return c.Get("log_path").(string)
}

func (c appConfig) EntryPoint() string {
	return c.Get("entry_point").(string)
}

func (c appConfig) ConfigPath() string {
	return c.Get("config_path").(string)
}

func (c appConfig) TaskEntryPoint() string {
	return c.Get("task_entry_point").(string)
}
func (c appConfig) GenerateElasticSearchClient() *elastic.Client {

	client, err := elastic.NewClient(
		elastic.SetURL(c.ElasticSearch().Get("host").(string)),
		elastic.SetBasicAuth(c.ElasticSearch().Get("username").(string), c.ElasticSearch().Get("password").(string)),
	)

	if err != nil {
		return nil
	}

	return client
}
