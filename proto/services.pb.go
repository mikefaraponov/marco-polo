// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	Send(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Send(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/proto.Server/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	Send(context.Context, *Phrase) (*Phrase, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phrase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Send(ctx, req.(*Phrase))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Server_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// AdapterClient is the client API for Adapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdapterClient interface {
	Forward(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error)
}

type adapterClient struct {
	cc *grpc.ClientConn
}

func NewAdapterClient(cc *grpc.ClientConn) AdapterClient {
	return &adapterClient{cc}
}

func (c *adapterClient) Forward(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/proto.Adapter/Forward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdapterServer is the server API for Adapter service.
type AdapterServer interface {
	Forward(context.Context, *Phrase) (*Phrase, error)
}

func RegisterAdapterServer(s *grpc.Server, srv AdapterServer) {
	s.RegisterService(&_Adapter_serviceDesc, srv)
}

func _Adapter_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phrase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Adapter/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServer).Forward(ctx, req.(*Phrase))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Adapter",
	HandlerType: (*AdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Forward",
			Handler:    _Adapter_Forward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_services_554fa65f7b16ae42) }

var fileDescriptor_services_554fa65f7b16ae42 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x7c,
	0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x61, 0x23, 0x03, 0x2e, 0xb6, 0xe0, 0xd4, 0xa2, 0xb2,
	0xd4, 0x22, 0x21, 0x35, 0x2e, 0x96, 0xe0, 0xd4, 0xbc, 0x14, 0x21, 0x5e, 0x88, 0x8c, 0x5e, 0x40,
	0x46, 0x51, 0x62, 0x71, 0xaa, 0x14, 0x2a, 0x57, 0x89, 0xc1, 0xc8, 0x84, 0x8b, 0xdd, 0x31, 0x25,
	0xb1, 0xa0, 0x24, 0xb5, 0x48, 0x48, 0x93, 0x8b, 0xdd, 0x2d, 0xbf, 0xa8, 0x3c, 0xb1, 0x88, 0xa0,
	0xae, 0x24, 0x36, 0x30, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x5a, 0x6b, 0x67, 0x97,
	0x00, 0x00, 0x00,
}
