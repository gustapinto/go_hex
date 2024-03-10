// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: transaction.proto

package gen

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Create(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreatedResponse, error)
	DeleteByIDAndAccountID(ctx context.Context, in *DeleteTransactionByIDAndAccountIDRequest, opts ...grpc.CallOption) (*Empty, error)
	GetByAccountID(ctx context.Context, in *GetTransactionByAccountIDRequest, opts ...grpc.CallOption) (*RepeatedTransactionResponse, error)
	GetByIdAndAccountId(ctx context.Context, in *GetTransactionByIDAndAccountIDRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Create(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/TransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteByIDAndAccountID(ctx context.Context, in *DeleteTransactionByIDAndAccountIDRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/TransactionService/DeleteByIDAndAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetByAccountID(ctx context.Context, in *GetTransactionByAccountIDRequest, opts ...grpc.CallOption) (*RepeatedTransactionResponse, error) {
	out := new(RepeatedTransactionResponse)
	err := c.cc.Invoke(ctx, "/TransactionService/GetByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetByIdAndAccountId(ctx context.Context, in *GetTransactionByIDAndAccountIDRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/TransactionService/GetByIdAndAccountId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	Create(context.Context, *CreateTransactionRequest) (*CreatedResponse, error)
	DeleteByIDAndAccountID(context.Context, *DeleteTransactionByIDAndAccountIDRequest) (*Empty, error)
	GetByAccountID(context.Context, *GetTransactionByAccountIDRequest) (*RepeatedTransactionResponse, error)
	GetByIdAndAccountId(context.Context, *GetTransactionByIDAndAccountIDRequest) (*TransactionResponse, error)
}

// UnimplementedTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) Create(context.Context, *CreateTransactionRequest) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteByIDAndAccountID(context.Context, *DeleteTransactionByIDAndAccountIDRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByIDAndAccountID not implemented")
}
func (UnimplementedTransactionServiceServer) GetByAccountID(context.Context, *GetTransactionByAccountIDRequest) (*RepeatedTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccountID not implemented")
}
func (UnimplementedTransactionServiceServer) GetByIdAndAccountId(context.Context, *GetTransactionByIDAndAccountIDRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdAndAccountId not implemented")
}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Create(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteByIDAndAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionByIDAndAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteByIDAndAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/DeleteByIDAndAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteByIDAndAccountID(ctx, req.(*DeleteTransactionByIDAndAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/GetByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetByAccountID(ctx, req.(*GetTransactionByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetByIdAndAccountId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByIDAndAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetByIdAndAccountId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/GetByIdAndAccountId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetByIdAndAccountId(ctx, req.(*GetTransactionByIDAndAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TransactionService_Create_Handler,
		},
		{
			MethodName: "DeleteByIDAndAccountID",
			Handler:    _TransactionService_DeleteByIDAndAccountID_Handler,
		},
		{
			MethodName: "GetByAccountID",
			Handler:    _TransactionService_GetByAccountID_Handler,
		},
		{
			MethodName: "GetByIdAndAccountId",
			Handler:    _TransactionService_GetByIdAndAccountId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}