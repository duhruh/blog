package app

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	applog "github.com/duhruh/blog/app/log"
	"github.com/duhruh/blog/config"
)

// This is where we define our application logger
// here we initialize the logger to only output to
// stdout
func NewLogger(c config.ApplicationConfig) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(applog.NewColorWriter(log.NewSyncWriter(os.Stderr)))
	logger = level.NewFilter(logger, c.LogOption())
	logger = applog.NewFileLogger(c.LogFile(), logger, c)
	logger = log.With(
		logger,
		"timestamp", log.DefaultTimestampUTC,
		"environment", c.Environment(),
		"gitCommit", config.GitCommit,
		"version", c.Version(),
		"buildNumber", config.BuildNumber,
		"buildTime", config.BuildTime,
		"caller", log.DefaultCaller,
	)

	return logger
}
