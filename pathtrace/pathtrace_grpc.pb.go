// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pathtrace

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

// PathTracerClient is the client API for PathTracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PathTracerClient interface {
	// PathTraceJobRequest comes from a path trace agent, it is a stream of job
	// requests and job replies. Each new request carries the result of a previous
	// job.
	PathTraceJobRequest(ctx context.Context, opts ...grpc.CallOption) (PathTracer_PathTraceJobRequestClient, error)
	// PathTraceRequest is initiated by a client application
	PathTraceRequest(ctx context.Context, in *PathTraceJob, opts ...grpc.CallOption) (*PathTraceReply, error)
}

type pathTracerClient struct {
	cc grpc.ClientConnInterface
}

func NewPathTracerClient(cc grpc.ClientConnInterface) PathTracerClient {
	return &pathTracerClient{cc}
}

func (c *pathTracerClient) PathTraceJobRequest(ctx context.Context, opts ...grpc.CallOption) (PathTracer_PathTraceJobRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &PathTracer_ServiceDesc.Streams[0], "/pathtrace.PathTracer/PathTraceJobRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &pathTracerPathTraceJobRequestClient{stream}
	return x, nil
}

type PathTracer_PathTraceJobRequestClient interface {
	Send(*GetPathTraceJob) error
	Recv() (*PathTraceJob, error)
	grpc.ClientStream
}

type pathTracerPathTraceJobRequestClient struct {
	grpc.ClientStream
}

func (x *pathTracerPathTraceJobRequestClient) Send(m *GetPathTraceJob) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pathTracerPathTraceJobRequestClient) Recv() (*PathTraceJob, error) {
	m := new(PathTraceJob)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pathTracerClient) PathTraceRequest(ctx context.Context, in *PathTraceJob, opts ...grpc.CallOption) (*PathTraceReply, error) {
	out := new(PathTraceReply)
	err := c.cc.Invoke(ctx, "/pathtrace.PathTracer/PathTraceRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PathTracerServer is the server API for PathTracer service.
// All implementations must embed UnimplementedPathTracerServer
// for forward compatibility
type PathTracerServer interface {
	// PathTraceJobRequest comes from a path trace agent, it is a stream of job
	// requests and job replies. Each new request carries the result of a previous
	// job.
	PathTraceJobRequest(PathTracer_PathTraceJobRequestServer) error
	// PathTraceRequest is initiated by a client application
	PathTraceRequest(context.Context, *PathTraceJob) (*PathTraceReply, error)
	mustEmbedUnimplementedPathTracerServer()
}

// UnimplementedPathTracerServer must be embedded to have forward compatible implementations.
type UnimplementedPathTracerServer struct {
}

func (UnimplementedPathTracerServer) PathTraceJobRequest(PathTracer_PathTraceJobRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method PathTraceJobRequest not implemented")
}
func (UnimplementedPathTracerServer) PathTraceRequest(context.Context, *PathTraceJob) (*PathTraceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PathTraceRequest not implemented")
}
func (UnimplementedPathTracerServer) mustEmbedUnimplementedPathTracerServer() {}

// UnsafePathTracerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PathTracerServer will
// result in compilation errors.
type UnsafePathTracerServer interface {
	mustEmbedUnimplementedPathTracerServer()
}

func RegisterPathTracerServer(s grpc.ServiceRegistrar, srv PathTracerServer) {
	s.RegisterService(&PathTracer_ServiceDesc, srv)
}

func _PathTracer_PathTraceJobRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PathTracerServer).PathTraceJobRequest(&pathTracerPathTraceJobRequestServer{stream})
}

type PathTracer_PathTraceJobRequestServer interface {
	Send(*PathTraceJob) error
	Recv() (*GetPathTraceJob, error)
	grpc.ServerStream
}

type pathTracerPathTraceJobRequestServer struct {
	grpc.ServerStream
}

func (x *pathTracerPathTraceJobRequestServer) Send(m *PathTraceJob) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pathTracerPathTraceJobRequestServer) Recv() (*GetPathTraceJob, error) {
	m := new(GetPathTraceJob)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PathTracer_PathTraceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PathTraceJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PathTracerServer).PathTraceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathtrace.PathTracer/PathTraceRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PathTracerServer).PathTraceRequest(ctx, req.(*PathTraceJob))
	}
	return interceptor(ctx, in, info, handler)
}

// PathTracer_ServiceDesc is the grpc.ServiceDesc for PathTracer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PathTracer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pathtrace.PathTracer",
	HandlerType: (*PathTracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PathTraceRequest",
			Handler:    _PathTracer_PathTraceRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PathTraceJobRequest",
			Handler:       _PathTracer_PathTraceJobRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pathtrace.proto",
}