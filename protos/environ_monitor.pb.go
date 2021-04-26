// Code generated by protoc-gen-go. DO NOT EDIT.
// source: environ_monitor

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

// 环境数据分段平均值
type EnvironMonitorSectionAverageData struct {
	Section              string   `protobuf:"bytes,1,opt,name=Section,proto3" json:"Section,omitempty"`
	Temperature          float64  `protobuf:"fixed64,2,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
	PM025                float64  `protobuf:"fixed64,3,opt,name=PM025,proto3" json:"PM025,omitempty"`
	Noise                float64  `protobuf:"fixed64,4,opt,name=Noise,proto3" json:"Noise,omitempty"`
	Humidity             float64  `protobuf:"fixed64,5,opt,name=Humidity,proto3" json:"Humidity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironMonitorSectionAverageData) Reset()         { *m = EnvironMonitorSectionAverageData{} }
func (m *EnvironMonitorSectionAverageData) String() string { return proto.CompactTextString(m) }
func (*EnvironMonitorSectionAverageData) ProtoMessage()    {}
func (*EnvironMonitorSectionAverageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_environ_monitor_92b514bee594ab3f, []int{0}
}
func (m *EnvironMonitorSectionAverageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironMonitorSectionAverageData.Unmarshal(m, b)
}
func (m *EnvironMonitorSectionAverageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironMonitorSectionAverageData.Marshal(b, m, deterministic)
}
func (dst *EnvironMonitorSectionAverageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironMonitorSectionAverageData.Merge(dst, src)
}
func (m *EnvironMonitorSectionAverageData) XXX_Size() int {
	return xxx_messageInfo_EnvironMonitorSectionAverageData.Size(m)
}
func (m *EnvironMonitorSectionAverageData) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironMonitorSectionAverageData.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironMonitorSectionAverageData proto.InternalMessageInfo

func (m *EnvironMonitorSectionAverageData) GetSection() string {
	if m != nil {
		return m.Section
	}
	return ""
}

func (m *EnvironMonitorSectionAverageData) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *EnvironMonitorSectionAverageData) GetPM025() float64 {
	if m != nil {
		return m.PM025
	}
	return 0
}

func (m *EnvironMonitorSectionAverageData) GetNoise() float64 {
	if m != nil {
		return m.Noise
	}
	return 0
}

func (m *EnvironMonitorSectionAverageData) GetHumidity() float64 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

type EnvironMonitorDevice struct {
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
	Actived              string   `protobuf:"bytes,25,opt,name=Actived,proto3" json:"Actived,omitempty"`
	CreatedAt            string   `protobuf:"bytes,26,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Temperature          float64  `protobuf:"fixed64,30,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
	PM025                float64  `protobuf:"fixed64,31,opt,name=PM025,proto3" json:"PM025,omitempty"`
	Noise                float64  `protobuf:"fixed64,32,opt,name=Noise,proto3" json:"Noise,omitempty"`
	Humidity             float64  `protobuf:"fixed64,33,opt,name=Humidity,proto3" json:"Humidity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironMonitorDevice) Reset()         { *m = EnvironMonitorDevice{} }
func (m *EnvironMonitorDevice) String() string { return proto.CompactTextString(m) }
func (*EnvironMonitorDevice) ProtoMessage()    {}
func (*EnvironMonitorDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_environ_monitor_92b514bee594ab3f, []int{1}
}
func (m *EnvironMonitorDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironMonitorDevice.Unmarshal(m, b)
}
func (m *EnvironMonitorDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironMonitorDevice.Marshal(b, m, deterministic)
}
func (dst *EnvironMonitorDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironMonitorDevice.Merge(dst, src)
}
func (m *EnvironMonitorDevice) XXX_Size() int {
	return xxx_messageInfo_EnvironMonitorDevice.Size(m)
}
func (m *EnvironMonitorDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironMonitorDevice.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironMonitorDevice proto.InternalMessageInfo

func (m *EnvironMonitorDevice) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *EnvironMonitorDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvironMonitorDevice) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *EnvironMonitorDevice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EnvironMonitorDevice) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *EnvironMonitorDevice) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *EnvironMonitorDevice) GetGroupID() uint64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *EnvironMonitorDevice) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *EnvironMonitorDevice) GetDeviceTypeID() uint64 {
	if m != nil {
		return m.DeviceTypeID
	}
	return 0
}

func (m *EnvironMonitorDevice) GetDeviceModel() string {
	if m != nil {
		return m.DeviceModel
	}
	return ""
}

func (m *EnvironMonitorDevice) GetDeviceModelID() uint64 {
	if m != nil {
		return m.DeviceModelID
	}
	return 0
}

func (m *EnvironMonitorDevice) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *EnvironMonitorDevice) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *EnvironMonitorDevice) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *EnvironMonitorDevice) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *EnvironMonitorDevice) GetStatusCode() uint64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *EnvironMonitorDevice) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EnvironMonitorDevice) GetActived() string {
	if m != nil {
		return m.Actived
	}
	return ""
}

func (m *EnvironMonitorDevice) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *EnvironMonitorDevice) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *EnvironMonitorDevice) GetPM025() float64 {
	if m != nil {
		return m.PM025
	}
	return 0
}

func (m *EnvironMonitorDevice) GetNoise() float64 {
	if m != nil {
		return m.Noise
	}
	return 0
}

func (m *EnvironMonitorDevice) GetHumidity() float64 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

type EnvironMonitorDeviceList struct {
	List                 []*EnvironMonitorDevice `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64                  `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64                  `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64                  `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64                  `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EnvironMonitorDeviceList) Reset()         { *m = EnvironMonitorDeviceList{} }
