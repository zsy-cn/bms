// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wifi

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

type WifiDeviceGroup struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CustomerID           uint64   `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	DeviceTypeID         uint64   `protobuf:"varint,4,opt,name=DeviceTypeID,proto3" json:"DeviceTypeID,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	DeviceTotal          uint64   `protobuf:"varint,11,opt,name=DeviceTotal,proto3" json:"DeviceTotal,omitempty"`
	DeviceOn             uint64   `protobuf:"varint,12,opt,name=DeviceOn,proto3" json:"DeviceOn,omitempty"`
	DeviceOff            uint64   `protobuf:"varint,13,opt,name=DeviceOff,proto3" json:"DeviceOff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiDeviceGroup) Reset()         { *m = WifiDeviceGroup{} }
func (m *WifiDeviceGroup) String() string { return proto.CompactTextString(m) }
func (*WifiDeviceGroup) ProtoMessage()    {}
func (*WifiDeviceGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_9d234d4328aff37c, []int{0}
}
func (m *WifiDeviceGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiDeviceGroup.Unmarshal(m, b)
}
func (m *WifiDeviceGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiDeviceGroup.Marshal(b, m, deterministic)
}
func (dst *WifiDeviceGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiDeviceGroup.Merge(dst, src)
}
func (m *WifiDeviceGroup) XXX_Size() int {
	return xxx_messageInfo_WifiDeviceGroup.Size(m)
}
func (m *WifiDeviceGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiDeviceGroup.DiscardUnknown(m)
}

var xxx_messageInfo_WifiDeviceGroup proto.InternalMessageInfo

func (m *WifiDeviceGroup) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *WifiDeviceGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WifiDeviceGroup) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *WifiDeviceGroup) GetDeviceTypeID() uint64 {
	if m != nil {
		return m.DeviceTypeID
	}
	return 0
}

func (m *WifiDeviceGroup) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WifiDeviceGroup) GetDeviceTotal() uint64 {
	if m != nil {
		return m.DeviceTotal
	}
	return 0
}

func (m *WifiDeviceGroup) GetDeviceOn() uint64 {
	if m != nil {
		return m.DeviceOn
	}
	return 0
}

func (m *WifiDeviceGroup) GetDeviceOff() uint64 {
	if m != nil {
		return m.DeviceOff
	}
	return 0
}

type WifiDeviceGroupList struct {
	List                 []*WifiDeviceGroup `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64             `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64             `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64             `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64             `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WifiDeviceGroupList) Reset()         { *m = WifiDeviceGroupList{} }
func (m *WifiDeviceGroupList) String() string { return proto.CompactTextString(m) }
func (*WifiDeviceGroupList) ProtoMessage()    {}
func (*WifiDeviceGroupList) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_9d234d4328aff37c, []int{1}
}
func (m *WifiDeviceGroupList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiDeviceGroupList.Unmarshal(m, b)
}
func (m *WifiDeviceGroupList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiDeviceGroupList.Marshal(b, m, deterministic)
}
func (dst *WifiDeviceGroupList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiDeviceGroupList.Merge(dst, src)
}
func (m *WifiDeviceGroupList) XXX_Size() int {
	return xxx_messageInfo_WifiDeviceGroupList.Size(m)
}
func (m *WifiDeviceGroupList) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiDeviceGroupList.DiscardUnknown(m)
}

var xxx_messageInfo_WifiDeviceGroupList proto.InternalMessageInfo

