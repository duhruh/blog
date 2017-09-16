package app

import (
	"context"
	"github.com/duhruh/blog/app/blog"
	http3 "github.com/duhruh/scaffold/app/transport/http"
	"github.com/duhruh/tackle"
	http2 "github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
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

	var transports []http2.HttpTransport
	transports = append(transports, blogApp.HttpTransport())

	httpTransport := http3.NewHttpTransport(a.logger, a.config.GetHTTPBindAddress())

	httpTransport.Mount(transports)
}
