// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Service_Login_FullMethodName                 = "/proto.Service/Login"
	Service_Register_FullMethodName              = "/proto.Service/Register"
	Service_GetUserByID_FullMethodName           = "/proto.Service/GetUserByID"
	Service_GetUserWatchedHistory_FullMethodName = "/proto.Service/GetUserWatchedHistory"
	Service_GetVideosByIds_FullMethodName        = "/proto.Service/GetVideosByIds"
	Service_GetVideoByID_FullMethodName          = "/proto.Service/GetVideoByID"
	Service_UpdateVideo_FullMethodName           = "/proto.Service/UpdateVideo"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*User, error)
	GetUserWatchedHistory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Videos, error)
	GetVideosByIds(ctx context.Context, in *GetVideosByIdsRequest, opts ...grpc.CallOption) (*Videos, error)
	GetVideoByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Video, error)
	UpdateVideo(ctx context.Context, in *Video, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Service_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Service_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetUserByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Service_GetUserByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetUserWatchedHistory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Videos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Videos)
	err := c.cc.Invoke(ctx, Service_GetUserWatchedHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetVideosByIds(ctx context.Context, in *GetVideosByIdsRequest, opts ...grpc.CallOption) (*Videos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Videos)
	err := c.cc.Invoke(ctx, Service_GetVideosByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetVideoByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Video, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Video)
	err := c.cc.Invoke(ctx, Service_GetVideoByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateVideo(ctx context.Context, in *Video, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*User, error)
	GetUserByID(context.Context, *IdRequest) (*User, error)
	GetUserWatchedHistory(context.Context, *IdRequest) (*Videos, error)
	GetVideosByIds(context.Context, *GetVideosByIdsRequest) (*Videos, error)
	GetVideoByID(context.Context, *IdRequest) (*Video, error)
	UpdateVideo(context.Context, *Video) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedServiceServer) Register(context.Context, *RegisterRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedServiceServer) GetUserByID(context.Context, *IdRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedServiceServer) GetUserWatchedHistory(context.Context, *IdRequest) (*Videos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWatchedHistory not implemented")
}
func (UnimplementedServiceServer) GetVideosByIds(context.Context, *GetVideosByIdsRequest) (*Videos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideosByIds not implemented")
}
func (UnimplementedServiceServer) GetVideoByID(context.Context, *IdRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoByID not implemented")
}
func (UnimplementedServiceServer) UpdateVideo(context.Context, *Video) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVideo not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}
func (UnimplementedServiceServer) testEmbeddedByValue()                 {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUserByID(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetUserWatchedHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUserWatchedHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetUserWatchedHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUserWatchedHistory(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetVideosByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideosByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetVideosByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetVideosByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetVideosByIds(ctx, req.(*GetVideosByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetVideoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetVideoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetVideoByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetVideoByID(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Video)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateVideo(ctx, req.(*Video))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Service_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Service_Register_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _Service_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserWatchedHistory",
			Handler:    _Service_GetUserWatchedHistory_Handler,
		},
		{
			MethodName: "GetVideosByIds",
			Handler:    _Service_GetVideosByIds_Handler,
		},
		{
			MethodName: "GetVideoByID",
			Handler:    _Service_GetVideoByID_Handler,
		},
		{
			MethodName: "UpdateVideo",
			Handler:    _Service_UpdateVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
