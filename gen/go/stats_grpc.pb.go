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
	Statistics_CalculateStatistics_FullMethodName = "/stats.Statistics/CalculateStatistics"
)

// StatisticsClient is the client API for Statistics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticsClient interface {
	CalculateStatistics(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatisticsForUser, error)
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

// StatisticsServer is the server API for Statistics service.
// All implementations must embed UnimplementedStatisticsServer
// for forward compatibility
type StatisticsServer interface {
	CalculateStatistics(context.Context, *User) (*StatisticsForUser, error)
	mustEmbedUnimplementedStatisticsServer()
}

// UnimplementedStatisticsServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticsServer struct {
}

func (UnimplementedStatisticsServer) CalculateStatistics(context.Context, *User) (*StatisticsForUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateStatistics not implemented")
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats.proto",
}
