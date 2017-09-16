package app

import (
	"github.com/go-kit/kit/log"
	"os"

	cfg "github.com/duhruh/scaffold/config"
)

// This is where we define our application logger
// here we initialize the logger to only output to
// stdout
func NewLogger(_ Config) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(
		logger,
		"timestamp", log.DefaultTimestampUTC,
		"gitCommit", cfg.GitCommit,
		"version", cfg.Version,
		"buildNumber", cfg.BuildNumber,
		"buildTime", cfg.BuildTime,
	)

	logger.Log("message", "application boot")

	return logger
}
