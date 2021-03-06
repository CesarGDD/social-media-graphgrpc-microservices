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

// HashtagsServiceClient is the client API for HashtagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashtagsServiceClient interface {
	CreateHashtag(ctx context.Context, in *CreateHashtagRequest, opts ...grpc.CallOption) (*CreateHashtagResponse, error)
	GetHashtagByTitle(ctx context.Context, in *GetHashtagByTitleRequest, opts ...grpc.CallOption) (*GetHashtagByTitleResponse, error)
	GetHashtag(ctx context.Context, in *GetHashtagRequest, opts ...grpc.CallOption) (*GetHashtagResponse, error)
	UpdateHashtag(ctx context.Context, in *UpdateHashtagRequest, opts ...grpc.CallOption) (*UpdateHashtagResponse, error)
	DeleteHashtag(ctx context.Context, in *DeleteHashtagRequest, opts ...grpc.CallOption) (*DeleteHashtagResponse, error)
	ListHashtags(ctx context.Context, in *ListHashtagsRequest, opts ...grpc.CallOption) (*ListHashtagsResponse, error)
	ListHashtagsById(ctx context.Context, in *ListHashtagsByIdRequest, opts ...grpc.CallOption) (*ListHashtagsByIdResponse, error)
}

type hashtagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHashtagsServiceClient(cc grpc.ClientConnInterface) HashtagsServiceClient {
	return &hashtagsServiceClient{cc}
}

