package app

import (
	"context"

	"github.com/duhruh/blog/app/blog"
	apphttp "github.com/duhruh/blog/app/transport/http"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/transport/http"
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

	var transports []http.HttpTransport
	transports = append(transports, blogApp.HttpTransport())

	httpTransport := apphttp.NewHttpTransport(a.logger, a.config.GetHTTPBindAddress())

	httpTransport.Mount(transports)
}
