package main

import (
	"context"
	"flag"

	"github.com/duhruh/tackle"
	"github.com/duhruh/blog/app"
)

const (
	defaultEnvironment = string(tackle.Development)
	defaultAppConfig   = "config/app.yml"
)

var (
	environment = flag.String("environment", defaultEnvironment, "application environment")
	appConfig   = flag.String("config", defaultAppConfig, "application config file")
	help        = flag.Bool("help", false, "prints the help information")
	version     = flag.Bool("version", false, "prints the version information")
)

func main() {
	flag.Parse()

	c := app.NewConfigFromYamlFile(tackle.Environment(*environment), *appConfig)

	logger := app.NewLogger(c)

	if app.Info(c, logger, *help, *version){
		return
	}

	defer app.RecoverHandler(logger)

	application := app.NewApplication(context.Background(), c, logger)

	application.Build()

	application.Start()
}


