package app

import (
	"context"
	"sync"
	"time"

	"github.com/duhruh/blog/app/blog"
	appgrpc "github.com/duhruh/blog/app/transport/grpc"
	apphttp "github.com/duhruh/blog/app/transport/http"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/transport/grpc"
	"github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type application struct {
	context       context.Context
	config        tackle.Config
	logger        log.Logger
	httpTransport http.AppHttpTransport
	grpcTransport grpc.AppGrpcTransport
}

func NewApplication(cxt context.Context, config tackle.Config, logger log.Logger) tackle.Application {
	return &application{context: cxt, config: config, logger: logger}
}

func (a *application) Build() {
	defer func(begin time.Time) {
		level.Info(a.logger).Log("message", "application built", "took", time.Since(begin))
	}(time.Now())

	_ := NewDatabaseConnection(a.config)

	var blogApp blog.App
	blogApp = blog.NewImplementedService(a.context, a.logger)

	var httpTransports []http.HttpTransport
	httpTransports = append(httpTransports, blogApp.HttpTransport())

	var grpcTransports []grpc.GrpcTransport
	grpcTransports = append(grpcTransports, blogApp.GrpcTransport())

	a.httpTransport = apphttp.NewHttpTransport(a.logger, a.config.HttpBindAddress())

	a.grpcTransport = appgrpc.NewGrpcTransport(a.logger, a.config.GrpcBindAddress())

	a.httpTransport.Build(httpTransports)

	a.grpcTransport.Build(grpcTransports)
}

func (a *application) Start() {
	var wg sync.WaitGroup

	a.httpTransport.Start(&wg)
	a.grpcTransport.Start(&wg)

	level.Info(a.logger).Log("message", "application ready")

	wg.Wait()

	level.Info(a.logger).Log("message", "application halting")
}

func (a *application) HttpTransport() http.AppHttpTransport {
	return a.httpTransport
}
func (a *application) GrpcTransport() grpc.AppGrpcTransport {
	return a.grpcTransport
}
