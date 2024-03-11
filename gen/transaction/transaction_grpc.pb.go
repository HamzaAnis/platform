// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: transaction.proto

package transaction

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
	Transaction_SayGoodbye_FullMethodName = "/transaction.Transaction/SayGoodbye"
)

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeReply, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeReply, error) {
	out := new(GoodbyeReply)
	err := c.cc.Invoke(ctx, Transaction_SayGoodbye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	SayGoodbye(context.Context, *GoodbyeRequest) (*GoodbyeReply, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) SayGoodbye(context.Context, *GoodbyeRequest) (*GoodbyeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_SayGoodbye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).SayGoodbye(ctx, req.(*GoodbyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayGoodbye",
			Handler:    _Transaction_SayGoodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