func (m *EnvironMonitorDeviceList) String() string { return proto.CompactTextString(m) }
func (*EnvironMonitorDeviceList) ProtoMessage()    {}
func (*EnvironMonitorDeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_environ_monitor_92b514bee594ab3f, []int{2}
}
func (m *EnvironMonitorDeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironMonitorDeviceList.Unmarshal(m, b)
}
func (m *EnvironMonitorDeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironMonitorDeviceList.Marshal(b, m, deterministic)
}
func (dst *EnvironMonitorDeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironMonitorDeviceList.Merge(dst, src)
}
func (m *EnvironMonitorDeviceList) XXX_Size() int {
	return xxx_messageInfo_EnvironMonitorDeviceList.Size(m)
}
func (m *EnvironMonitorDeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironMonitorDeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironMonitorDeviceList proto.InternalMessageInfo

func (m *EnvironMonitorDeviceList) GetList() []*EnvironMonitorDevice {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *EnvironMonitorDeviceList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *EnvironMonitorDeviceList) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *EnvironMonitorDeviceList) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *EnvironMonitorDeviceList) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*EnvironMonitorSectionAverageData)(nil), "protos.EnvironMonitorSectionAverageData")
	proto.RegisterType((*EnvironMonitorDevice)(nil), "protos.EnvironMonitorDevice")
	proto.RegisterType((*EnvironMonitorDeviceList)(nil), "protos.EnvironMonitorDeviceList")
}

func init() { proto.RegisterFile("environ_monitor", fileDescriptor_environ_monitor_92b514bee594ab3f) }

