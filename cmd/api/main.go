package main

import (
	"context"
	"flag"

	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/log"

	"fmt"
	"github.com/duhruh/blog/app"
	"github.com/duhruh/blog/config"
	"github.com/go-kit/kit/log/level"
	"os"
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

	if *help {
		usage(c)
		return
	}

	if *version {
		versionInfo(logger)
		return
	}

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

func usage(c config.ApplicationConfig) {
	use := c.Description()
	fmt.Fprintf(os.Stderr, "Usage of blog [options]:\n\n%s\n\n", use)
	flag.PrintDefaults()
}
func versionInfo(logger log.Logger) {
	level.Debug(logger).Log("message", "version info")
}
