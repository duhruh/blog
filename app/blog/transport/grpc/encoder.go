package grpc

import (
	"context"

	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/blog/app/blog/proto"
	"github.com/duhruh/tackle"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type encoderFactory struct {
	tacklegrpc.EncoderFactory
	serializer Serializer
}

func NewEncoderFactory(s Serializer) tacklegrpc.EncoderFactory {
	return encoderFactory{
		EncoderFactory: tacklegrpc.NewEncoderFactory(),
		serializer:     s,
	}
}

func (ef encoderFactory) Generate(e string) (tacklegrpc.Encoder, error) {
	return ef.GenerateWithInstance(ef, e)
}

func (hs encoderFactory) ListBlogsEncoder() tacklegrpc.Encoder {
	return tacklegrpc.NewEncoder(hs.listBlogsRequest(), hs.listBlogsResponse())
}

func (hs encoderFactory) listBlogsRequest() kitgrpc.DecodeRequestFunc {
	return kitgrpc.DecodeRequestFunc(func(_ context.Context, grpcReq interface{}) (interface{}, error) {
		return tackle.NewPacket(), nil
	})
}

func (hs encoderFactory) listBlogsResponse() kitgrpc.EncodeResponseFunc {
	return kitgrpc.EncodeResponseFunc(func(_ context.Context, grpcReply interface{}) (interface{}, error) {
		pkt := grpcReply.(tackle.Packet)
		er := pkt.Get("error")
		if er != nil {
			return nil, er.(error)
		}

		data := pkt.Get("data")

		blogs := data.([]entity.Blog)

		protoBlogs := hs.serializer.ProtoBlogs(blogs)

		return &proto.ListBlogsResponse{Blogs: protoBlogs}, nil
	})
}
