// ----------------------------------------------------------------------------
// xr_debug_data.proto - XR debug data protobuf definitions
// Arpil 2024,Copyright (c) 2016 by Cisco Systems, Inc.
// ----------------------------------------------------------------------------

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.5.0
// source: serv.proto

package xrbinarysrvpb

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
	UploadService_UploadRequest_FullMethodName = "/xrbinarysrv.UploadService/UploadRequest"
)

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	// client side streaming supported.
	UploadRequest(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadRequestClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadRequest(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[0], UploadService_UploadRequest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadRequestClient{stream}
	return x, nil
}

type UploadService_UploadRequestClient interface {
	Send(*XrDebugRequest) error
	CloseAndRecv() (*XrDebugResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadRequestClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadRequestClient) Send(m *XrDebugRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadRequestClient) CloseAndRecv() (*XrDebugResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(XrDebugResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	// client side streaming supported.
	UploadRequest(UploadService_UploadRequestServer) error
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) UploadRequest(UploadService_UploadRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadRequest not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadRequest(&uploadServiceUploadRequestServer{stream})
}

type UploadService_UploadRequestServer interface {
	SendAndClose(*XrDebugResponse) error
	Recv() (*XrDebugRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadRequestServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadRequestServer) SendAndClose(m *XrDebugResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadRequestServer) Recv() (*XrDebugRequest, error) {
	m := new(XrDebugRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xrbinarysrv.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadRequest",
			Handler:       _UploadService_UploadRequest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "serv.proto",
}
