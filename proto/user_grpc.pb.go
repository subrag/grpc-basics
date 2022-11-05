// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetProject(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*Project, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetProject(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/user.UserService/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetProject(context.Context, *UserProfile) (*Project, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetProject(context.Context, *UserProfile) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProject(ctx, req.(*UserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProject",
			Handler:    _UserService_GetProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	GetAssignment(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Assignment, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) GetAssignment(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Assignment, error) {
	out := new(Assignment)
	err := c.cc.Invoke(ctx, "/user.ProjectService/GetAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	GetAssignment(context.Context, *Project) (*Assignment, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) GetAssignment(context.Context, *Project) (*Assignment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignment not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_GetAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ProjectService/GetAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetAssignment(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssignment",
			Handler:    _ProjectService_GetAssignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}