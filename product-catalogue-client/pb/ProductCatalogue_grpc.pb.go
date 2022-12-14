// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: ProductCatalogue.proto

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

// ProductCatalogClient is the client API for ProductCatalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogClient interface {
	// Get product details
	GetProductDetails(ctx context.Context, in *ProductDetailsRequest, opts ...grpc.CallOption) (*ProductDetailsReply, error)
}

type productCatalogClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogClient(cc grpc.ClientConnInterface) ProductCatalogClient {
	return &productCatalogClient{cc}
}

func (c *productCatalogClient) GetProductDetails(ctx context.Context, in *ProductDetailsRequest, opts ...grpc.CallOption) (*ProductDetailsReply, error) {
	out := new(ProductDetailsReply)
	err := c.cc.Invoke(ctx, "/product.ProductCatalog/GetProductDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServer is the server API for ProductCatalog service.
// All implementations must embed UnimplementedProductCatalogServer
// for forward compatibility
type ProductCatalogServer interface {
	// Get product details
	GetProductDetails(context.Context, *ProductDetailsRequest) (*ProductDetailsReply, error)
	mustEmbedUnimplementedProductCatalogServer()
}

// UnimplementedProductCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServer struct {
}

func (UnimplementedProductCatalogServer) GetProductDetails(context.Context, *ProductDetailsRequest) (*ProductDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetails not implemented")
}
func (UnimplementedProductCatalogServer) mustEmbedUnimplementedProductCatalogServer() {}

// UnsafeProductCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServer will
// result in compilation errors.
type UnsafeProductCatalogServer interface {
	mustEmbedUnimplementedProductCatalogServer()
}

func RegisterProductCatalogServer(s grpc.ServiceRegistrar, srv ProductCatalogServer) {
	s.RegisterService(&ProductCatalog_ServiceDesc, srv)
}

func _ProductCatalog_GetProductDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServer).GetProductDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductCatalog/GetProductDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServer).GetProductDetails(ctx, req.(*ProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalog_ServiceDesc is the grpc.ServiceDesc for ProductCatalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductCatalog",
	HandlerType: (*ProductCatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductDetails",
			Handler:    _ProductCatalog_GetProductDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ProductCatalogue.proto",
}
