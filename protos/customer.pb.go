// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer

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

type Customer struct {
	ID                   uint64     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Passwd1              string     `protobuf:"bytes,3,opt,name=Passwd1,proto3" json:"Passwd1,omitempty"`
	Passwd2              string     `protobuf:"bytes,4,opt,name=Passwd2,proto3" json:"Passwd2,omitempty"`
	Title                string     `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Address              string     `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	Path                 string     `protobuf:"bytes,7,opt,name=Path,proto3" json:"Path,omitempty"`
	Enable               bool       `protobuf:"varint,8,opt,name=Enable,proto3" json:"Enable,omitempty"`
	Contacts             []*Contact `protobuf:"bytes,20,rep,name=Contacts,proto3" json:"Contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_04198dc38a516628, []int{0}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetPasswd1() string {
	if m != nil {
		return m.Passwd1
	}
	return ""
}

func (m *Customer) GetPasswd2() string {
	if m != nil {
		return m.Passwd2
	}
	return ""
}

func (m *Customer) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Customer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Customer) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Customer) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Customer) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type CustomerList struct {
	List                 []*Customer `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64      `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CustomerList) Reset()         { *m = CustomerList{} }
func (m *CustomerList) String() string { return proto.CompactTextString(m) }
func (*CustomerList) ProtoMessage()    {}
func (*CustomerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_04198dc38a516628, []int{1}
}
func (m *CustomerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerList.Unmarshal(m, b)
}
func (m *CustomerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerList.Marshal(b, m, deterministic)
}
func (dst *CustomerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerList.Merge(dst, src)
}
func (m *CustomerList) XXX_Size() int {
	return xxx_messageInfo_CustomerList.Size(m)
}
func (m *CustomerList) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerList.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerList proto.InternalMessageInfo

func (m *CustomerList) GetList() []*Customer {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *CustomerList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetCustomerRequest struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_04198dc38a516628, []int{2}
}
func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(dst, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetCustomersRequest struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCustomersRequest) Reset()         { *m = GetCustomersRequest{} }
func (m *GetCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomersRequest) ProtoMessage()    {}
func (*GetCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_04198dc38a516628, []int{3}
}
func (m *GetCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomersRequest.Unmarshal(m, b)
}
func (m *GetCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomersRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomersRequest.Merge(dst, src)
}
func (m *GetCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomersRequest.Size(m)
}
func (m *GetCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomersRequest proto.InternalMessageInfo

func (m *GetCustomersRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "protos.Customer")
	proto.RegisterType((*CustomerList)(nil), "protos.CustomerList")
	proto.RegisterType((*GetCustomerRequest)(nil), "protos.GetCustomerRequest")
	proto.RegisterType((*GetCustomersRequest)(nil), "protos.GetCustomersRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	Get(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	GetList(ctx context.Context, in *GetCustomersRequest, opts ...grpc.CallOption) (*CustomerList, error)
	Add(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) Get(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/protos.CustomerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetList(ctx context.Context, in *GetCustomersRequest, opts ...grpc.CallOption) (*CustomerList, error) {
	out := new(CustomerList)
	err := c.cc.Invoke(ctx, "/protos.CustomerService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Add(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Update(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Delete(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	Get(context.Context, *GetCustomerRequest) (*Customer, error)
	GetList(context.Context, *GetCustomersRequest) (*CustomerList, error)
	Add(context.Context, *Customer) (*Empty, error)
	Update(context.Context, *Customer) (*Empty, error)
	Delete(context.Context, *Customer) (*Empty, error)
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Get(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetList(ctx, req.(*GetCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Add(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Update(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Delete(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CustomerService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _CustomerService_GetList_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _CustomerService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CustomerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer",
}

func init() { proto.RegisterFile("customer", fileDescriptor_customer_04198dc38a516628) }

var fileDescriptor_customer_04198dc38a516628 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xe2, 0x40,
	0x14, 0xc6, 0xc9, 0x1f, 0x93, 0xec, 0x71, 0x77, 0x5d, 0x66, 0x65, 0x19, 0xb2, 0x37, 0x21, 0xc8,
	0x92, 0xa5, 0x20, 0x34, 0xde, 0xf6, 0x46, 0x54, 0xc4, 0x52, 0x8a, 0xa4, 0xed, 0x03, 0x8c, 0xc9,
	0xd0, 0x06, 0x4c, 0xc6, 0x66, 0x8e, 0x2d, 0xbd, 0xed, 0xeb, 0xf6, 0x25, 0x4a, 0x26, 0x19, 0x15,
	0xed, 0x85, 0x57, 0x99, 0xef, 0x9c, 0xdf, 0x37, 0xe7, 0x7c, 0x43, 0xc0, 0x4b, 0xb7, 0x12, 0x45,
	0xc1, 0x2b, 0xe2, 0x6c, 0x2a, 0x81, 0x42, 0xfa, 0x4e, 0x2a, 0x8a, 0x42, 0x94, 0xbe, 0x9b, 0x8a,
	0x12, 0x59, 0x8a, 0xe1, 0x87, 0x01, 0xde, 0x44, 0x53, 0x3f, 0xc1, 0x5c, 0x4c, 0xa9, 0x11, 0x18,
	0x91, 0x9d, 0x98, 0x8b, 0x29, 0x21, 0x60, 0xdf, 0xb2, 0x82, 0x53, 0x33, 0x30, 0xa2, 0x6f, 0x89,
	0x3a, 0x13, 0x0a, 0xee, 0x92, 0x49, 0xf9, 0x9a, 0x5d, 0x52, 0x4b, 0x95, 0xb5, 0xdc, 0x77, 0x62,
	0x6a, 0x1f, 0x76, 0x62, 0xd2, 0x87, 0xce, 0x7d, 0x8e, 0x6b, 0x4e, 0x3b, 0xaa, 0xde, 0x88, 0x9a,
	0x1f, 0x67, 0x59, 0xc5, 0xa5, 0xa4, 0x4e, 0xc3, 0xb7, 0xb2, 0x9e, 0xbb, 0x64, 0xf8, 0x44, 0xdd,
	0x66, 0x6e, 0x7d, 0x26, 0x7f, 0xc0, 0x99, 0x95, 0x6c, 0xb5, 0xe6, 0xd4, 0x0b, 0x8c, 0xc8, 0x4b,
	0x5a, 0x45, 0x2e, 0xc0, 0x9b, 0x34, 0x59, 0x24, 0xed, 0x07, 0x56, 0xd4, 0x8d, 0x7b, 0xc3, 0x26,
	0xec, 0xb0, 0xad, 0x27, 0x3b, 0x20, 0xbc, 0x86, 0xef, 0x3a, 0xec, 0x4d, 0x2e, 0x91, 0x0c, 0xc0,
	0xae, 0xbf, 0xd4, 0x50, 0xc6, 0x5f, 0x3b, 0x63, 0xcb, 0x24, 0xaa, 0x5b, 0xaf, 0x3f, 0x11, 0xdb,
	0x12, 0xd5, 0x3b, 0xd8, 0x49, 0x23, 0xc2, 0x01, 0x90, 0x39, 0xc7, 0x1d, 0xca, 0x9f, 0xb7, 0x5c,
	0xe2, 0xf1, 0x13, 0x86, 0x0b, 0xf8, 0x7d, 0x40, 0x49, 0x8d, 0xc5, 0x00, 0x4b, 0xf6, 0x98, 0x97,
	0x0c, 0x73, 0x51, 0x2a, 0xbc, 0x1b, 0x13, 0x3d, 0x7e, 0xdf, 0x49, 0x0e, 0xa8, 0xf8, 0xdd, 0x84,
	0x9e, 0xbe, 0xe8, 0x8e, 0x57, 0x2f, 0x79, 0xca, 0xc9, 0x08, 0xac, 0x39, 0x47, 0xe2, 0x6b, 0xeb,
	0xe9, 0x46, 0xfe, 0x49, 0x2a, 0x72, 0x05, 0xee, 0x9c, 0xa3, 0x8a, 0xf6, 0xf7, 0x0b, 0xa3, 0x5e,
	0xd2, 0xef, 0x1f, 0x3b, 0x95, 0xe5, 0x1f, 0x58, 0xe3, 0x2c, 0x23, 0x27, 0xd7, 0xfa, 0x3f, 0x74,
	0x65, 0x56, 0x6c, 0xf0, 0x8d, 0xfc, 0x07, 0xe7, 0x61, 0x93, 0x31, 0xe4, 0x67, 0xa1, 0x53, 0xbe,
	0xe6, 0x67, 0xa0, 0xab, 0xe6, 0x47, 0x1e, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x93, 0xcf,
	0xd0, 0xdb, 0x02, 0x00, 0x00,
}