func (m *WifiDeviceGroupList) GetList() []*WifiDeviceGroup {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *WifiDeviceGroupList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *WifiDeviceGroupList) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *WifiDeviceGroupList) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *WifiDeviceGroupList) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type WifiDevice struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	SerialNumber         string   `protobuf:"bytes,3,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Position             string   `protobuf:"bytes,5,opt,name=Position,proto3" json:"Position,omitempty"`
	Group                string   `protobuf:"bytes,11,opt,name=Group,proto3" json:"Group,omitempty"`
	GroupID              uint64   `protobuf:"varint,12,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	DeviceType           string   `protobuf:"bytes,13,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`
	DeviceTypeID         uint64   `protobuf:"varint,14,opt,name=DeviceTypeID,proto3" json:"DeviceTypeID,omitempty"`
	DeviceModel          string   `protobuf:"bytes,15,opt,name=DeviceModel,proto3" json:"DeviceModel,omitempty"`
	DeviceModelID        uint64   `protobuf:"varint,16,opt,name=DeviceModelID,proto3" json:"DeviceModelID,omitempty"`
	Customer             string   `protobuf:"bytes,17,opt,name=Customer,proto3" json:"Customer,omitempty"`
	CustomerID           uint64   `protobuf:"varint,18,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Latitude             float64  `protobuf:"fixed64,21,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,22,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	StatusCode           uint64   `protobuf:"varint,23,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Status               string   `protobuf:"bytes,24,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,25,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Connections          uint64   `protobuf:"varint,31,opt,name=Connections,proto3" json:"Connections,omitempty"`
	UpSpeed              string   `protobuf:"bytes,32,opt,name=UpSpeed,proto3" json:"UpSpeed,omitempty"`
	DownSpeed            string   `protobuf:"bytes,33,opt,name=DownSpeed,proto3" json:"DownSpeed,omitempty"`
	DataTraffic          string   `protobuf:"bytes,34,opt,name=DataTraffic,proto3" json:"DataTraffic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiDevice) Reset()         { *m = WifiDevice{} }
func (m *WifiDevice) String() string { return proto.CompactTextString(m) }
func (*WifiDevice) ProtoMessage()    {}
func (*WifiDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_9d234d4328aff37c, []int{2}
}
func (m *WifiDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiDevice.Unmarshal(m, b)
}
func (m *WifiDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiDevice.Marshal(b, m, deterministic)
}
func (dst *WifiDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiDevice.Merge(dst, src)
}
func (m *WifiDevice) XXX_Size() int {
	return xxx_messageInfo_WifiDevice.Size(m)
}
func (m *WifiDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiDevice.DiscardUnknown(m)
}

var xxx_messageInfo_WifiDevice proto.InternalMessageInfo

func (m *WifiDevice) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *WifiDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WifiDevice) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *WifiDevice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WifiDevice) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *WifiDevice) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *WifiDevice) GetGroupID() uint64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *WifiDevice) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *WifiDevice) GetDeviceTypeID() uint64 {
	if m != nil {
		return m.DeviceTypeID
	}
	return 0
}

func (m *WifiDevice) GetDeviceModel() string {
	if m != nil {
		return m.DeviceModel
	}
	return ""
}

func (m *WifiDevice) GetDeviceModelID() uint64 {
	if m != nil {
		return m.DeviceModelID
	}
	return 0
}

func (m *WifiDevice) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *WifiDevice) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *WifiDevice) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *WifiDevice) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *WifiDevice) GetStatusCode() uint64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *WifiDevice) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *WifiDevice) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WifiDevice) GetConnections() uint64 {
	if m != nil {
		return m.Connections
	}
	return 0
}

func (m *WifiDevice) GetUpSpeed() string {
	if m != nil {
		return m.UpSpeed
	}
	return ""
}

func (m *WifiDevice) GetDownSpeed() string {
	if m != nil {
		return m.DownSpeed
	}
	return ""
}

func (m *WifiDevice) GetDataTraffic() string {
	if m != nil {
		return m.DataTraffic
	}
	return ""
}

type WifiDeviceList struct {
	List                 []*WifiDevice `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64        `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64        `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64        `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64        `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WifiDeviceList) Reset()         { *m = WifiDeviceList{} }
