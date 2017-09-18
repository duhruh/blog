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

var uuidRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var routes = []struct {
	Method   string
	Path     string
	Endpoint string
	Encoder  string
}{
	{
		"GET",
		"/blogs",
		"ListBlogsEndpoint",
		"ListBlogsEncoder",
	},
	{
		"POST",
		"/blogs",
		"CreateBlogEndpoint",
		"CreateBlogEncoder",
	},
	{
		"PUT",
		"/blogs/{id:" + uuidRegex + "}",
		"UpdateBlogEndpoint",
		"UpdateBlogEncoder",
	},
	{
		"GET",
		"/blogs/{id:" + uuidRegex + "}",
		"ShowBlogEndpoint",
		"ShowBlogEncoder",
	},

	{
		"GET",
		"/blogs/{id:" + uuidRegex + "}/posts",
		"ListPostsEndpoint",
		"ListPostsEncoder",
	},
	{
		"POST",
		"/blogs/{id:" + uuidRegex + "}/posts",
		"CreatePostEndpoint",
		"CreatePostEncoder",
	},
	{
		"GET",
		"/posts/{id:" + uuidRegex + "}",
		"ShowPostEndpoint",
		"ShowPostEncoder",
	},
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

	for _, route := range routes {
		endpoint, err := h.endpointFactory.Generate(route.Endpoint)
		if err != nil {
			panic(err)
		}
		encoder, err := h.encoderFactory.Generate(route.Encoder)
		if err != nil {
			panic(err)
		}

		router.Handle(route.Path, tacklehttp.NewServer(
			endpoint,
			encoder,
			options,
		)).Methods(route.Method)
	}
	handler.Handle("/", router)
	return handler
}
