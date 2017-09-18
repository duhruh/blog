package blog

import (
	"context"
	"github.com/duhruh/blog/app/blog/entity"
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

func (ef endpointFactory) ListBlogsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, err := ef.service.ListBlogs()

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

func (ef endpointFactory) ListPostsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		packet := request.(tackle.Packet)
		bid := packet.Get("blog_id")
		blog := entity.NewBlog()
		blog.SetIdentity(domain.NewIdentity(bid.(string)))

		r, err := ef.service.ListPosts(blog)

		return r, err
	}
}
func (ef endpointFactory) ShowPostEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		packet := request.(tackle.Packet)

		id := domain.NewIdentity(packet.Get("id"))
		r, err := ef.service.ShowPost(id)

		return r, err
	}
}

func (ef endpointFactory) CreatePostEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		packet := request.(tackle.Packet)
		bid := packet.Get("blog_id")
		blog := entity.NewBlog()
		blog.SetIdentity(domain.NewIdentity(bid))

		r, err := ef.service.CreatePost(blog, packet.Get("body").(string))

		return r, err
	}
}

func (ef endpointFactory) UpdateBlogEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		packet := request.(tackle.Packet)
		bid := packet.Get("id")
		blog := entity.NewBlog()
		blog.SetIdentity(domain.NewIdentity(bid))
		blog.SetName(packet.Get("name").(string))

		return ef.service.UpdateBlog(blog)
	}
}
