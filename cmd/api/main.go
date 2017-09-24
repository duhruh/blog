package main

import (
	"context"
	"flag"

	"github.com/duhruh/blog/app"

	"github.com/duhruh/tackle"
)

const (
	defaultHttpPort    = ":8080"
	defaultGrpcPort    = ":8082"
	defaultEnvironment = string(tackle.Production)
)

var (
	environment     = flag.String("environment", defaultEnvironment, "application environment")
	httpBindAddress = flag.String("http-bind-address", defaultHttpPort, "http: Port to bind the server to")
	grpcBindAddress = flag.String("grpc-bind-address", defaultGrpcPort, "grpc: Port to bind the server to")
)

func main() {
	flag.Parse()

	config := app.NewConfig(tackle.Environment(*environment), *httpBindAddress, *grpcBindAddress)

	logger := app.NewLogger(config)

	application := app.NewApplication(context.Background(), config, logger)

	application.Build()

	application.Start()
}
