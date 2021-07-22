// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementServiceClient interface {
	RegisterPeer(ctx context.Context, in *RegisterPeerRequest, opts ...grpc.CallOption) (*RegisterPeerResponse, error)
	// Sync enables peer synchronization. Each peer that is connected to this stream will receive updates from the server.
	// For example, if a new peer has been added to an account all other connected peers will receive this peer's Wireguard public key as an update
	// The initial SyncResponse contains all of the available peers so the local state can be refreshed
	Sync(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (ManagementService_SyncClient, error)
	// Exposes a Wireguard public key of the Management service.
	// This key is used to support message encryption between client and server
	GetServerKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerKeyResponse, error)
	// health check endpoint
	IsHealthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type managementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementServiceClient(cc grpc.ClientConnInterface) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) RegisterPeer(ctx context.Context, in *RegisterPeerRequest, opts ...grpc.CallOption) (*RegisterPeerResponse, error) {
	out := new(RegisterPeerResponse)
	err := c.cc.Invoke(ctx, "/management.ManagementService/RegisterPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) Sync(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (ManagementService_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManagementService_ServiceDesc.Streams[0], "/management.ManagementService/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &managementServiceSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManagementService_SyncClient interface {
	Recv() (*EncryptedMessage, error)
	grpc.ClientStream
}

type managementServiceSyncClient struct {
	grpc.ClientStream
}

func (x *managementServiceSyncClient) Recv() (*EncryptedMessage, error) {
	m := new(EncryptedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managementServiceClient) GetServerKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerKeyResponse, error) {
	out := new(ServerKeyResponse)
	err := c.cc.Invoke(ctx, "/management.ManagementService/GetServerKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) IsHealthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/management.ManagementService/isHealthy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
// All implementations must embed UnimplementedManagementServiceServer
// for forward compatibility
type ManagementServiceServer interface {
	RegisterPeer(context.Context, *RegisterPeerRequest) (*RegisterPeerResponse, error)
	// Sync enables peer synchronization. Each peer that is connected to this stream will receive updates from the server.
	// For example, if a new peer has been added to an account all other connected peers will receive this peer's Wireguard public key as an update
	// The initial SyncResponse contains all of the available peers so the local state can be refreshed
	Sync(*EncryptedMessage, ManagementService_SyncServer) error
	// Exposes a Wireguard public key of the Management service.
	// This key is used to support message encryption between client and server
	GetServerKey(context.Context, *Empty) (*ServerKeyResponse, error)
	// health check endpoint
	IsHealthy(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedManagementServiceServer()
}

// UnimplementedManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagementServiceServer struct {
}

func (UnimplementedManagementServiceServer) RegisterPeer(context.Context, *RegisterPeerRequest) (*RegisterPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPeer not implemented")
}
func (UnimplementedManagementServiceServer) Sync(*EncryptedMessage, ManagementService_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedManagementServiceServer) GetServerKey(context.Context, *Empty) (*ServerKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerKey not implemented")
}
func (UnimplementedManagementServiceServer) IsHealthy(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHealthy not implemented")
}
func (UnimplementedManagementServiceServer) mustEmbedUnimplementedManagementServiceServer() {}

// UnsafeManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServiceServer will
// result in compilation errors.
type UnsafeManagementServiceServer interface {
	mustEmbedUnimplementedManagementServiceServer()
}

func RegisterManagementServiceServer(s grpc.ServiceRegistrar, srv ManagementServiceServer) {
	s.RegisterService(&ManagementService_ServiceDesc, srv)
}

func _ManagementService_RegisterPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).RegisterPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.ManagementService/RegisterPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).RegisterPeer(ctx, req.(*RegisterPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EncryptedMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagementServiceServer).Sync(m, &managementServiceSyncServer{stream})
}

type ManagementService_SyncServer interface {
	Send(*EncryptedMessage) error
	grpc.ServerStream
}

type managementServiceSyncServer struct {
	grpc.ServerStream
}

func (x *managementServiceSyncServer) Send(m *EncryptedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ManagementService_GetServerKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).GetServerKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.ManagementService/GetServerKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).GetServerKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_IsHealthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).IsHealthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.ManagementService/isHealthy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).IsHealthy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagementService_ServiceDesc is the grpc.ServiceDesc for ManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPeer",
			Handler:    _ManagementService_RegisterPeer_Handler,
		},
		{
			MethodName: "GetServerKey",
			Handler:    _ManagementService_GetServerKey_Handler,
		},
		{
			MethodName: "isHealthy",
			Handler:    _ManagementService_IsHealthy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _ManagementService_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "management.proto",
}