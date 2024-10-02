// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sample/v1/shop.proto

package shop

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

const (
	ShopService_FindShopByName_FullMethodName = "/sample.v1.ShopService/FindShopByName"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	// 店舗名で店舗を検索する
	FindShopByName(ctx context.Context, in *FindShopByNameRequest, opts ...grpc.CallOption) (*FindShopByNameResponse, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) FindShopByName(ctx context.Context, in *FindShopByNameRequest, opts ...grpc.CallOption) (*FindShopByNameResponse, error) {
	out := new(FindShopByNameResponse)
	err := c.cc.Invoke(ctx, ShopService_FindShopByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations should embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	// 店舗名で店舗を検索する
	FindShopByName(context.Context, *FindShopByNameRequest) (*FindShopByNameResponse, error)
}

// UnimplementedShopServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) FindShopByName(context.Context, *FindShopByNameRequest) (*FindShopByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindShopByName not implemented")
}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_FindShopByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindShopByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).FindShopByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_FindShopByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).FindShopByName(ctx, req.(*FindShopByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.v1.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindShopByName",
			Handler:    _ShopService_FindShopByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample/v1/shop.proto",
}
