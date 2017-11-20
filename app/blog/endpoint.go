package blog

import (
	"context"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/endpoint"

	"github.com/duhruh/blog/app/blog/entity"
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

var uhh = 0

func (ef endpointFactory) ListBlogsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uhh += 1
		c := context.WithValue(ctx, "count", uhh)
		s := ef.service.WithContext(c)
		r, err := s.ListBlogs()

		pkt := tackle.NewPacket()
		pkt.Put("data", r)
		pkt.Put("error", err)
		return pkt, nil
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
