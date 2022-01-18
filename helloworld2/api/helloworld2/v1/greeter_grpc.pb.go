// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Greeter2Client is the client API for Greeter2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Greeter2Client interface {
	// Sends a greeting
	SayHello2(ctx context.Context, in *Hello2Request, opts ...grpc.CallOption) (*Hello2Reply, error)
}

type greeter2Client struct {
	cc grpc.ClientConnInterface
}

func NewGreeter2Client(cc grpc.ClientConnInterface) Greeter2Client {
	return &greeter2Client{cc}
}

func (c *greeter2Client) SayHello2(ctx context.Context, in *Hello2Request, opts ...grpc.CallOption) (*Hello2Reply, error) {
	out := new(Hello2Reply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Greeter2/SayHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Greeter2Server is the server API for Greeter2 service.
// All implementations must embed UnimplementedGreeter2Server
// for forward compatibility
type Greeter2Server interface {
	// Sends a greeting
	SayHello2(context.Context, *Hello2Request) (*Hello2Reply, error)
	mustEmbedUnimplementedGreeter2Server()
}

// UnimplementedGreeter2Server must be embedded to have forward compatible implementations.
type UnimplementedGreeter2Server struct {
}

func (UnimplementedGreeter2Server) SayHello2(context.Context, *Hello2Request) (*Hello2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello2 not implemented")
}
func (UnimplementedGreeter2Server) mustEmbedUnimplementedGreeter2Server() {}

// UnsafeGreeter2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Greeter2Server will
// result in compilation errors.
type UnsafeGreeter2Server interface {
	mustEmbedUnimplementedGreeter2Server()
}

func RegisterGreeter2Server(s grpc.ServiceRegistrar, srv Greeter2Server) {
	s.RegisterService(&Greeter2_ServiceDesc, srv)
}

func _Greeter2_SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Greeter2Server).SayHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Greeter2/SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Greeter2Server).SayHello2(ctx, req.(*Hello2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter2_ServiceDesc is the grpc.ServiceDesc for Greeter2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter2",
	HandlerType: (*Greeter2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello2",
			Handler:    _Greeter2_SayHello2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/helloworld2/v1/greeter.proto",
}