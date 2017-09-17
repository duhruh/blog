// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/blog/proto/blog.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	app/blog/proto/blog.proto

It has these top-level messages:
	ListBlogsRequest
	ListBlogsResponse
	Blog
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ListBlogsRequest struct {
}

func (m *ListBlogsRequest) Reset()                    { *m = ListBlogsRequest{} }
func (m *ListBlogsRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListBlogsRequest) ProtoMessage()               {}
func (*ListBlogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListBlogsResponse struct {
	Blogs []*Blog `protobuf:"bytes,1,rep,name=blogs" json:"blogs,omitempty"`
}

func (m *ListBlogsResponse) Reset()                    { *m = ListBlogsResponse{} }
func (m *ListBlogsResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListBlogsResponse) ProtoMessage()               {}
func (*ListBlogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListBlogsResponse) GetBlogs() []*Blog {
	if m != nil {
		return m.Blogs
	}
	return nil
}

type Blog struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Blog) Reset()                    { *m = Blog{} }
func (m *Blog) String() string            { return proto1.CompactTextString(m) }
func (*Blog) ProtoMessage()               {}
func (*Blog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto1.RegisterType((*ListBlogsRequest)(nil), "proto.ListBlogsRequest")
	proto1.RegisterType((*ListBlogsResponse)(nil), "proto.ListBlogsResponse")
	proto1.RegisterType((*Blog)(nil), "proto.Blog")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlogService service

type BlogServiceClient interface {
	ListBlogs(ctx context.Context, in *ListBlogsRequest, opts ...grpc.CallOption) (*ListBlogsResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) ListBlogs(ctx context.Context, in *ListBlogsRequest, opts ...grpc.CallOption) (*ListBlogsResponse, error) {
	out := new(ListBlogsResponse)
	err := grpc.Invoke(ctx, "/proto.BlogService/ListBlogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlogService service

type BlogServiceServer interface {
	ListBlogs(context.Context, *ListBlogsRequest) (*ListBlogsResponse, error)
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_ListBlogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ListBlogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BlogService/ListBlogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ListBlogs(ctx, req.(*ListBlogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBlogs",
			Handler:    _BlogService_ListBlogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/blog/proto/blog.proto",
}

func init() { proto1.RegisterFile("app/blog/proto/blog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2c, 0x28, 0xd0,
	0x4f, 0xca, 0xc9, 0x4f, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x33, 0xf5, 0xc0, 0x4c, 0x21,
	0x56, 0x30, 0xa5, 0x24, 0xc4, 0x25, 0xe0, 0x93, 0x59, 0x5c, 0xe2, 0x94, 0x93, 0x9f, 0x5e, 0x1c,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x64, 0xc6, 0x25, 0x88, 0x24, 0x56, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x2a, 0xa4, 0xc8, 0xc5, 0x0a, 0xd2, 0x5d, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d,
	0xc4, 0x0d, 0x31, 0x46, 0x0f, 0xa4, 0x28, 0x08, 0x22, 0xa3, 0xa4, 0xc5, 0xc5, 0x02, 0xe2, 0x0a,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x23, 0x7f, 0x2e,
	0x6e, 0x90, 0xda, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x07, 0x2e, 0x4e, 0xb8, 0x95,
	0x42, 0xe2, 0x50, 0xb3, 0xd1, 0x1d, 0x26, 0x25, 0x81, 0x29, 0x01, 0x71, 0x9d, 0x12, 0x43, 0x12,
	0x1b, 0x58, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x94, 0xb9, 0x84, 0xc2, 0xf3, 0x00, 0x00,
	0x00,
}