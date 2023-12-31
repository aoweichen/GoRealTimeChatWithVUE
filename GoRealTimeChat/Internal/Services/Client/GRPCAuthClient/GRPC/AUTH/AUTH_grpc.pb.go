// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: AUTH.proto

package AUTHGRPC

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
	IMAuthHandler_CheckAuth_FullMethodName = "/IMAuthHandler/CheckAuth"
)

// IMAuthHandlerClient is the client API for IMAuthHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IMAuthHandlerClient interface {
	// CheckAuth是一个RPC方法，用于验证身份信息
	// 参数CheckAuthRequest是一个消息类型，表示包含token的请求
	// 返回CheckAuthResponse消息类型，表示包含验证结果的响应
	CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error)
}

type iMAuthHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewIMAuthHandlerClient(cc grpc.ClientConnInterface) IMAuthHandlerClient {
	return &iMAuthHandlerClient{cc}
}

func (c *iMAuthHandlerClient) CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error) {
	out := new(CheckAuthResponse)
	err := c.cc.Invoke(ctx, IMAuthHandler_CheckAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMAuthHandlerServer is the server API for IMAuthHandler service.
// All implementations must embed UnimplementedIMAuthHandlerServer
// for forward compatibility
type IMAuthHandlerServer interface {
	// CheckAuth是一个RPC方法，用于验证身份信息
	// 参数CheckAuthRequest是一个消息类型，表示包含token的请求
	// 返回CheckAuthResponse消息类型，表示包含验证结果的响应
	CheckAuth(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error)
	mustEmbedUnimplementedIMAuthHandlerServer()
}

// UnimplementedIMAuthHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedIMAuthHandlerServer struct {
}

func (UnimplementedIMAuthHandlerServer) CheckAuth(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedIMAuthHandlerServer) mustEmbedUnimplementedIMAuthHandlerServer() {}

// UnsafeIMAuthHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IMAuthHandlerServer will
// result in compilation errors.
type UnsafeIMAuthHandlerServer interface {
	mustEmbedUnimplementedIMAuthHandlerServer()
}

func RegisterIMAuthHandlerServer(s grpc.ServiceRegistrar, srv IMAuthHandlerServer) {
	s.RegisterService(&IMAuthHandler_ServiceDesc, srv)
}

func _IMAuthHandler_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMAuthHandlerServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IMAuthHandler_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMAuthHandlerServer).CheckAuth(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IMAuthHandler_ServiceDesc is the grpc.ServiceDesc for IMAuthHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IMAuthHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IMAuthHandler",
	HandlerType: (*IMAuthHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _IMAuthHandler_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AUTH.proto",
}