func (m *WifiDeviceList) String() string { return proto.CompactTextString(m) }
func (*WifiDeviceList) ProtoMessage()    {}
func (*WifiDeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_9d234d4328aff37c, []int{3}
}
func (m *WifiDeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiDeviceList.Unmarshal(m, b)
}
func (m *WifiDeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiDeviceList.Marshal(b, m, deterministic)
}
func (dst *WifiDeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiDeviceList.Merge(dst, src)
}
func (m *WifiDeviceList) XXX_Size() int {
	return xxx_messageInfo_WifiDeviceList.Size(m)
}
func (m *WifiDeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiDeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_WifiDeviceList proto.InternalMessageInfo

func (m *WifiDeviceList) GetList() []*WifiDevice {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *WifiDeviceList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *WifiDeviceList) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *WifiDeviceList) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *WifiDeviceList) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*WifiDeviceGroup)(nil), "protos.WifiDeviceGroup")
	proto.RegisterType((*WifiDeviceGroupList)(nil), "protos.WifiDeviceGroupList")
	proto.RegisterType((*WifiDevice)(nil), "protos.WifiDevice")
	proto.RegisterType((*WifiDeviceList)(nil), "protos.WifiDeviceList")
}

func init() { proto.RegisterFile("wifi", fileDescriptor_wifi_9d234d4328aff37c) }

