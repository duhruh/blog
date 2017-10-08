package main

import (
	"context"
	"flag"
	"io"
	"os"

	//"github.com/duhruh/blog/app"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/config"
	//"io/ioutil"
	"github.com/duhruh/blog/app"
)

const (
	defaultEnvironment = string(tackle.Development)
	defaultAppConfig   = "config/app.yml"
)

var (
	environment = flag.String("environment", defaultEnvironment, "application environment")
	appConfig   = flag.String("config", defaultAppConfig, "application config file")
)

func main() {
	flag.Parse()

	var r io.Reader
	r, err := os.Open(*appConfig)
	if err != nil {
		panic(err)
	}

	yaml := config.NewYamlLoader()
	cfg, err := yaml.LoadFromFile(r)
	if err != nil {
		panic(err)
	}

	//name := cfg.Get("name").(string)
	//httpPort := cfg.Get("http").(config.OptionMap).Get("port").(int)
	//databaseHost := cfg.Get("database").(config.OptionMap).Get("development").(config.OptionMap).Get("host").(string)
	//
	//println(name)
	//println(httpPort)
	//println(databaseHost)

	//dbFile, err := ioutil.ReadFile(*databaseConfig)
	//
	config := app.NewConfig(tackle.Environment(*environment), cfg)

	logger := app.NewLogger(config)

	application := app.NewApplication(context.Background(), config, logger)

	application.Build()

	application.Start()
}
