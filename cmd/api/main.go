package main

import (
	"context"
	"flag"

	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/log"

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

	c := app.NewConfigFromYamlFile(tackle.Environment(*environment), *appConfig)

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
