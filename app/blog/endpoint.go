package blog

import (
	"context"
	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/endpoint"
)

type endpointFactory struct {
	tackle.EndpointFactory
	service Service
}

func newEndpointFactory(s Service) tackle.EndpointFactory {
	return endpointFactory{
		EndpointFactory: tackle.NewEndpointFactory(),
		service:         s,
	}
}

func (ef endpointFactory) Generate(end string) (endpoint.Endpoint, error) {
	return ef.EndpointFactory.GenerateWithInstance(ef, end)
}

func (ef endpointFactory) AllBlogsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, err := ef.service.AllBlogs()

		return r, err
	}
}
func (ef endpointFactory) ShowBlogEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		packet := request.(tackle.Packet)

		id := domain.NewIdentity(packet.Get("id"))
		r, err := ef.service.ShowBlog(id)

		return r, err
	}
}

func (ef endpointFactory) CreateBlogEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		packet := request.(tackle.Packet)

		r, err := ef.service.CreateBlog(packet.Get("name").(string))

		return r, err
	}
}
