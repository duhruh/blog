package grpc

import (
	"context"
	"github.com/duhruh/tackle"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type encoderFactory struct {
	tacklegrpc.EncoderFactory
}

func NewEncoderFactory() tacklegrpc.EncoderFactory {
	return encoderFactory{
		EncoderFactory: tacklegrpc.NewEncoderFactory(),
	}
}

func (ef encoderFactory) Generate(e string) (tacklegrpc.Encoder, error) {
	return ef.GenerateWithInstance(ef, e)
}

func (hs encoderFactory) ListBlogEncoder() tacklegrpc.Encoder {
	return tacklegrpc.NewEncoder(hs.listBlogsRequest(), hs.listBlogsResponse())
}

func (hs encoderFactory) listBlogsRequest() kitgrpc.DecodeRequestFunc {
	return kitgrpc.DecodeRequestFunc(func(_ context.Context, grpcReq interface{}) (interface{}, error) {
		return tackle.NewPacket(), nil
	})
}

func (hs encoderFactory) listBlogsResponse() kitgrpc.EncodeResponseFunc {
	return kitgrpc.EncodeResponseFunc(func(_ context.Context, grpcReply interface{}) (interface{}, error) {
		return tackle.NewPacket(), nil
	})
}
