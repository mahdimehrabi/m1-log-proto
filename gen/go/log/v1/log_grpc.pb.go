// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: log/v1/log.proto

package petv1

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
	LogService_StoreLog_FullMethodName = "/pet.v1.LogService/StoreLog"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	StoreLog(ctx context.Context, opts ...grpc.CallOption) (LogService_StoreLogClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) StoreLog(ctx context.Context, opts ...grpc.CallOption) (LogService_StoreLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], LogService_StoreLog_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceStoreLogClient{stream}
	return x, nil
}

type LogService_StoreLogClient interface {
	Send(*Log) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type logServiceStoreLogClient struct {
	grpc.ClientStream
}

func (x *logServiceStoreLogClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logServiceStoreLogClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	StoreLog(LogService_StoreLogServer) error
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) StoreLog(LogService_StoreLogServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreLog not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_StoreLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).StoreLog(&logServiceStoreLogServer{stream})
}

type LogService_StoreLogServer interface {
	SendAndClose(*Empty) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type logServiceStoreLogServer struct {
	grpc.ServerStream
}

func (x *logServiceStoreLogServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logServiceStoreLogServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet.v1.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreLog",
			Handler:       _LogService_StoreLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "log/v1/log.proto",
}
