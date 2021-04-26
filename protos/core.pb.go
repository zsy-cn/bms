// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	GetGroupList(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GroupList, error)
	AddGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error)
	UpdateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	GetManufacturerList(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*ManufacturerList, error)
	AddManufacturer(ctx context.Context, in *Manufacturer, opts ...grpc.CallOption) (*Empty, error)
	UpdateManufacturer(ctx context.Context, in *Manufacturer, opts ...grpc.CallOption) (*Empty, error)
	DeleteManufacturer(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	GetDeviceTypeList(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*DeviceTypeList, error)
	AddDeviceType(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*Empty, error)
	UpdateDeviceType(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*Empty, error)
	DeleteDeviceType(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	GetDeviceModelList(ctx context.Context, in *GetDeviceModelsRequest, opts ...grpc.CallOption) (*DeviceModelList, error)
	AddDeviceModel(ctx context.Context, in *DeviceModel, opts ...grpc.CallOption) (*Empty, error)
	UpdateDeviceModel(ctx context.Context, in *DeviceModel, opts ...grpc.CallOption) (*Empty, error)
	DeleteDeviceModel(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) GetGroupList(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GroupList, error) {
	out := new(GroupList)
	err := c.cc.Invoke(ctx, "/protos.CoreService/GetGroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) AddGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetManufacturerList(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*ManufacturerList, error) {
	out := new(ManufacturerList)
	err := c.cc.Invoke(ctx, "/protos.CoreService/GetManufacturerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) AddManufacturer(ctx context.Context, in *Manufacturer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/AddManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateManufacturer(ctx context.Context, in *Manufacturer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/UpdateManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteManufacturer(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/DeleteManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetDeviceTypeList(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*DeviceTypeList, error) {
	out := new(DeviceTypeList)
	err := c.cc.Invoke(ctx, "/protos.CoreService/GetDeviceTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) AddDeviceType(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/AddDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateDeviceType(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/UpdateDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteDeviceType(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/DeleteDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetDeviceModelList(ctx context.Context, in *GetDeviceModelsRequest, opts ...grpc.CallOption) (*DeviceModelList, error) {
	out := new(DeviceModelList)
	err := c.cc.Invoke(ctx, "/protos.CoreService/GetDeviceModelList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) AddDeviceModel(ctx context.Context, in *DeviceModel, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/AddDeviceModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) UpdateDeviceModel(ctx context.Context, in *DeviceModel, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/UpdateDeviceModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteDeviceModel(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CoreService/DeleteDeviceModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	GetGroupList(context.Context, *GetGroupsRequest) (*GroupList, error)
	AddGroup(context.Context, *Group) (*Empty, error)
	UpdateGroup(context.Context, *Group) (*Empty, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*Empty, error)
	GetManufacturerList(context.Context, *GetManufacturersRequest) (*ManufacturerList, error)
	AddManufacturer(context.Context, *Manufacturer) (*Empty, error)
	UpdateManufacturer(context.Context, *Manufacturer) (*Empty, error)
	DeleteManufacturer(context.Context, *DeleteRequest) (*Empty, error)
	GetDeviceTypeList(context.Context, *GetDeviceTypesRequest) (*DeviceTypeList, error)
	AddDeviceType(context.Context, *DeviceType) (*Empty, error)
	UpdateDeviceType(context.Context, *DeviceType) (*Empty, error)
	DeleteDeviceType(context.Context, *DeleteRequest) (*Empty, error)
	GetDeviceModelList(context.Context, *GetDeviceModelsRequest) (*DeviceModelList, error)
	AddDeviceModel(context.Context, *DeviceModel) (*Empty, error)
	UpdateDeviceModel(context.Context, *DeviceModel) (*Empty, error)
	DeleteDeviceModel(context.Context, *DeleteRequest) (*Empty, error)
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_GetGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/GetGroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetGroupList(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).AddGroup(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateGroup(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetManufacturerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManufacturersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetManufacturerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/GetManufacturerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetManufacturerList(ctx, req.(*GetManufacturersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_AddManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manufacturer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).AddManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/AddManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).AddManufacturer(ctx, req.(*Manufacturer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manufacturer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/UpdateManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateManufacturer(ctx, req.(*Manufacturer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/DeleteManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteManufacturer(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetDeviceTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetDeviceTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/GetDeviceTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetDeviceTypeList(ctx, req.(*GetDeviceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_AddDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).AddDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/AddDeviceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).AddDeviceType(ctx, req.(*DeviceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/UpdateDeviceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateDeviceType(ctx, req.(*DeviceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/DeleteDeviceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteDeviceType(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetDeviceModelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetDeviceModelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/GetDeviceModelList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetDeviceModelList(ctx, req.(*GetDeviceModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_AddDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).AddDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/AddDeviceModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).AddDeviceModel(ctx, req.(*DeviceModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_UpdateDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).UpdateDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/UpdateDeviceModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).UpdateDeviceModel(ctx, req.(*DeviceModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CoreService/DeleteDeviceModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteDeviceModel(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupList",
			Handler:    _CoreService_GetGroupList_Handler,
		},
		{
			MethodName: "AddGroup",
			Handler:    _CoreService_AddGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _CoreService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _CoreService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetManufacturerList",
			Handler:    _CoreService_GetManufacturerList_Handler,
		},
		{
			MethodName: "AddManufacturer",
			Handler:    _CoreService_AddManufacturer_Handler,
		},
		{
			MethodName: "UpdateManufacturer",
			Handler:    _CoreService_UpdateManufacturer_Handler,
		},
		{
			MethodName: "DeleteManufacturer",
			Handler:    _CoreService_DeleteManufacturer_Handler,
		},
		{
			MethodName: "GetDeviceTypeList",
			Handler:    _CoreService_GetDeviceTypeList_Handler,
		},
		{
			MethodName: "AddDeviceType",
			Handler:    _CoreService_AddDeviceType_Handler,
		},
		{
			MethodName: "UpdateDeviceType",
			Handler:    _CoreService_UpdateDeviceType_Handler,
		},
		{
			MethodName: "DeleteDeviceType",
			Handler:    _CoreService_DeleteDeviceType_Handler,
		},
		{
			MethodName: "GetDeviceModelList",
			Handler:    _CoreService_GetDeviceModelList_Handler,
		},
		{
			MethodName: "AddDeviceModel",
			Handler:    _CoreService_AddDeviceModel_Handler,
		},
		{
			MethodName: "UpdateDeviceModel",
			Handler:    _CoreService_UpdateDeviceModel_Handler,
		},
		{
			MethodName: "DeleteDeviceModel",
			Handler:    _CoreService_DeleteDeviceModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core",
}

func init() { proto.RegisterFile("core", fileDescriptor_core_f781c88c7b89c5af) }

var fileDescriptor_core_f781c88c7b89c5af = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x5f, 0x74, 0xca, 0xed, 0xa6, 0xeb, 0x9d, 0x5f, 0x04, 0xd4, 0x57, 0x41, 0xd8, 0xc3,
	0xfc, 0xda, 0xf4, 0x69, 0x38, 0x29, 0x88, 0x03, 0xf1, 0xe3, 0x59, 0x66, 0x73, 0x95, 0xc1, 0xba,
	0xd4, 0x34, 0x15, 0xf6, 0x0f, 0xfa, 0x77, 0xc9, 0x12, 0x9b, 0xa6, 0xdd, 0x84, 0xee, 0x31, 0xf7,
	0x9c, 0x93, 0x9c, 0xdf, 0xdd, 0x28, 0xac, 0x85, 0x42, 0x12, 0xd6, 0x62, 0x29, 0x94, 0x48, 0x58,
	0x2d, 0x14, 0x51, 0x24, 0xa6, 0x6c, 0xfd, 0x53, 0x8a, 0x34, 0x66, 0xf5, 0x68, 0x34, 0x4d, 0x3f,
	0x46, 0xa1, 0x4a, 0x25, 0x49, 0xe6, 0x71, 0xfa, 0x1e, 0x87, 0xf4, 0xa6, 0x66, 0x31, 0xb1, 0xfa,
	0xdf, 0x21, 0x12, 0x9c, 0x26, 0x9d, 0x9f, 0x0d, 0xf0, 0x6e, 0x85, 0xa4, 0x67, 0x92, 0xf3, 0x29,
	0xde, 0x40, 0x3d, 0x20, 0x15, 0xcc, 0x2f, 0x79, 0x18, 0x27, 0x0a, 0x0f, 0xda, 0xe6, 0x81, 0x76,
	0x36, 0x4d, 0x9e, 0xe8, 0x2b, 0xa5, 0x44, 0x31, 0xdf, 0x2a, 0xd6, 0x7c, 0x02, 0x9b, 0x7d, 0xce,
	0xf5, 0x19, 0x1b, 0x05, 0x99, 0xd9, 0xe3, 0x5d, 0x14, 0xab, 0x19, 0x9e, 0x82, 0xf7, 0x1a, 0xf3,
	0x91, 0xa2, 0x2a, 0xe6, 0x2e, 0x78, 0x03, 0x9a, 0x50, 0x66, 0x66, 0x99, 0xea, 0x0c, 0xb3, 0x52,
	0xa5, 0xe4, 0x23, 0xb4, 0x02, 0x52, 0x43, 0x67, 0x17, 0xba, 0xe7, 0xb1, 0x03, 0xe5, 0x8a, 0x96,
	0xcd, 0x52, 0x2f, 0x44, 0x2f, 0x61, 0xbb, 0xcf, 0xb9, 0x3b, 0xc6, 0x9d, 0x65, 0xe6, 0x72, 0x93,
	0x1e, 0xa0, 0x01, 0x5e, 0x3d, 0x7a, 0x0d, 0x68, 0x48, 0x0b, 0xd1, 0xdd, 0xe2, 0x16, 0xfe, 0x59,
	0xc0, 0x3d, 0xf8, 0x01, 0xa9, 0x81, 0xfe, 0xc5, 0x5f, 0x66, 0x31, 0x69, 0x86, 0x43, 0x07, 0x3f,
	0x97, 0x2c, 0xfc, 0x5e, 0x7e, 0x73, 0x21, 0xd6, 0x81, 0x46, 0x9f, 0xf3, 0x7c, 0x88, 0xb8, 0x68,
	0x2c, 0xbf, 0x7f, 0x01, 0x4d, 0x83, 0xbd, 0x5a, 0xac, 0x0b, 0x4d, 0x83, 0xe5, 0xc4, 0xaa, 0x01,
	0x0f, 0x01, 0x2d, 0xd5, 0x70, 0xfe, 0x0f, 0xd7, 0xd5, 0x8f, 0x16, 0x88, 0xb5, 0x66, 0x91, 0xf7,
	0x8b, 0x95, 0xf2, 0xe0, 0x39, 0x6c, 0x59, 0x66, 0x3d, 0xc5, 0xd6, 0x12, 0x6b, 0xb9, 0xc4, 0x15,
	0xf8, 0x2e, 0x75, 0xf5, 0x60, 0x0f, 0x7c, 0x97, 0xdb, 0x04, 0x2b, 0x81, 0xbf, 0x9b, 0x0f, 0xc1,
	0xd9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xa6, 0x09, 0x6d, 0x17, 0x04, 0x00, 0x00,
}