// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserByUsernameResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	ListUsersById(ctx context.Context, in *ListUsersByIdRequest, opts ...grpc.CallOption) (*ListUsersByIdResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserByUsernameResponse, error) {
	out := new(GetUserByUsernameResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListUsersById(ctx context.Context, in *ListUsersByIdRequest, opts ...grpc.CallOption) (*ListUsersByIdResponse, error) {
	out := new(ListUsersByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.UsersService/ListUsersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserByUsernameResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	ListUsersById(context.Context, *ListUsersByIdRequest) (*ListUsersByIdResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsersServiceServer) GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUsersServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUsersServiceServer) ListUsersById(context.Context, *ListUsersByIdRequest) (*ListUsersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersById not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListUsersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ListUsersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UsersService/ListUsersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ListUsersById(ctx, req.(*ListUsersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UsersService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UsersService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UsersService_GetUserByUsername_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UsersService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UsersService_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UsersService_ListUsers_Handler,
		},
		{
			MethodName: "ListUsersById",
			Handler:    _UsersService_ListUsersById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// FollowersServiceClient is the client API for FollowersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersServiceClient interface {
	CreateFollower(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error)
	GetFollower(ctx context.Context, in *GetFollowerRequest, opts ...grpc.CallOption) (*GetFollowerResponse, error)
	DeleteFollower(ctx context.Context, in *DeleteFollowerRequest, opts ...grpc.CallOption) (*DeleteFollowerResponse, error)
	ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error)
	ListFollowersById(ctx context.Context, in *ListFollowersByIdRequest, opts ...grpc.CallOption) (*ListFollowersByIdResponse, error)
	ListFollowersByLeaderId(ctx context.Context, in *ListFollowersByLeaderIdRequest, opts ...grpc.CallOption) (*ListFollowersByLeaderIdResponse, error)
}

type followersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersServiceClient(cc grpc.ClientConnInterface) FollowersServiceClient {
	return &followersServiceClient{cc}
}

func (c *followersServiceClient) CreateFollower(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error) {
	out := new(CreateFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/CreateFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollower(ctx context.Context, in *GetFollowerRequest, opts ...grpc.CallOption) (*GetFollowerResponse, error) {
	out := new(GetFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/GetFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) DeleteFollower(ctx context.Context, in *DeleteFollowerRequest, opts ...grpc.CallOption) (*DeleteFollowerResponse, error) {
	out := new(DeleteFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/DeleteFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error) {
	out := new(ListFollowersResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/ListFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) ListFollowersById(ctx context.Context, in *ListFollowersByIdRequest, opts ...grpc.CallOption) (*ListFollowersByIdResponse, error) {
	out := new(ListFollowersByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/ListFollowersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) ListFollowersByLeaderId(ctx context.Context, in *ListFollowersByLeaderIdRequest, opts ...grpc.CallOption) (*ListFollowersByLeaderIdResponse, error) {
	out := new(ListFollowersByLeaderIdResponse)
	err := c.cc.Invoke(ctx, "/pb.FollowersService/ListFollowersByLeaderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServiceServer is the server API for FollowersService service.
// All implementations must embed UnimplementedFollowersServiceServer
// for forward compatibility
type FollowersServiceServer interface {
	CreateFollower(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error)
	GetFollower(context.Context, *GetFollowerRequest) (*GetFollowerResponse, error)
	DeleteFollower(context.Context, *DeleteFollowerRequest) (*DeleteFollowerResponse, error)
	ListFollowers(context.Context, *ListFollowersRequest) (*ListFollowersResponse, error)
	ListFollowersById(context.Context, *ListFollowersByIdRequest) (*ListFollowersByIdResponse, error)
	ListFollowersByLeaderId(context.Context, *ListFollowersByLeaderIdRequest) (*ListFollowersByLeaderIdResponse, error)
	mustEmbedUnimplementedFollowersServiceServer()
}

// UnimplementedFollowersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServiceServer struct {
}

func (UnimplementedFollowersServiceServer) CreateFollower(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFollower not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollower(context.Context, *GetFollowerRequest) (*GetFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollower not implemented")
}
func (UnimplementedFollowersServiceServer) DeleteFollower(context.Context, *DeleteFollowerRequest) (*DeleteFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFollower not implemented")
}
func (UnimplementedFollowersServiceServer) ListFollowers(context.Context, *ListFollowersRequest) (*ListFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowers not implemented")
}
func (UnimplementedFollowersServiceServer) ListFollowersById(context.Context, *ListFollowersByIdRequest) (*ListFollowersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowersById not implemented")
}
func (UnimplementedFollowersServiceServer) ListFollowersByLeaderId(context.Context, *ListFollowersByLeaderIdRequest) (*ListFollowersByLeaderIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowersByLeaderId not implemented")
}
func (UnimplementedFollowersServiceServer) mustEmbedUnimplementedFollowersServiceServer() {}

// UnsafeFollowersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServiceServer will
// result in compilation errors.
type UnsafeFollowersServiceServer interface {
	mustEmbedUnimplementedFollowersServiceServer()
}

func RegisterFollowersServiceServer(s grpc.ServiceRegistrar, srv FollowersServiceServer) {
	s.RegisterService(&FollowersService_ServiceDesc, srv)
}

func _FollowersService_CreateFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).CreateFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/CreateFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).CreateFollower(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/GetFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollower(ctx, req.(*GetFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_DeleteFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).DeleteFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/DeleteFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).DeleteFollower(ctx, req.(*DeleteFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_ListFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).ListFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/ListFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).ListFollowers(ctx, req.(*ListFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_ListFollowersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).ListFollowersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/ListFollowersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).ListFollowersById(ctx, req.(*ListFollowersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_ListFollowersByLeaderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowersByLeaderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).ListFollowersByLeaderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FollowersService/ListFollowersByLeaderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).ListFollowersByLeaderId(ctx, req.(*ListFollowersByLeaderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowersService_ServiceDesc is the grpc.ServiceDesc for FollowersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FollowersService",
	HandlerType: (*FollowersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFollower",
			Handler:    _FollowersService_CreateFollower_Handler,
		},
		{
			MethodName: "GetFollower",
			Handler:    _FollowersService_GetFollower_Handler,
		},
		{
			MethodName: "DeleteFollower",
			Handler:    _FollowersService_DeleteFollower_Handler,
		},
		{
			MethodName: "ListFollowers",
			Handler:    _FollowersService_ListFollowers_Handler,
		},
		{
			MethodName: "ListFollowersById",
			Handler:    _FollowersService_ListFollowersById_Handler,
		},
		{
			MethodName: "ListFollowersByLeaderId",
			Handler:    _FollowersService_ListFollowersByLeaderId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}
