// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: proto/hydrology_buffer_service.proto

package HL_BufferService

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
	HydrologyBufferService_AddTelegram_FullMethodName          = "/hydrologybuffer.HydrologyBufferService/AddTelegram"
	HydrologyBufferService_RemoveTelegrams_FullMethodName      = "/hydrologybuffer.HydrologyBufferService/RemoveTelegrams"
	HydrologyBufferService_UpdateTelegramByInfo_FullMethodName = "/hydrologybuffer.HydrologyBufferService/UpdateTelegramByInfo"
	HydrologyBufferService_UpdateTelegramByCode_FullMethodName = "/hydrologybuffer.HydrologyBufferService/UpdateTelegramByCode"
	HydrologyBufferService_GetTelegram_FullMethodName          = "/hydrologybuffer.HydrologyBufferService/GetTelegram"
	HydrologyBufferService_GetTelegrams_FullMethodName         = "/hydrologybuffer.HydrologyBufferService/GetTelegrams"
	HydrologyBufferService_TransferToSystem_FullMethodName     = "/hydrologybuffer.HydrologyBufferService/TransferToSystem"
)

// HydrologyBufferServiceClient is the client API for HydrologyBufferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HydrologyBufferServiceClient interface {
	AddTelegram(ctx context.Context, in *AddTelegramRequest, opts ...grpc.CallOption) (*AddTelegramResponse, error)
	RemoveTelegrams(ctx context.Context, in *RemoveTelegramsRequest, opts ...grpc.CallOption) (*RemoveTelegramsResponse, error)
	UpdateTelegramByInfo(ctx context.Context, in *UpdateTelegramByInfoRequest, opts ...grpc.CallOption) (*UpdateTelegramResponse, error)
	UpdateTelegramByCode(ctx context.Context, in *UpdateTelegramByCodeRequest, opts ...grpc.CallOption) (*UpdateTelegramResponse, error)
	GetTelegram(ctx context.Context, in *GetTelegramRequest, opts ...grpc.CallOption) (*GetTelegramResponse, error)
	GetTelegrams(ctx context.Context, in *GetTelegramsRequest, opts ...grpc.CallOption) (*GetTelegramsResponse, error)
	TransferToSystem(ctx context.Context, in *TransferToSystemRequest, opts ...grpc.CallOption) (*TransferToSystemResponse, error)
}

type hydrologyBufferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHydrologyBufferServiceClient(cc grpc.ClientConnInterface) HydrologyBufferServiceClient {
	return &hydrologyBufferServiceClient{cc}
}

