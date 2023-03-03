// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: pb/book/book-service.proto

package book

import (
	pb "book/pb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Create(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookCreateResponse, error)
	GetBook(ctx context.Context, in *BookFindOneRequest, opts ...grpc.CallOption) (*BookFindOneResponse, error)
	GetBooks(ctx context.Context, in *BookFindAllRequest, opts ...grpc.CallOption) (*BookFindAllResponse, error)
	Delete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error)
	SoftDelete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error)
	Update(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error)
	Restore(ctx context.Context, in *BookRestoreRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Create(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookCreateResponse, error) {
	out := new(BookCreateResponse)
	err := c.cc.Invoke(ctx, "/BookService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *BookFindOneRequest, opts ...grpc.CallOption) (*BookFindOneResponse, error) {
	out := new(BookFindOneResponse)
	err := c.cc.Invoke(ctx, "/BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBooks(ctx context.Context, in *BookFindAllRequest, opts ...grpc.CallOption) (*BookFindAllResponse, error) {
	out := new(BookFindAllResponse)
	err := c.cc.Invoke(ctx, "/BookService/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error) {
	out := new(pb.OperationResponse)
	err := c.cc.Invoke(ctx, "/BookService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) SoftDelete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error) {
	out := new(pb.OperationResponse)
	err := c.cc.Invoke(ctx, "/BookService/SoftDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Update(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error) {
	out := new(pb.OperationResponse)
	err := c.cc.Invoke(ctx, "/BookService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Restore(ctx context.Context, in *BookRestoreRequest, opts ...grpc.CallOption) (*pb.OperationResponse, error) {
	out := new(pb.OperationResponse)
	err := c.cc.Invoke(ctx, "/BookService/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	Create(context.Context, *BookCreateRequest) (*BookCreateResponse, error)
	GetBook(context.Context, *BookFindOneRequest) (*BookFindOneResponse, error)
	GetBooks(context.Context, *BookFindAllRequest) (*BookFindAllResponse, error)
	Delete(context.Context, *BookDeleteRequest) (*pb.OperationResponse, error)
	SoftDelete(context.Context, *BookDeleteRequest) (*pb.OperationResponse, error)
	Update(context.Context, *BookUpdateRequest) (*pb.OperationResponse, error)
	Restore(context.Context, *BookRestoreRequest) (*pb.OperationResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) Create(context.Context, *BookCreateRequest) (*BookCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookServiceServer) GetBook(context.Context, *BookFindOneRequest) (*BookFindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) GetBooks(context.Context, *BookFindAllRequest) (*BookFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBookServiceServer) Delete(context.Context, *BookDeleteRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookServiceServer) SoftDelete(context.Context, *BookDeleteRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDelete not implemented")
}
func (UnimplementedBookServiceServer) Update(context.Context, *BookUpdateRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookServiceServer) Restore(context.Context, *BookRestoreRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Create(ctx, req.(*BookCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookFindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*BookFindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBooks(ctx, req.(*BookFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_SoftDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).SoftDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/SoftDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).SoftDelete(ctx, req.(*BookDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Update(ctx, req.(*BookUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Restore(ctx, req.(*BookRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookService_Create_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _BookService_GetBooks_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
		{
			MethodName: "SoftDelete",
			Handler:    _BookService_SoftDelete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookService_Update_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _BookService_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/book/book-service.proto",
}
