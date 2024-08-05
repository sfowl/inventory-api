// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/k8s_clusters_service.proto

package v1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	K8SClustersService_CreateK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.K8sClustersService/CreateK8sCluster"
	K8SClustersService_UpdateK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.K8sClustersService/UpdateK8sCluster"
	K8SClustersService_DeleteK8SCluster_FullMethodName = "/kessel.inventory.v1beta1.K8sClustersService/DeleteK8sCluster"
)

// K8SClustersServiceClient is the client API for K8SClustersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type K8SClustersServiceClient interface {
	CreateK8SCluster(ctx context.Context, in *CreateK8SClusterRequest, opts ...grpc.CallOption) (*CreateK8SClusterResponse, error)
	UpdateK8SCluster(ctx context.Context, in *UpdateK8SClusterRequest, opts ...grpc.CallOption) (*UpdateK8SClusterResponse, error)
	DeleteK8SCluster(ctx context.Context, in *DeleteK8SClusterRequest, opts ...grpc.CallOption) (*DeleteK8SClusterResponse, error)
}

type k8SClustersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SClustersServiceClient(cc grpc.ClientConnInterface) K8SClustersServiceClient {
	return &k8SClustersServiceClient{cc}
}

func (c *k8SClustersServiceClient) CreateK8SCluster(ctx context.Context, in *CreateK8SClusterRequest, opts ...grpc.CallOption) (*CreateK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateK8SClusterResponse)
	err := c.cc.Invoke(ctx, K8SClustersService_CreateK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SClustersServiceClient) UpdateK8SCluster(ctx context.Context, in *UpdateK8SClusterRequest, opts ...grpc.CallOption) (*UpdateK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateK8SClusterResponse)
	err := c.cc.Invoke(ctx, K8SClustersService_UpdateK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SClustersServiceClient) DeleteK8SCluster(ctx context.Context, in *DeleteK8SClusterRequest, opts ...grpc.CallOption) (*DeleteK8SClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteK8SClusterResponse)
	err := c.cc.Invoke(ctx, K8SClustersService_DeleteK8SCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SClustersServiceServer is the server API for K8SClustersService service.
// All implementations must embed UnimplementedK8SClustersServiceServer
// for forward compatibility.
type K8SClustersServiceServer interface {
	CreateK8SCluster(context.Context, *CreateK8SClusterRequest) (*CreateK8SClusterResponse, error)
	UpdateK8SCluster(context.Context, *UpdateK8SClusterRequest) (*UpdateK8SClusterResponse, error)
	DeleteK8SCluster(context.Context, *DeleteK8SClusterRequest) (*DeleteK8SClusterResponse, error)
	mustEmbedUnimplementedK8SClustersServiceServer()
}

// UnimplementedK8SClustersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedK8SClustersServiceServer struct{}

func (UnimplementedK8SClustersServiceServer) CreateK8SCluster(context.Context, *CreateK8SClusterRequest) (*CreateK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateK8SCluster not implemented")
}
func (UnimplementedK8SClustersServiceServer) UpdateK8SCluster(context.Context, *UpdateK8SClusterRequest) (*UpdateK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateK8SCluster not implemented")
}
func (UnimplementedK8SClustersServiceServer) DeleteK8SCluster(context.Context, *DeleteK8SClusterRequest) (*DeleteK8SClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteK8SCluster not implemented")
}
func (UnimplementedK8SClustersServiceServer) mustEmbedUnimplementedK8SClustersServiceServer() {}
func (UnimplementedK8SClustersServiceServer) testEmbeddedByValue()                            {}

// UnsafeK8SClustersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to K8SClustersServiceServer will
// result in compilation errors.
type UnsafeK8SClustersServiceServer interface {
	mustEmbedUnimplementedK8SClustersServiceServer()
}

func RegisterK8SClustersServiceServer(s grpc.ServiceRegistrar, srv K8SClustersServiceServer) {
	// If the following call pancis, it indicates UnimplementedK8SClustersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&K8SClustersService_ServiceDesc, srv)
}

func _K8SClustersService_CreateK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SClustersServiceServer).CreateK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: K8SClustersService_CreateK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SClustersServiceServer).CreateK8SCluster(ctx, req.(*CreateK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SClustersService_UpdateK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SClustersServiceServer).UpdateK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: K8SClustersService_UpdateK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SClustersServiceServer).UpdateK8SCluster(ctx, req.(*UpdateK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SClustersService_DeleteK8SCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteK8SClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SClustersServiceServer).DeleteK8SCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: K8SClustersService_DeleteK8SCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SClustersServiceServer).DeleteK8SCluster(ctx, req.(*DeleteK8SClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// K8SClustersService_ServiceDesc is the grpc.ServiceDesc for K8SClustersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var K8SClustersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kessel.inventory.v1beta1.K8sClustersService",
	HandlerType: (*K8SClustersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateK8sCluster",
			Handler:    _K8SClustersService_CreateK8SCluster_Handler,
		},
		{
			MethodName: "UpdateK8sCluster",
			Handler:    _K8SClustersService_UpdateK8SCluster_Handler,
		},
		{
			MethodName: "DeleteK8sCluster",
			Handler:    _K8SClustersService_DeleteK8SCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kessel/inventory/v1beta1/k8s_clusters_service.proto",
}