func (c *hydrologyBufferServiceClient) AddTelegram(ctx context.Context, in *AddTelegramRequest, opts ...grpc.CallOption) (*AddTelegramResponse, error) {
	out := new(AddTelegramResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_AddTelegram_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) RemoveTelegrams(ctx context.Context, in *RemoveTelegramsRequest, opts ...grpc.CallOption) (*RemoveTelegramsResponse, error) {
	out := new(RemoveTelegramsResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_RemoveTelegrams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) UpdateTelegramByInfo(ctx context.Context, in *UpdateTelegramByInfoRequest, opts ...grpc.CallOption) (*UpdateTelegramResponse, error) {
	out := new(UpdateTelegramResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_UpdateTelegramByInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) UpdateTelegramByCode(ctx context.Context, in *UpdateTelegramByCodeRequest, opts ...grpc.CallOption) (*UpdateTelegramResponse, error) {
	out := new(UpdateTelegramResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_UpdateTelegramByCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) GetTelegram(ctx context.Context, in *GetTelegramRequest, opts ...grpc.CallOption) (*GetTelegramResponse, error) {
	out := new(GetTelegramResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_GetTelegram_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) GetTelegrams(ctx context.Context, in *GetTelegramsRequest, opts ...grpc.CallOption) (*GetTelegramsResponse, error) {
	out := new(GetTelegramsResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_GetTelegrams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hydrologyBufferServiceClient) TransferToSystem(ctx context.Context, in *TransferToSystemRequest, opts ...grpc.CallOption) (*TransferToSystemResponse, error) {
	out := new(TransferToSystemResponse)
	err := c.cc.Invoke(ctx, HydrologyBufferService_TransferToSystem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HydrologyBufferServiceServer is the server API for HydrologyBufferService service.
// All implementations must embed UnimplementedHydrologyBufferServiceServer
// for forward compatibility
type HydrologyBufferServiceServer interface {
	AddTelegram(context.Context, *AddTelegramRequest) (*AddTelegramResponse, error)
	RemoveTelegrams(context.Context, *RemoveTelegramsRequest) (*RemoveTelegramsResponse, error)
	UpdateTelegramByInfo(context.Context, *UpdateTelegramByInfoRequest) (*UpdateTelegramResponse, error)
	UpdateTelegramByCode(context.Context, *UpdateTelegramByCodeRequest) (*UpdateTelegramResponse, error)
	GetTelegram(context.Context, *GetTelegramRequest) (*GetTelegramResponse, error)
	GetTelegrams(context.Context, *GetTelegramsRequest) (*GetTelegramsResponse, error)
	TransferToSystem(context.Context, *TransferToSystemRequest) (*TransferToSystemResponse, error)
	mustEmbedUnimplementedHydrologyBufferServiceServer()
}

// UnimplementedHydrologyBufferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHydrologyBufferServiceServer struct {
}

func (UnimplementedHydrologyBufferServiceServer) AddTelegram(context.Context, *AddTelegramRequest) (*AddTelegramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTelegram not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) RemoveTelegrams(context.Context, *RemoveTelegramsRequest) (*RemoveTelegramsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTelegrams not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) UpdateTelegramByInfo(context.Context, *UpdateTelegramByInfoRequest) (*UpdateTelegramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTelegramByInfo not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) UpdateTelegramByCode(context.Context, *UpdateTelegramByCodeRequest) (*UpdateTelegramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTelegramByCode not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) GetTelegram(context.Context, *GetTelegramRequest) (*GetTelegramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelegram not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) GetTelegrams(context.Context, *GetTelegramsRequest) (*GetTelegramsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelegrams not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) TransferToSystem(context.Context, *TransferToSystemRequest) (*TransferToSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToSystem not implemented")
}
func (UnimplementedHydrologyBufferServiceServer) mustEmbedUnimplementedHydrologyBufferServiceServer() {
}

// UnsafeHydrologyBufferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HydrologyBufferServiceServer will
// result in compilation errors.
type UnsafeHydrologyBufferServiceServer interface {
	mustEmbedUnimplementedHydrologyBufferServiceServer()
}

func RegisterHydrologyBufferServiceServer(s grpc.ServiceRegistrar, srv HydrologyBufferServiceServer) {
	s.RegisterService(&HydrologyBufferService_ServiceDesc, srv)
}

func _HydrologyBufferService_AddTelegram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTelegramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).AddTelegram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_AddTelegram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).AddTelegram(ctx, req.(*AddTelegramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_RemoveTelegrams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTelegramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).RemoveTelegrams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_RemoveTelegrams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).RemoveTelegrams(ctx, req.(*RemoveTelegramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_UpdateTelegramByInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTelegramByInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).UpdateTelegramByInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_UpdateTelegramByInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).UpdateTelegramByInfo(ctx, req.(*UpdateTelegramByInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_UpdateTelegramByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTelegramByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).UpdateTelegramByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_UpdateTelegramByCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).UpdateTelegramByCode(ctx, req.(*UpdateTelegramByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_GetTelegram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTelegramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).GetTelegram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_GetTelegram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).GetTelegram(ctx, req.(*GetTelegramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_GetTelegrams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTelegramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).GetTelegrams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_GetTelegrams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).GetTelegrams(ctx, req.(*GetTelegramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HydrologyBufferService_TransferToSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferToSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HydrologyBufferServiceServer).TransferToSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HydrologyBufferService_TransferToSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HydrologyBufferServiceServer).TransferToSystem(ctx, req.(*TransferToSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HydrologyBufferService_ServiceDesc is the grpc.ServiceDesc for HydrologyBufferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HydrologyBufferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hydrologybuffer.HydrologyBufferService",
	HandlerType: (*HydrologyBufferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTelegram",
			Handler:    _HydrologyBufferService_AddTelegram_Handler,
		},
		{
			MethodName: "RemoveTelegrams",
			Handler:    _HydrologyBufferService_RemoveTelegrams_Handler,
		},
		{
			MethodName: "UpdateTelegramByInfo",
			Handler:    _HydrologyBufferService_UpdateTelegramByInfo_Handler,
		},
		{
			MethodName: "UpdateTelegramByCode",
			Handler:    _HydrologyBufferService_UpdateTelegramByCode_Handler,
		},
		{
			MethodName: "GetTelegram",
			Handler:    _HydrologyBufferService_GetTelegram_Handler,
		},
		{
			MethodName: "GetTelegrams",
			Handler:    _HydrologyBufferService_GetTelegrams_Handler,
		},
		{
			MethodName: "TransferToSystem",
			Handler:    _HydrologyBufferService_TransferToSystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hydrology_buffer_service.proto",
}
