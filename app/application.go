package app

import (
	"context"

	"github.com/duhruh/blog/app/blog"
	appgrpc "github.com/duhruh/blog/app/transport/grpc"
	apphttp "github.com/duhruh/blog/app/transport/http"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/transport/grpc"
	"github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
	"sync"
)

type application struct {
	context context.Context
	config  Config
	logger  log.Logger
}

func NewApplication(cxt context.Context, config Config, logger log.Logger) tackle.Application {
	return application{context: cxt, config: config, logger: logger}
}

func (a application) Start() {
	var blogApp blog.App
	blogApp = blog.NewImplementedService(a.context, a.logger)

	var httpTransports []http.HttpTransport
	httpTransports = append(httpTransports, blogApp.HttpTransport())

	var grpcTransports []grpc.GrpcTransport
	grpcTransports = append(grpcTransports, blogApp.GrpcTransport())

	httpTransport := apphttp.NewHttpTransport(a.logger, a.config.GetHTTPBindAddress())

	grpcTransport := appgrpc.NewGrpcTransport(a.logger, a.config.GetGRPCBindAddress())

	var wg sync.WaitGroup

	httpTransport.Mount(httpTransports, &wg)

	grpcTransport.Mount(grpcTransports, &wg)

	wg.Wait()
	a.logger.Log("message", "application halting")
}
