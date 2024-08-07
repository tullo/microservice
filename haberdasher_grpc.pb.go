// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: haberdasher.proto

package haberdasher

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

// HaberdasherServiceClient is the client API for HaberdasherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HaberdasherServiceClient interface {
	// MakeHat produces a hat of mysterious, randomly-selected color!
	MakeHat(ctx context.Context, in *Size, opts ...grpc.CallOption) (*Hat, error)
}

type haberdasherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHaberdasherServiceClient(cc grpc.ClientConnInterface) HaberdasherServiceClient {
	return &haberdasherServiceClient{cc}
}

func (c *haberdasherServiceClient) MakeHat(ctx context.Context, in *Size, opts ...grpc.CallOption) (*Hat, error) {
	out := new(Hat)
	err := c.cc.Invoke(ctx, "/tullo.microservice.haberdasher.HaberdasherService/MakeHat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HaberdasherServiceServer is the server API for HaberdasherService service.
// All implementations must embed UnimplementedHaberdasherServiceServer
// for forward compatibility
type HaberdasherServiceServer interface {
	// MakeHat produces a hat of mysterious, randomly-selected color!
	MakeHat(context.Context, *Size) (*Hat, error)
	mustEmbedUnimplementedHaberdasherServiceServer()
}

// UnimplementedHaberdasherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHaberdasherServiceServer struct {
}

func (UnimplementedHaberdasherServiceServer) MakeHat(context.Context, *Size) (*Hat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeHat not implemented")
}
func (UnimplementedHaberdasherServiceServer) mustEmbedUnimplementedHaberdasherServiceServer() {}

// UnsafeHaberdasherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HaberdasherServiceServer will
// result in compilation errors.
type UnsafeHaberdasherServiceServer interface {
	mustEmbedUnimplementedHaberdasherServiceServer()
}

func RegisterHaberdasherServiceServer(s grpc.ServiceRegistrar, srv HaberdasherServiceServer) {
	s.RegisterService(&HaberdasherService_ServiceDesc, srv)
}

func _HaberdasherService_MakeHat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Size)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HaberdasherServiceServer).MakeHat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tullo.microservice.haberdasher.HaberdasherService/MakeHat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HaberdasherServiceServer).MakeHat(ctx, req.(*Size))
	}
	return interceptor(ctx, in, info, handler)
}

// HaberdasherService_ServiceDesc is the grpc.ServiceDesc for HaberdasherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HaberdasherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tullo.microservice.haberdasher.HaberdasherService",
	HandlerType: (*HaberdasherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeHat",
			Handler:    _HaberdasherService_MakeHat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "haberdasher.proto",
}