var fileDescriptor_environ_monitor_92b514bee594ab3f = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xef, 0x6e, 0xda, 0x3c,
	0x14, 0xc6, 0x15, 0x1a, 0xfa, 0xbe, 0x98, 0xfe, 0xd9, 0xac, 0xae, 0x3b, 0xab, 0xaa, 0x2e, 0x43,
	0xfb, 0xc0, 0xa7, 0xaa, 0xea, 0xb4, 0x0b, 0x40, 0x64, 0xda, 0x22, 0x01, 0x42, 0x81, 0xef, 0x93,
	0x4b, 0x8e, 0x90, 0x25, 0x12, 0x23, 0xc7, 0x41, 0xea, 0x2e, 0x68, 0xda, 0x65, 0xec, 0xd2, 0x26,
	0x9f, 0x43, 0x20, 0x54, 0x6c, 0x9f, 0xea, 0xdf, 0xf3, 0xd8, 0xee, 0x73, 0xcc, 0x03, 0xe2, 0x12,
	0x8b, 0x8d, 0xb6, 0xa6, 0xf8, 0x9e, 0x9b, 0x42, 0x3b, 0x63, 0xe5, 0xe9, 0xda, 0x1a, 0x67, 0xca,
	0xde, 0xcf, 0x40, 0x44, 0x5f, 0xd8, 0x1b, 0xb3, 0x35, 0xc3, 0x85, 0xd3, 0xa6, 0x18, 0x6c, 0xd0,
	0xaa, 0x25, 0xc6, 0xca, 0x29, 0x09, 0xe2, 0xbf, 0xad, 0x0a, 0x41, 0x14, 0xf4, 0x3b, 0x69, 0x8d,
	0x32, 0x12, 0xdd, 0x39, 0xe6, 0x6b, 0xb4, 0xca, 0x55, 0x16, 0xa1, 0x15, 0x05, 0xfd, 0x20, 0x6d,
	0x4a, 0xf2, 0x4a, 0xb4, 0xa7, 0xe3, 0x87, 0xc7, 0xcf, 0x70, 0x42, 0x1e, 0x83, 0x57, 0x27, 0x46,
	0x97, 0x08, 0x21, 0xab, 0x04, 0xf2, 0x46, 0xfc, 0xff, 0xad, 0xca, 0x75, 0xa6, 0xdd, 0x33, 0xb4,
	0xc9, 0xd8, 0x71, 0xef, 0x57, 0x5b, 0x5c, 0x1d, 0x06, 0x8d, 0x71, 0xa3, 0x17, 0x28, 0x2f, 0x44,
	0x2b, 0x89, 0x29, 0x57, 0x98, 0xb6, 0x92, 0x58, 0x4a, 0x11, 0x4e, 0x54, 0xce, 0x59, 0x3a, 0x29,
	0xad, 0x65, 0x4f, 0x9c, 0xcd, 0xd0, 0x6a, 0xb5, 0x9a, 0x54, 0xf9, 0x13, 0x5a, 0xca, 0xd2, 0x49,
	0x0f, 0x34, 0x3f, 0x4a, 0x8c, 0xe5, 0xc2, 0xea, 0x35, 0x0d, 0x1a, 0xd2, 0x96, 0xa6, 0xe4, 0xe3,
	0x4d, 0x4d, 0xa9, 0xc9, 0x6e, 0x93, 0xbd, 0x63, 0x3f, 0xd0, 0x57, 0x6b, 0xaa, 0x35, 0x74, 0xc9,
	0x60, 0xf0, 0x0f, 0x47, 0x8b, 0x24, 0x86, 0x33, 0x0a, 0x58, 0xa3, 0xbc, 0x13, 0x82, 0xf3, 0xcf,
	0x9f, 0xd7, 0x08, 0xe7, 0x74, 0xa8, 0xa1, 0xf8, 0xc4, 0x7b, 0x4a, 0x62, 0xb8, 0xa0, 0xe3, 0x07,
	0x1a, 0x27, 0xf6, 0x3c, 0x36, 0x19, 0xae, 0xe0, 0xb2, 0x4e, 0xbc, 0x93, 0xe4, 0x47, 0x71, 0xde,
	0xc0, 0x24, 0x86, 0x57, 0x74, 0xcd, 0xa1, 0xe8, 0xe7, 0x1a, 0x56, 0xa5, 0x33, 0x39, 0x5a, 0x78,
	0xcd, 0x73, 0xd5, 0xec, 0x73, 0xd6, 0xeb, 0x24, 0x06, 0x49, 0xc7, 0x1b, 0x8a, 0x3f, 0x3b, 0x52,
	0x4e, 0xbb, 0x2a, 0x43, 0x78, 0xc3, 0x1f, 0x59, 0xcd, 0xf2, 0x56, 0x74, 0x46, 0xa6, 0x58, 0xb2,
	0x79, 0x4d, 0xe6, 0x5e, 0xf0, 0x37, 0xcf, 0x9c, 0x72, 0x55, 0x39, 0x34, 0x19, 0xc2, 0x5b, 0xbe,
	0x79, 0xaf, 0xc8, 0x6b, 0x71, 0xca, 0x04, 0x40, 0x99, 0xb6, 0xe4, 0xdf, 0x74, 0xb0, 0x70, 0x7a,
	0x83, 0x19, 0xbc, 0xe3, 0x32, 0x6e, 0xd1, 0xff, 0xbf, 0xa1, 0x45, 0xe5, 0x30, 0x1b, 0x38, 0xb8,
	0x21, 0x6f, 0x2f, 0xbc, 0xac, 0xea, 0xdd, 0x3f, 0xaa, 0xfa, 0xfe, 0x68, 0x55, 0xa3, 0xbf, 0x55,
	0xf5, 0xc3, 0x8b, 0xaa, 0xfe, 0x0e, 0x04, 0x1c, 0xab, 0xea, 0x48, 0x97, 0x4e, 0x3e, 0x88, 0xd0,
	0xff, 0x85, 0x20, 0x3a, 0xe9, 0x77, 0x1f, 0x6f, 0xef, 0xf9, 0x7b, 0x78, 0x7f, 0x6c, 0x7f, 0x4a,
	0x3b, 0x7d, 0x80, 0xa1, 0xa9, 0x0a, 0x47, 0x8d, 0x0e, 0x53, 0x06, 0x3f, 0xce, 0xb0, 0xb2, 0x16,
	0x0b, 0x37, 0x55, 0x4b, 0x04, 0x41, 0x5e, 0x53, 0xa2, 0xba, 0xaa, 0x25, 0xce, 0xf4, 0x0f, 0xa4,
	0x56, 0x86, 0xe9, 0x8e, 0xfd, 0xe3, 0xcf, 0x8d, 0x53, 0x2b, 0xbe, 0x98, 0xbb, 0xd9, 0x50, 0x9e,
	0xf8, 0xe7, 0xe1, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x29, 0xe0, 0x14, 0x38, 0x04,
	0x00, 0x00,
}
