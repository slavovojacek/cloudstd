// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: acme/shelf/v1/service.proto

package shelfv1

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

// ShelfServiceClient is the client API for ShelfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShelfServiceClient interface {
	// ListShelves retrieves a list of shelves resources from the server.
	ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// GetShelf retrieves a single shelf resource from the server.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*GetShelfResponse, error)
	// CreateShelf creates a new shelf resource on the server.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error)
	// UpdateShelf updates the shelf resource on the server.
	UpdateShelf(ctx context.Context, in *UpdateShelfRequest, opts ...grpc.CallOption) (*UpdateShelfResponse, error)
	// DeleteShelf deletes the shelf resource from the server.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*DeleteShelfResponse, error)
}

type shelfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShelfServiceClient(cc grpc.ClientConnInterface) ShelfServiceClient {
	return &shelfServiceClient{cc}
}

func (c *shelfServiceClient) ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, "/acme.shelf.v1.ShelfService/ListShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*GetShelfResponse, error) {
	out := new(GetShelfResponse)
	err := c.cc.Invoke(ctx, "/acme.shelf.v1.ShelfService/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error) {
	out := new(CreateShelfResponse)
	err := c.cc.Invoke(ctx, "/acme.shelf.v1.ShelfService/CreateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) UpdateShelf(ctx context.Context, in *UpdateShelfRequest, opts ...grpc.CallOption) (*UpdateShelfResponse, error) {
	out := new(UpdateShelfResponse)
	err := c.cc.Invoke(ctx, "/acme.shelf.v1.ShelfService/UpdateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*DeleteShelfResponse, error) {
	out := new(DeleteShelfResponse)
	err := c.cc.Invoke(ctx, "/acme.shelf.v1.ShelfService/DeleteShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShelfServiceServer is the server API for ShelfService service.
// All implementations should embed UnimplementedShelfServiceServer
// for forward compatibility
type ShelfServiceServer interface {
	// ListShelves retrieves a list of shelves resources from the server.
	ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error)
	// GetShelf retrieves a single shelf resource from the server.
	GetShelf(context.Context, *GetShelfRequest) (*GetShelfResponse, error)
	// CreateShelf creates a new shelf resource on the server.
	CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error)
	// UpdateShelf updates the shelf resource on the server.
	UpdateShelf(context.Context, *UpdateShelfRequest) (*UpdateShelfResponse, error)
	// DeleteShelf deletes the shelf resource from the server.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*DeleteShelfResponse, error)
}

// UnimplementedShelfServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShelfServiceServer struct {
}

func (UnimplementedShelfServiceServer) ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (UnimplementedShelfServiceServer) GetShelf(context.Context, *GetShelfRequest) (*GetShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (UnimplementedShelfServiceServer) CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (UnimplementedShelfServiceServer) UpdateShelf(context.Context, *UpdateShelfRequest) (*UpdateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShelf not implemented")
}
func (UnimplementedShelfServiceServer) DeleteShelf(context.Context, *DeleteShelfRequest) (*DeleteShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}

// UnsafeShelfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShelfServiceServer will
// result in compilation errors.
type UnsafeShelfServiceServer interface {
	mustEmbedUnimplementedShelfServiceServer()
}

func RegisterShelfServiceServer(s grpc.ServiceRegistrar, srv ShelfServiceServer) {
	s.RegisterService(&ShelfService_ServiceDesc, srv)
}

func _ShelfService_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.shelf.v1.ShelfService/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).ListShelves(ctx, req.(*ListShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.shelf.v1.ShelfService/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.shelf.v1.ShelfService/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_UpdateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).UpdateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.shelf.v1.ShelfService/UpdateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).UpdateShelf(ctx, req.(*UpdateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acme.shelf.v1.ShelfService/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShelfService_ServiceDesc is the grpc.ServiceDesc for ShelfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShelfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "acme.shelf.v1.ShelfService",
	HandlerType: (*ShelfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _ShelfService_ListShelves_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _ShelfService_GetShelf_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _ShelfService_CreateShelf_Handler,
		},
		{
			MethodName: "UpdateShelf",
			Handler:    _ShelfService_UpdateShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _ShelfService_DeleteShelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acme/shelf/v1/service.proto",
}
