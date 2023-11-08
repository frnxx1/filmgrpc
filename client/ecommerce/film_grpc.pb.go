// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: film.proto

package ecommerce

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
	Films_GetGenre_FullMethodName = "/ecommerce.Films/GetGenre"
	Films_GetFilm_FullMethodName  = "/ecommerce.Films/GetFilm"
)

// FilmsClient is the client API for Films service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmsClient interface {
	GetGenre(ctx context.Context, in *FilmGenre, opts ...grpc.CallOption) (*FilmGenreRole, error)
	GetFilm(ctx context.Context, in *FilmInfo, opts ...grpc.CallOption) (*FilmStatus, error)
}

type filmsClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmsClient(cc grpc.ClientConnInterface) FilmsClient {
	return &filmsClient{cc}
}

func (c *filmsClient) GetGenre(ctx context.Context, in *FilmGenre, opts ...grpc.CallOption) (*FilmGenreRole, error) {
	out := new(FilmGenreRole)
	err := c.cc.Invoke(ctx, Films_GetGenre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetFilm(ctx context.Context, in *FilmInfo, opts ...grpc.CallOption) (*FilmStatus, error) {
	out := new(FilmStatus)
	err := c.cc.Invoke(ctx, Films_GetFilm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmsServer is the server API for Films service.
// All implementations must embed UnimplementedFilmsServer
// for forward compatibility
type FilmsServer interface {
	GetGenre(context.Context, *FilmGenre) (*FilmGenreRole, error)
	GetFilm(context.Context, *FilmInfo) (*FilmStatus, error)
	mustEmbedUnimplementedFilmsServer()
}

// UnimplementedFilmsServer must be embedded to have forward compatible implementations.
type UnimplementedFilmsServer struct {
}

func (UnimplementedFilmsServer) GetGenre(context.Context, *FilmGenre) (*FilmGenreRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenre not implemented")
}
func (UnimplementedFilmsServer) GetFilm(context.Context, *FilmInfo) (*FilmStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilm not implemented")
}
func (UnimplementedFilmsServer) mustEmbedUnimplementedFilmsServer() {}

// UnsafeFilmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmsServer will
// result in compilation errors.
type UnsafeFilmsServer interface {
	mustEmbedUnimplementedFilmsServer()
}

func RegisterFilmsServer(s grpc.ServiceRegistrar, srv FilmsServer) {
	s.RegisterService(&Films_ServiceDesc, srv)
}

func _Films_GetGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmGenre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Films_GetGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetGenre(ctx, req.(*FilmGenre))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Films_GetFilm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetFilm(ctx, req.(*FilmInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Films_ServiceDesc is the grpc.ServiceDesc for Films service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Films_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.Films",
	HandlerType: (*FilmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenre",
			Handler:    _Films_GetGenre_Handler,
		},
		{
			MethodName: "GetFilm",
			Handler:    _Films_GetFilm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "film.proto",
}
