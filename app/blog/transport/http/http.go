package http

import (
	"net/http"

	"github.com/duhruh/tackle"
	tacklehttp "github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type httpTransport struct {
	encoderFactory  tacklehttp.EncoderFactory
	endpointFactory tackle.EndpointFactory
	logger          log.Logger
}

func NewHttpTransport(endpointFactory tackle.EndpointFactory, logger log.Logger) tacklehttp.HttpTransport {
	return httpTransport{
		encoderFactory:  NewEncoderFactory(),
		endpointFactory: endpointFactory,
		logger:          logger,
	}
}

func (h httpTransport) NewHandler(handler *http.ServeMux) http.Handler {
	routes := h.Routes()
	router := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(h.encoderFactory.ErrorEncoder()),
		kithttp.ServerErrorLogger(level.Error(h.logger)),
	}

	for _, route := range routes {
		endpoint, err := h.endpointFactory.Generate(route.Endpoint())
		if err != nil {
			panic(err)
		}
		encoder, err := h.encoderFactory.Generate(route.Encoder())
		if err != nil {
			panic(err)
		}

		router.Handle(route.Path(), tacklehttp.NewServer(
			endpoint,
			encoder,
			options,
		)).Methods(route.Method())
	}
	handler.Handle("/", router)
	return handler
}

func (h httpTransport) Routes() []tacklehttp.Route {
	return getRoutes()
}
