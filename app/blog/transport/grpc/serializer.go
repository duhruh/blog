package grpc

import (
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/blog/app/blog/proto"
)

type Serializer interface {
	ProtoBlogs(b []entity.Blog) []*proto.Blog
}

type protoSerializer struct {
}

func NewSerializer() Serializer {
	return protoSerializer{}
}

func (ps protoSerializer) ProtoBlogs(blogs []entity.Blog) []*proto.Blog {
	var protoBlogs []*proto.Blog
	for _, b := range blogs {
		protoBlog := &proto.Blog{
			Id:   b.Identity().Identity().(string),
			Name: b.Name(),
		}

		protoBlogs = append(protoBlogs, protoBlog)

	}

	return protoBlogs
}
