// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: v1/pb/ipam.proto

package pb

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

// IpamClient is the client API for Ipam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IpamClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type ipamClient struct {
	cc grpc.ClientConnInterface
}

func NewIpamClient(cc grpc.ClientConnInterface) IpamClient {
	return &ipamClient{cc}
}

func (c *ipamClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/ipam.ipam/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipamClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/ipam.ipam/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IpamServer is the server API for Ipam service.
// All implementations must embed UnimplementedIpamServer
// for forward compatibility
type IpamServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Del(context.Context, *DelRequest) (*DelResponse, error)
	mustEmbedUnimplementedIpamServer()
}

// UnimplementedIpamServer must be embedded to have forward compatible implementations.
type UnimplementedIpamServer struct {
}

func (UnimplementedIpamServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedIpamServer) Del(context.Context, *DelRequest) (*DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedIpamServer) mustEmbedUnimplementedIpamServer() {}

// UnsafeIpamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IpamServer will
// result in compilation errors.
type UnsafeIpamServer interface {
	mustEmbedUnimplementedIpamServer()
}

func RegisterIpamServer(s grpc.ServiceRegistrar, srv IpamServer) {
	s.RegisterService(&Ipam_ServiceDesc, srv)
}

func _Ipam_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipam.ipam/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipam_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipam.ipam/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ipam_ServiceDesc is the grpc.ServiceDesc for Ipam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ipam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ipam.ipam",
	HandlerType: (*IpamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Ipam_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Ipam_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/pb/ipam.proto",
}
