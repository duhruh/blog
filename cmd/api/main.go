package main

import (
	"context"
	"flag"
	"github.com/duhruh/blog/app"
)

const (
	defaultPort = ":8080"
)

var (
	// HTTP configuration
	httpBindAddress = flag.String("http-bind-address", defaultPort, "http: Port to bind the server to")
)

func main() {
	flag.Parse()

	config := app.NewConfig(*httpBindAddress)

	logger := app.NewLogger(config)

	application := app.NewApplication(context.Background(), config, logger)

	application.Start()
}
