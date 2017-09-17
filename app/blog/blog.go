package blog

import (
	"context"

	"github.com/duhruh/blog/app/blog/transport/http"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	tacklehttp "github.com/duhruh/tackle/transport/http"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	"github.com/duhruh/blog/app/blog/repository"
	"github.com/duhruh/blog/app/blog/transport/grpc"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type App struct {
	service       Service
	httpTransport tacklehttp.HttpTransport
	grpcTransport tacklegrpc.GrpcTransport
}

func (h App) Service() Service                        { return h.service }
func (h App) HttpTransport() tacklehttp.HttpTransport { return h.httpTransport }
func (h App) GrpcTransport() tacklegrpc.GrpcTransport { return h.grpcTransport }

func NewImplementedService(cxt context.Context, logger log.Logger) App {
	fieldKeys := []string{"method"}

	var blogRepo repository.BlogRepository
	blogRepo = repository.NewBlogRepository()

	var postRepo repository.PostRepository
	postRepo = repository.NewPostRepository()

	var service Service
	service = newService(blogRepo, postRepo)
	service = newLoggingService(log.With(logger, "component", "blog"), service)
	service = newInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "blog_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "blog_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		service,
	)

	endpointFactory := newEndpointFactory(service)

	httpTransport := http.NewHttpTransport(endpointFactory, log.With(logger, "component", "http"))
	grpcTransport := grpc.NewGrpcTransport(endpointFactory, log.With(logger, "component", "grpc"))

	return App{
		service:       service,
		httpTransport: httpTransport,
		grpcTransport: grpcTransport,
	}
}
