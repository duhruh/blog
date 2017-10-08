package main

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/config"

	"github.com/duhruh/blog/app"

	"github.com/go-kit/kit/log"
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

	c := app.NewConfig(tackle.Environment(*environment), cfg)

	logger := app.NewLogger(c)

	defer recoverHandler(logger)

	application := app.NewApplication(context.Background(), c, logger)

	application.Build()

	application.Start()
}

func recoverHandler(l log.Logger) {
	if r := recover(); r != nil {
		l.Log("panic", r)
	}
}
