package blog

import (
	"context"

	"github.com/duhruh/blog/app/blog/transport/http"
	tacklehttp "github.com/duhruh/tackle/transport/http"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	"github.com/duhruh/blog/app/blog/repository"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type App struct {
	service       Service
	httpTransport tacklehttp.HttpTransport
}

func (h App) Service() Service                        { return h.service }
func (h App) HttpTransport() tacklehttp.HttpTransport { return h.httpTransport }

func NewImplementedService(cxt context.Context, logger log.Logger) App {
	fieldKeys := []string{"method"}

	var repo repository.BlogRepository
	repo = repository.NewBlogRepository()

	var service Service
	service = newService(repo)
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

	return App{
		service:       service,
		httpTransport: httpTransport,
	}
}
