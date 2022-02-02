// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/getuserbyname.proto

package proto

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

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddServiceClient interface {
	GetUserByName(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*User, error)
}

type addServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceClient(cc grpc.ClientConnInterface) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) GetUserByName(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.AddService/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
// All implementations must embed UnimplementedAddServiceServer
// for forward compatibility
type AddServiceServer interface {
	GetUserByName(context.Context, *SearchQuery) (*User, error)
	mustEmbedUnimplementedAddServiceServer()
}

// UnimplementedAddServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (UnimplementedAddServiceServer) GetUserByName(context.Context, *SearchQuery) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedAddServiceServer) mustEmbedUnimplementedAddServiceServer() {}

// UnsafeAddServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddServiceServer will
// result in compilation errors.
type UnsafeAddServiceServer interface {
	mustEmbedUnimplementedAddServiceServer()
}

func RegisterAddServiceServer(s grpc.ServiceRegistrar, srv AddServiceServer) {
	s.RegisterService(&AddService_ServiceDesc, srv)
}

func _AddService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetUserByName(ctx, req.(*SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// AddService_ServiceDesc is the grpc.ServiceDesc for AddService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _AddService_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/getuserbyname.proto",
}