func (c *hashtagsServiceClient) CreateHashtag(ctx context.Context, in *CreateHashtagRequest, opts ...grpc.CallOption) (*CreateHashtagResponse, error) {
	out := new(CreateHashtagResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/CreateHashtag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) GetHashtagByTitle(ctx context.Context, in *GetHashtagByTitleRequest, opts ...grpc.CallOption) (*GetHashtagByTitleResponse, error) {
	out := new(GetHashtagByTitleResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/GetHashtagByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) GetHashtag(ctx context.Context, in *GetHashtagRequest, opts ...grpc.CallOption) (*GetHashtagResponse, error) {
	out := new(GetHashtagResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/GetHashtag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) UpdateHashtag(ctx context.Context, in *UpdateHashtagRequest, opts ...grpc.CallOption) (*UpdateHashtagResponse, error) {
	out := new(UpdateHashtagResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/UpdateHashtag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) DeleteHashtag(ctx context.Context, in *DeleteHashtagRequest, opts ...grpc.CallOption) (*DeleteHashtagResponse, error) {
	out := new(DeleteHashtagResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/DeleteHashtag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) ListHashtags(ctx context.Context, in *ListHashtagsRequest, opts ...grpc.CallOption) (*ListHashtagsResponse, error) {
	out := new(ListHashtagsResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/ListHashtags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagsServiceClient) ListHashtagsById(ctx context.Context, in *ListHashtagsByIdRequest, opts ...grpc.CallOption) (*ListHashtagsByIdResponse, error) {
	out := new(ListHashtagsByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagsService/ListHashtagsById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashtagsServiceServer is the server API for HashtagsService service.
// All implementations must embed UnimplementedHashtagsServiceServer
// for forward compatibility
type HashtagsServiceServer interface {
	CreateHashtag(context.Context, *CreateHashtagRequest) (*CreateHashtagResponse, error)
	GetHashtagByTitle(context.Context, *GetHashtagByTitleRequest) (*GetHashtagByTitleResponse, error)
	GetHashtag(context.Context, *GetHashtagRequest) (*GetHashtagResponse, error)
	UpdateHashtag(context.Context, *UpdateHashtagRequest) (*UpdateHashtagResponse, error)
	DeleteHashtag(context.Context, *DeleteHashtagRequest) (*DeleteHashtagResponse, error)
	ListHashtags(context.Context, *ListHashtagsRequest) (*ListHashtagsResponse, error)
	ListHashtagsById(context.Context, *ListHashtagsByIdRequest) (*ListHashtagsByIdResponse, error)
	mustEmbedUnimplementedHashtagsServiceServer()
}

// UnimplementedHashtagsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHashtagsServiceServer struct {
}

func (UnimplementedHashtagsServiceServer) CreateHashtag(context.Context, *CreateHashtagRequest) (*CreateHashtagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHashtag not implemented")
}
func (UnimplementedHashtagsServiceServer) GetHashtagByTitle(context.Context, *GetHashtagByTitleRequest) (*GetHashtagByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashtagByTitle not implemented")
}
func (UnimplementedHashtagsServiceServer) GetHashtag(context.Context, *GetHashtagRequest) (*GetHashtagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashtag not implemented")
}
func (UnimplementedHashtagsServiceServer) UpdateHashtag(context.Context, *UpdateHashtagRequest) (*UpdateHashtagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHashtag not implemented")
}
func (UnimplementedHashtagsServiceServer) DeleteHashtag(context.Context, *DeleteHashtagRequest) (*DeleteHashtagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHashtag not implemented")
}
func (UnimplementedHashtagsServiceServer) ListHashtags(context.Context, *ListHashtagsRequest) (*ListHashtagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHashtags not implemented")
}
func (UnimplementedHashtagsServiceServer) ListHashtagsById(context.Context, *ListHashtagsByIdRequest) (*ListHashtagsByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHashtagsById not implemented")
}
func (UnimplementedHashtagsServiceServer) mustEmbedUnimplementedHashtagsServiceServer() {}

// UnsafeHashtagsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashtagsServiceServer will
// result in compilation errors.
type UnsafeHashtagsServiceServer interface {
	mustEmbedUnimplementedHashtagsServiceServer()
}

func RegisterHashtagsServiceServer(s grpc.ServiceRegistrar, srv HashtagsServiceServer) {
	s.RegisterService(&HashtagsService_ServiceDesc, srv)
}

func _HashtagsService_CreateHashtag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHashtagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).CreateHashtag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/CreateHashtag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).CreateHashtag(ctx, req.(*CreateHashtagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_GetHashtagByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashtagByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).GetHashtagByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/GetHashtagByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).GetHashtagByTitle(ctx, req.(*GetHashtagByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_GetHashtag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashtagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).GetHashtag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/GetHashtag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).GetHashtag(ctx, req.(*GetHashtagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_UpdateHashtag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHashtagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).UpdateHashtag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/UpdateHashtag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).UpdateHashtag(ctx, req.(*UpdateHashtagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_DeleteHashtag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHashtagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).DeleteHashtag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/DeleteHashtag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).DeleteHashtag(ctx, req.(*DeleteHashtagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_ListHashtags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHashtagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).ListHashtags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/ListHashtags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).ListHashtags(ctx, req.(*ListHashtagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagsService_ListHashtagsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHashtagsByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagsServiceServer).ListHashtagsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagsService/ListHashtagsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagsServiceServer).ListHashtagsById(ctx, req.(*ListHashtagsByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HashtagsService_ServiceDesc is the grpc.ServiceDesc for HashtagsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashtagsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HashtagsService",
	HandlerType: (*HashtagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHashtag",
			Handler:    _HashtagsService_CreateHashtag_Handler,
		},
		{
			MethodName: "GetHashtagByTitle",
			Handler:    _HashtagsService_GetHashtagByTitle_Handler,
		},
		{
			MethodName: "GetHashtag",
			Handler:    _HashtagsService_GetHashtag_Handler,
		},
		{
			MethodName: "UpdateHashtag",
			Handler:    _HashtagsService_UpdateHashtag_Handler,
		},
		{
			MethodName: "DeleteHashtag",
			Handler:    _HashtagsService_DeleteHashtag_Handler,
		},
		{
			MethodName: "ListHashtags",
			Handler:    _HashtagsService_ListHashtags_Handler,
		},
		{
			MethodName: "ListHashtagsById",
			Handler:    _HashtagsService_ListHashtagsById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// HashtagPostsServiceClient is the client API for HashtagPostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashtagPostsServiceClient interface {
	CreateHashtagPost(ctx context.Context, in *CreateHashtagPostRequest, opts ...grpc.CallOption) (*CreateHashtagPostResponse, error)
	GetHashtagPost(ctx context.Context, in *GetHashtagPostRequest, opts ...grpc.CallOption) (*GetHashtagPostResponse, error)
	DeleteHashtagPost(ctx context.Context, in *DeleteHashtagPostRequest, opts ...grpc.CallOption) (*DeleteHashtagPostResponse, error)
	ListHashtagPosts(ctx context.Context, in *ListHashtagPostsRequest, opts ...grpc.CallOption) (*ListHashtagPostsResponse, error)
	ListHashtagPostsById(ctx context.Context, in *ListHashtagPostsByIdRequest, opts ...grpc.CallOption) (*ListHashtagPostsByIdResponse, error)
}

type hashtagPostsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHashtagPostsServiceClient(cc grpc.ClientConnInterface) HashtagPostsServiceClient {
	return &hashtagPostsServiceClient{cc}
}

func (c *hashtagPostsServiceClient) CreateHashtagPost(ctx context.Context, in *CreateHashtagPostRequest, opts ...grpc.CallOption) (*CreateHashtagPostResponse, error) {
	out := new(CreateHashtagPostResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagPostsService/CreateHashtagPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagPostsServiceClient) GetHashtagPost(ctx context.Context, in *GetHashtagPostRequest, opts ...grpc.CallOption) (*GetHashtagPostResponse, error) {
	out := new(GetHashtagPostResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagPostsService/GetHashtagPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagPostsServiceClient) DeleteHashtagPost(ctx context.Context, in *DeleteHashtagPostRequest, opts ...grpc.CallOption) (*DeleteHashtagPostResponse, error) {
	out := new(DeleteHashtagPostResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagPostsService/DeleteHashtagPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagPostsServiceClient) ListHashtagPosts(ctx context.Context, in *ListHashtagPostsRequest, opts ...grpc.CallOption) (*ListHashtagPostsResponse, error) {
	out := new(ListHashtagPostsResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagPostsService/ListHashtagPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagPostsServiceClient) ListHashtagPostsById(ctx context.Context, in *ListHashtagPostsByIdRequest, opts ...grpc.CallOption) (*ListHashtagPostsByIdResponse, error) {
	out := new(ListHashtagPostsByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.HashtagPostsService/ListHashtagPostsById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashtagPostsServiceServer is the server API for HashtagPostsService service.
// All implementations must embed UnimplementedHashtagPostsServiceServer
// for forward compatibility
type HashtagPostsServiceServer interface {
	CreateHashtagPost(context.Context, *CreateHashtagPostRequest) (*CreateHashtagPostResponse, error)
	GetHashtagPost(context.Context, *GetHashtagPostRequest) (*GetHashtagPostResponse, error)
	DeleteHashtagPost(context.Context, *DeleteHashtagPostRequest) (*DeleteHashtagPostResponse, error)
	ListHashtagPosts(context.Context, *ListHashtagPostsRequest) (*ListHashtagPostsResponse, error)
	ListHashtagPostsById(context.Context, *ListHashtagPostsByIdRequest) (*ListHashtagPostsByIdResponse, error)
	mustEmbedUnimplementedHashtagPostsServiceServer()
}

// UnimplementedHashtagPostsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHashtagPostsServiceServer struct {
}

func (UnimplementedHashtagPostsServiceServer) CreateHashtagPost(context.Context, *CreateHashtagPostRequest) (*CreateHashtagPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHashtagPost not implemented")
}
func (UnimplementedHashtagPostsServiceServer) GetHashtagPost(context.Context, *GetHashtagPostRequest) (*GetHashtagPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashtagPost not implemented")
}
func (UnimplementedHashtagPostsServiceServer) DeleteHashtagPost(context.Context, *DeleteHashtagPostRequest) (*DeleteHashtagPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHashtagPost not implemented")
}
func (UnimplementedHashtagPostsServiceServer) ListHashtagPosts(context.Context, *ListHashtagPostsRequest) (*ListHashtagPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHashtagPosts not implemented")
}
func (UnimplementedHashtagPostsServiceServer) ListHashtagPostsById(context.Context, *ListHashtagPostsByIdRequest) (*ListHashtagPostsByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHashtagPostsById not implemented")
}
func (UnimplementedHashtagPostsServiceServer) mustEmbedUnimplementedHashtagPostsServiceServer() {}

// UnsafeHashtagPostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashtagPostsServiceServer will
// result in compilation errors.
type UnsafeHashtagPostsServiceServer interface {
	mustEmbedUnimplementedHashtagPostsServiceServer()
}

func RegisterHashtagPostsServiceServer(s grpc.ServiceRegistrar, srv HashtagPostsServiceServer) {
	s.RegisterService(&HashtagPostsService_ServiceDesc, srv)
}

func _HashtagPostsService_CreateHashtagPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHashtagPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagPostsServiceServer).CreateHashtagPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagPostsService/CreateHashtagPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagPostsServiceServer).CreateHashtagPost(ctx, req.(*CreateHashtagPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagPostsService_GetHashtagPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashtagPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagPostsServiceServer).GetHashtagPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagPostsService/GetHashtagPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagPostsServiceServer).GetHashtagPost(ctx, req.(*GetHashtagPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagPostsService_DeleteHashtagPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHashtagPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagPostsServiceServer).DeleteHashtagPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagPostsService/DeleteHashtagPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagPostsServiceServer).DeleteHashtagPost(ctx, req.(*DeleteHashtagPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagPostsService_ListHashtagPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHashtagPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagPostsServiceServer).ListHashtagPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagPostsService/ListHashtagPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagPostsServiceServer).ListHashtagPosts(ctx, req.(*ListHashtagPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagPostsService_ListHashtagPostsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHashtagPostsByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagPostsServiceServer).ListHashtagPostsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HashtagPostsService/ListHashtagPostsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagPostsServiceServer).ListHashtagPostsById(ctx, req.(*ListHashtagPostsByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HashtagPostsService_ServiceDesc is the grpc.ServiceDesc for HashtagPostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashtagPostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HashtagPostsService",
	HandlerType: (*HashtagPostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHashtagPost",
			Handler:    _HashtagPostsService_CreateHashtagPost_Handler,
		},
		{
			MethodName: "GetHashtagPost",
			Handler:    _HashtagPostsService_GetHashtagPost_Handler,
		},
		{
			MethodName: "DeleteHashtagPost",
			Handler:    _HashtagPostsService_DeleteHashtagPost_Handler,
		},
		{
			MethodName: "ListHashtagPosts",
			Handler:    _HashtagPostsService_ListHashtagPosts_Handler,
		},
		{
			MethodName: "ListHashtagPostsById",
			Handler:    _HashtagPostsService_ListHashtagPostsById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}
