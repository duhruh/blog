package grpc

import (
	"github.com/duhruh/blog/app/blog/proto"

	"github.com/duhruh/tackle"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type grpcTransport struct {
	logger          log.Logger
	endpointFactory tackle.EndpointFactory
	encoderFactory  tacklegrpc.EncoderFactory
	handlers        map[string]kitgrpc.Handler
}

func NewGrpcTransport(endpointFactory tackle.EndpointFactory, logger log.Logger) tacklegrpc.GrpcTransport {
	return grpcTransport{
		logger:          logger,
		endpointFactory: endpointFactory,
		encoderFactory:  NewEncoderFactory(),
		handlers:        make(map[string]kitgrpc.Handler),
	}

}
func (gt grpcTransport) NewHandler(g *grpc.Server) {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(gt.logger),
	}

	for _, handler := range gt.Handlers() {
		ef, _ := gt.endpointFactory.Generate(handler.Endpoint())
		ec, _ := gt.encoderFactory.Generate(handler.Encoder())

		gt.handlers[handler.Name()] = tacklegrpc.NewServer(ef, ec, options)
	}

	proto.RegisterBlogServiceServer(g, gt)
}

func (gt grpcTransport) Handlers() []tacklegrpc.Handler {
	return getHandlers()
}

func (gt grpcTransport) ListBlogs(ctx context.Context, req *proto.ListBlogsRequest) (*proto.ListBlogsResponse, error) {
	_, rep, err := gt.handlers["BlogService.listBlogs"].ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.ListBlogsResponse), nil
}
