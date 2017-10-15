package app

import (
	applog "github.com/duhruh/blog/app/log"
	"github.com/duhruh/blog/config"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// This is where we define our application logger
// here we initialize the logger to only output to
// stdout
func NewLogger(c config.ApplicationConfig) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(applog.NewColorWriter(log.NewSyncWriter(os.Stderr)))
	logger, err := applog.NewElasticSearchLogger(c.GenerateElasticSearchClient(), c.Host(), c.Name(), logger)
	if err != nil {
		level.Error(logger).Log("error", err)
	}
	logger = level.NewFilter(logger, c.LogOption())
	logger = log.With(
		logger,
		"timestamp", log.DefaultTimestampUTC,
		"environment", c.Environment(),
		"gitCommit", config.GitCommit,
		"version", config.Version,
		"buildNumber", config.BuildNumber,
		"buildTime", config.BuildTime,
		"caller", log.DefaultCaller,
	)

	level.Info(logger).Log("message", "application booting")

	return logger
}
