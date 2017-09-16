package http

import (
	"net/http"

	"github.com/duhruh/tackle"
	tacklehttp "github.com/duhruh/tackle/transport/http"

	"github.com/go-kit/kit/log"
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
	router := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(h.encoderFactory.ErrorEncoder()),
		kithttp.ServerErrorLogger(h.logger),
	}

	allBlogsEndpoint, _ := h.endpointFactory.Generate("AllBlogsEndpoint")
	allBlogsEncoder, _ := h.encoderFactory.Generate("AllBlogsEncoder")
	router.Handle("/blogs", tacklehttp.NewServer(
		allBlogsEndpoint,
		allBlogsEncoder,
		options,
	)).Methods("GET")

	createBlogEndpoint, _ := h.endpointFactory.Generate("CreateBlogEndpoint")
	createBlogEncoder, _ := h.encoderFactory.Generate("CreateBlogEncoder")
	router.Handle("/blogs", tacklehttp.NewServer(
		createBlogEndpoint,
		createBlogEncoder,
		options,
	)).Methods("POST")

	showBlogEndpoint, _ := h.endpointFactory.Generate("ShowBlogEndpoint")
	showBlogEncoder, _ := h.encoderFactory.Generate("ShowBlogEncoder")
	router.Handle("/blogs/{id:[0-9]+}", tacklehttp.NewServer(
		showBlogEndpoint,
		showBlogEncoder,
		options,
	)).Methods("GET")

	handler.Handle("/", router)
	return handler
}