var fileDescriptor_wifi_9d234d4328aff37c = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0xa7, 0xd4, 0x93, 0x34, 0x85, 0x05, 0xda, 0xa1, 0x42, 0x60, 0x2c, 0x84, 0x22,
	0x21, 0xf5, 0x00, 0x5f, 0x80, 0x6c, 0x09, 0x59, 0x0a, 0x05, 0x39, 0x45, 0x9c, 0xb7, 0xc9, 0xba,
	0x5a, 0x29, 0xf1, 0x5a, 0xeb, 0x35, 0x15, 0x7c, 0x08, 0x5f, 0xc1, 0x99, 0x1f, 0xe3, 0x07, 0xd0,
	0xce, 0x3a, 0xb6, 0x93, 0xf4, 0xc0, 0xad, 0x27, 0xef, 0x7b, 0x33, 0x19, 0xcd, 0x9b, 0x79, 0x13,
	0xf0, 0x6f, 0x65, 0x2e, 0xd9, 0x61, 0xa9, 0x95, 0x51, 0x55, 0xf4, 0xd7, 0x83, 0x93, 0x6f, 0x32,
	0x97, 0x89, 0xf8, 0x2e, 0x17, 0xe2, 0xa3, 0x56, 0x75, 0xc9, 0x26, 0x30, 0x48, 0x13, 0xf4, 0x42,
	0x6f, 0xea, 0x67, 0x83, 0x34, 0x61, 0x0c, 0xfc, 0x4b, 0xbe, 0x16, 0x38, 0x08, 0xbd, 0x69, 0x90,
	0xd1, 0x9b, 0xbd, 0x00, 0x88, 0xeb, 0xca, 0xa8, 0xb5, 0xd0, 0x69, 0x82, 0x07, 0x94, 0xdb, 0x63,
	0x58, 0x04, 0x63, 0x57, 0xf2, 0xea, 0x47, 0x29, 0xd2, 0x04, 0x7d, 0xca, 0xd8, 0xe2, 0xd8, 0x73,
	0x08, 0x62, 0x2d, 0xb8, 0x11, 0xcb, 0x0f, 0x06, 0x87, 0x54, 0xbc, 0x23, 0x58, 0x08, 0xa3, 0x26,
	0x5b, 0x19, 0xbe, 0xc2, 0x11, 0x15, 0xe8, 0x53, 0xec, 0x1c, 0x8e, 0x1c, 0xfc, 0x5c, 0xe0, 0x98,
	0xc2, 0x2d, 0xb6, 0xb5, 0x9b, 0x77, 0x9e, 0xe3, 0x31, 0x05, 0x3b, 0x22, 0xfa, 0xe3, 0xc1, 0xe3,
	0x1d, 0xd5, 0x33, 0x59, 0x19, 0xf6, 0x16, 0x7c, 0xfb, 0x45, 0x2f, 0x3c, 0x98, 0x8e, 0xde, 0x9d,
	0x5d, 0xb8, 0x21, 0x5d, 0xec, 0xa4, 0x66, 0x94, 0xc4, 0x9e, 0xc0, 0x30, 0x56, 0x75, 0x61, 0x68,
	0x2e, 0x7e, 0xe6, 0x80, 0x6d, 0x3b, 0xae, 0xb5, 0x16, 0x85, 0xf9, 0xc2, 0x6f, 0x04, 0x82, 0x6b,
	0xbb, 0x47, 0xd9, 0xb6, 0xed, 0x77, 0x2e, 0x7f, 0x8a, 0x46, 0x55, 0x8b, 0xed, 0x58, 0x49, 0x9b,
	0x2b, 0xec, 0x44, 0xf5, 0x98, 0xe8, 0xd7, 0x10, 0xa0, 0xeb, 0xe6, 0xbf, 0x36, 0x15, 0xc1, 0x78,
	0x2e, 0xb4, 0xe4, 0xab, 0xcb, 0x7a, 0x7d, 0x2d, 0x34, 0xed, 0x2a, 0xc8, 0xb6, 0x38, 0x37, 0xeb,
	0x6a, 0xa1, 0x65, 0x69, 0xa4, 0x2a, 0x68, 0x59, 0x41, 0xd6, 0xa7, 0xa8, 0x69, 0x55, 0x49, 0x0a,
	0xbb, 0x55, 0xb5, 0xd8, 0x0e, 0x82, 0xe6, 0x42, 0x6a, 0x82, 0xcc, 0x01, 0x86, 0xf0, 0x80, 0x1e,
	0x69, 0xd2, 0xe8, 0xd8, 0x40, 0x2b, 0xb2, 0xf3, 0x01, 0x2d, 0x27, 0xc8, 0x7a, 0xcc, 0x9e, 0x77,
	0x26, 0x77, 0x78, 0xa7, 0x75, 0xc7, 0x27, 0xb5, 0x14, 0x2b, 0x3c, 0xd9, 0x74, 0xdc, 0x52, 0xec,
	0x35, 0x1c, 0xf7, 0x60, 0x9a, 0xe0, 0x43, 0x2a, 0xb3, 0x4d, 0x5a, 0x5d, 0x1b, 0xd7, 0xe2, 0x23,
	0xa7, 0x6b, 0x83, 0x77, 0x3c, 0xce, 0xf6, 0x3c, 0x7e, 0x0e, 0x47, 0x33, 0x6e, 0xa4, 0xa9, 0x97,
	0x02, 0x9f, 0x86, 0xde, 0xd4, 0xcb, 0x5a, 0x6c, 0xfd, 0x37, 0x53, 0xc5, 0x8d, 0x0b, 0x9e, 0x52,
	0xb0, 0x23, 0x6c, 0xe5, 0xb9, 0xe1, 0xa6, 0xae, 0x62, 0xb5, 0x14, 0x78, 0xe6, 0x2a, 0x77, 0x0c,
	0x3b, 0x85, 0x43, 0x87, 0x10, 0xa9, 0xa7, 0x06, 0x6d, 0x5f, 0xcc, 0xb3, 0x3b, 0x2e, 0x26, 0x56,
	0x45, 0x21, 0x16, 0x76, 0x2b, 0x15, 0xbe, 0x6c, 0xac, 0xd7, 0x51, 0x76, 0x27, 0x5f, 0xcb, 0x79,
	0x29, 0xc4, 0x12, 0x43, 0xfa, 0xf5, 0x06, 0xd2, 0xbd, 0xa8, 0xdb, 0xc2, 0xc5, 0x5e, 0xb9, 0xca,
	0x2d, 0x41, 0xd3, 0xe6, 0x86, 0x5f, 0x69, 0x9e, 0xe7, 0x72, 0x81, 0x51, 0x33, 0xed, 0x8e, 0x8a,
	0x7e, 0x7b, 0x30, 0xe9, 0x8c, 0x49, 0xf7, 0xf1, 0x66, 0xeb, 0x98, 0xd8, 0xfe, 0x31, 0xdd, 0xdf,
	0x1d, 0x5d, 0xbb, 0xbf, 0xbf, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x3a, 0xd0, 0x04,
	0x0d, 0x05, 0x00, 0x00,
}
