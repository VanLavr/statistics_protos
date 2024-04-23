// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: stats.proto

package statistics_protos

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
	Statistics_CalculateStatistics_FullMethodName        = "/stats.Statistics/CalculateStatistics"
	Statistics_CalculateAverageStatistics_FullMethodName = "/stats.Statistics/CalculateAverageStatistics"
	Statistics_GenerateNumbers_FullMethodName            = "/stats.Statistics/GenerateNumbers"
	Statistics_CalculateBunchOfStatistics_FullMethodName = "/stats.Statistics/CalculateBunchOfStatistics"
)

// StatisticsClient is the client API for Statistics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsClient interface {
	CalculateStatistics(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatisticsForUser, error)
	CalculateAverageStatistics(ctx context.Context, opts ...grpc.CallOption) (Statistics_CalculateAverageStatisticsClient, error)
	GenerateNumbers(ctx context.Context, in *Limit, opts ...grpc.CallOption) (Statistics_GenerateNumbersClient, error)
	CalculateBunchOfStatistics(ctx context.Context, opts ...grpc.CallOption) (Statistics_CalculateBunchOfStatisticsClient, error)
}

type statisticsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticsClient(cc grpc.ClientConnInterface) StatisticsClient {
	return &statisticsClient{cc}
}

func (c *statisticsClient) CalculateStatistics(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatisticsForUser, error) {
	out := new(StatisticsForUser)
	err := c.cc.Invoke(ctx, Statistics_CalculateStatistics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsClient) CalculateAverageStatistics(ctx context.Context, opts ...grpc.CallOption) (Statistics_CalculateAverageStatisticsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Statistics_ServiceDesc.Streams[0], Statistics_CalculateAverageStatistics_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &statisticsCalculateAverageStatisticsClient{stream}
	return x, nil
}

type Statistics_CalculateAverageStatisticsClient interface {
	Send(*User) error
	CloseAndRecv() (*AvgStat, error)
	grpc.ClientStream
}

type statisticsCalculateAverageStatisticsClient struct {
	grpc.ClientStream
}

func (x *statisticsCalculateAverageStatisticsClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statisticsCalculateAverageStatisticsClient) CloseAndRecv() (*AvgStat, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgStat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statisticsClient) GenerateNumbers(ctx context.Context, in *Limit, opts ...grpc.CallOption) (Statistics_GenerateNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Statistics_ServiceDesc.Streams[1], Statistics_GenerateNumbers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &statisticsGenerateNumbersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Statistics_GenerateNumbersClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type statisticsGenerateNumbersClient struct {
	grpc.ClientStream
}

func (x *statisticsGenerateNumbersClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statisticsClient) CalculateBunchOfStatistics(ctx context.Context, opts ...grpc.CallOption) (Statistics_CalculateBunchOfStatisticsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Statistics_ServiceDesc.Streams[2], Statistics_CalculateBunchOfStatistics_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &statisticsCalculateBunchOfStatisticsClient{stream}
	return x, nil
}

type Statistics_CalculateBunchOfStatisticsClient interface {
	Send(*User) error
	Recv() (*StatisticsForUser, error)
	grpc.ClientStream
}

type statisticsCalculateBunchOfStatisticsClient struct {
	grpc.ClientStream
}

func (x *statisticsCalculateBunchOfStatisticsClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statisticsCalculateBunchOfStatisticsClient) Recv() (*StatisticsForUser, error) {
	m := new(StatisticsForUser)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatisticsServer is the server API for Statistics service.
// All implementations must embed UnimplementedStatisticsServer
// for forward compatibility
type StatisticsServer interface {
	CalculateStatistics(context.Context, *User) (*StatisticsForUser, error)
	CalculateAverageStatistics(Statistics_CalculateAverageStatisticsServer) error
	GenerateNumbers(*Limit, Statistics_GenerateNumbersServer) error
	CalculateBunchOfStatistics(Statistics_CalculateBunchOfStatisticsServer) error
	mustEmbedUnimplementedStatisticsServer()
}

// UnimplementedStatisticsServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServer struct {
}

func (UnimplementedStatisticsServer) CalculateStatistics(context.Context, *User) (*StatisticsForUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateStatistics not implemented")
}
func (UnimplementedStatisticsServer) CalculateAverageStatistics(Statistics_CalculateAverageStatisticsServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateAverageStatistics not implemented")
}
func (UnimplementedStatisticsServer) GenerateNumbers(*Limit, Statistics_GenerateNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateNumbers not implemented")
}
func (UnimplementedStatisticsServer) CalculateBunchOfStatistics(Statistics_CalculateBunchOfStatisticsServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateBunchOfStatistics not implemented")
}
func (UnimplementedStatisticsServer) mustEmbedUnimplementedStatisticsServer() {}

// UnsafeStatisticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticsServer will
// result in compilation errors.
type UnsafeStatisticsServer interface {
	mustEmbedUnimplementedStatisticsServer()
}

func RegisterStatisticsServer(s grpc.ServiceRegistrar, srv StatisticsServer) {
	s.RegisterService(&Statistics_ServiceDesc, srv)
}

func _Statistics_CalculateStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).CalculateStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Statistics_CalculateStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).CalculateStatistics(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Statistics_CalculateAverageStatistics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatisticsServer).CalculateAverageStatistics(&statisticsCalculateAverageStatisticsServer{stream})
}

type Statistics_CalculateAverageStatisticsServer interface {
	SendAndClose(*AvgStat) error
	Recv() (*User, error)
	grpc.ServerStream
}

type statisticsCalculateAverageStatisticsServer struct {
	grpc.ServerStream
}

func (x *statisticsCalculateAverageStatisticsServer) SendAndClose(m *AvgStat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statisticsCalculateAverageStatisticsServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Statistics_GenerateNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Limit)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatisticsServer).GenerateNumbers(m, &statisticsGenerateNumbersServer{stream})
}

type Statistics_GenerateNumbersServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type statisticsGenerateNumbersServer struct {
	grpc.ServerStream
}

func (x *statisticsGenerateNumbersServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func _Statistics_CalculateBunchOfStatistics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatisticsServer).CalculateBunchOfStatistics(&statisticsCalculateBunchOfStatisticsServer{stream})
}

type Statistics_CalculateBunchOfStatisticsServer interface {
	Send(*StatisticsForUser) error
	Recv() (*User, error)
	grpc.ServerStream
}

type statisticsCalculateBunchOfStatisticsServer struct {
	grpc.ServerStream
}

func (x *statisticsCalculateBunchOfStatisticsServer) Send(m *StatisticsForUser) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statisticsCalculateBunchOfStatisticsServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Statistics_ServiceDesc is the grpc.ServiceDesc for Statistics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Statistics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats.Statistics",
	HandlerType: (*StatisticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateStatistics",
			Handler:    _Statistics_CalculateStatistics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculateAverageStatistics",
			Handler:       _Statistics_CalculateAverageStatistics_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateNumbers",
			Handler:       _Statistics_GenerateNumbers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CalculateBunchOfStatistics",
			Handler:       _Statistics_CalculateBunchOfStatistics_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stats.proto",
}
