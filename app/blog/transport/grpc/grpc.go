package grpc

import (
	"context"
	"github.com/duhruh/blog/app/blog/proto"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type grpcTransport struct {
	logger    log.Logger
	listBlogs kitgrpc.Handler
}

func NewGrpcTransport() {

}
func (gt grpcTransport) NewHandler(g *grpc.Server) {

	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(gt.logger),
	}

	gt.listBlogs = tacklegrpc.NewServer()

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

	proto.RegisterBlogServiceServer(g, gt)

}
func (gt grpcTransport) ListBlogs(ctx context.Context, req *proto.ListBlogsRequest) (*proto.ListBlogsResponse, error) {
	_, rep, err := gt.listBlogs.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.ListBlogsResponse), nil
}
