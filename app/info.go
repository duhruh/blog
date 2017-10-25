package app

import (
	"fmt"
	"os"
	"flag"
	"github.com/duhruh/blog/config"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)


func Info(c config.ApplicationConfig, l log.Logger, help bool, version bool) bool{
	if help {
		usage(c)
		return true
	}

	if version {
		versionInfo(l)
		return true
	}

	return false
}

func usage(c config.ApplicationConfig) {
	use := c.Description()
	fmt.Fprintf(os.Stderr, "Usage of blog [options]:\n\n%s\n\n", use)
	flag.PrintDefaults()
}
func versionInfo(logger log.Logger) {
	level.Debug(logger).Log("message", "version info")
}
