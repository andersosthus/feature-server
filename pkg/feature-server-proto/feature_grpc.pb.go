// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package featureserverpb

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

// FeatureClient is the client API for Feature service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureClient interface {
	IsEnabled(ctx context.Context, in *FeatureRequest, opts ...grpc.CallOption) (*FeatureResponse, error)
}

type featureClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureClient(cc grpc.ClientConnInterface) FeatureClient {
	return &featureClient{cc}
}

func (c *featureClient) IsEnabled(ctx context.Context, in *FeatureRequest, opts ...grpc.CallOption) (*FeatureResponse, error) {
	out := new(FeatureResponse)
	err := c.cc.Invoke(ctx, "/featureserver.Feature/IsEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureServer is the server API for Feature service.
// All implementations must embed UnimplementedFeatureServer
// for forward compatibility
type FeatureServer interface {
	IsEnabled(context.Context, *FeatureRequest) (*FeatureResponse, error)
	mustEmbedUnimplementedFeatureServer()
}

// UnimplementedFeatureServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureServer struct {
}

func (UnimplementedFeatureServer) IsEnabled(context.Context, *FeatureRequest) (*FeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEnabled not implemented")
}
func (UnimplementedFeatureServer) mustEmbedUnimplementedFeatureServer() {}

// UnsafeFeatureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureServer will
// result in compilation errors.
type UnsafeFeatureServer interface {
	mustEmbedUnimplementedFeatureServer()
}

func RegisterFeatureServer(s grpc.ServiceRegistrar, srv FeatureServer) {
	s.RegisterService(&Feature_ServiceDesc, srv)
}

func _Feature_IsEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureServer).IsEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/featureserver.Feature/IsEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureServer).IsEnabled(ctx, req.(*FeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feature_ServiceDesc is the grpc.ServiceDesc for Feature service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feature_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "featureserver.Feature",
	HandlerType: (*FeatureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsEnabled",
			Handler:    _Feature_IsEnabled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feature-server-proto/feature.proto",
}