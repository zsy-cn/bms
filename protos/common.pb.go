// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b39a7e17f4513f99, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Pagination struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize             uint64   `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	SortBy               string   `protobuf:"bytes,3,opt,name=SortBy,proto3" json:"SortBy,omitempty"`
	Order                bool     `protobuf:"varint,4,opt,name=Order,proto3" json:"Order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b39a7e17f4513f99, []int{1}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pagination) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *Pagination) GetOrder() bool {
	if m != nil {
		return m.Order
	}
	return false
}

// 通用delete请求
type DeleteRequest struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b39a7e17f4513f99, []int{2}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type DeviceCountInfo struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Alive                uint64   `protobuf:"varint,2,opt,name=Alive,proto3" json:"Alive,omitempty"`
	Alarm                uint64   `protobuf:"varint,3,opt,name=Alarm,proto3" json:"Alarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceCountInfo) Reset()         { *m = DeviceCountInfo{} }
func (m *DeviceCountInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceCountInfo) ProtoMessage()    {}
func (*DeviceCountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b39a7e17f4513f99, []int{3}
}
func (m *DeviceCountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceCountInfo.Unmarshal(m, b)
}
func (m *DeviceCountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceCountInfo.Marshal(b, m, deterministic)
}
func (dst *DeviceCountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCountInfo.Merge(dst, src)
}
func (m *DeviceCountInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceCountInfo.Size(m)
}
func (m *DeviceCountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCountInfo proto.InternalMessageInfo

func (m *DeviceCountInfo) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DeviceCountInfo) GetAlive() uint64 {
	if m != nil {
		return m.Alive
	}
	return 0
}

func (m *DeviceCountInfo) GetAlarm() uint64 {
	if m != nil {
		return m.Alarm
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "protos.Empty")
	proto.RegisterType((*Pagination)(nil), "protos.Pagination")
	proto.RegisterType((*DeleteRequest)(nil), "protos.DeleteRequest")
	proto.RegisterType((*DeviceCountInfo)(nil), "protos.DeviceCountInfo")
}

func init() { proto.RegisterFile("common", fileDescriptor_common_b39a7e17f4513f99) }

var fileDescriptor_common_b39a7e17f4513f99 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xc9, 0xba, 0x5d, 0xeb, 0x80, 0x0a, 0x41, 0x24, 0x78, 0x71, 0xc9, 0x69, 0xef, 0x3e,
	0x41, 0x35, 0x1e, 0xf6, 0x64, 0x49, 0xf1, 0x01, 0x62, 0x1d, 0x4b, 0xa4, 0xc9, 0xd4, 0x34, 0x5d,
	0x58, 0x9f, 0x5e, 0xf2, 0xc7, 0x9e, 0x66, 0x7e, 0x1f, 0x03, 0xbf, 0x6f, 0xa0, 0xdb, 0x92, 0x73,
	0xe4, 0x79, 0x77, 0x08, 0x14, 0xe9, 0x28, 0x2f, 0x61, 0xf1, 0xea, 0x0e, 0x71, 0x96, 0xdf, 0x00,
	0x6b, 0xb3, 0xb3, 0xde, 0x44, 0x4b, 0x9e, 0x73, 0x68, 0xd7, 0x66, 0x87, 0x82, 0xf5, 0x6c, 0x68,
	0x75, 0xde, 0xf9, 0x03, 0x2c, 0xd3, 0xdc, 0xd8, 0x5f, 0x14, 0x4d, 0xce, 0xcf, 0xcc, 0xef, 0xa1,
	0xdb, 0x50, 0x88, 0xcf, 0xb3, 0xb8, 0xe8, 0xd9, 0x70, 0xa5, 0x2b, 0xf1, 0x3b, 0x58, 0xbc, 0x85,
	0x4f, 0x0c, 0xa2, 0xed, 0xd9, 0xb0, 0xd4, 0x05, 0xe4, 0x23, 0x5c, 0x2b, 0xdc, 0x63, 0x44, 0x8d,
	0x3f, 0x27, 0x3c, 0x46, 0x7e, 0x03, 0xcd, 0xa8, 0xaa, 0xac, 0x19, 0x95, 0x7c, 0x87, 0x5b, 0x85,
	0x93, 0xdd, 0xe2, 0x0b, 0x9d, 0x7c, 0x1c, 0xfd, 0x17, 0x25, 0xc3, 0xca, 0x25, 0xaa, 0x67, 0x95,
	0x92, 0x61, 0xb5, 0xb7, 0xd3, 0x7f, 0xa5, 0x02, 0x25, 0x35, 0xc1, 0xe5, 0x3a, 0x39, 0x35, 0xc1,
	0x7d, 0x94, 0xa7, 0x9f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xb8, 0xb6, 0x70, 0x05, 0x01,
	0x00, 0x00,
}