// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sound_box

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

type GetSoundBoxMediasRequest struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CustomerID           uint64      `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetSoundBoxMediasRequest) Reset()         { *m = GetSoundBoxMediasRequest{} }
func (m *GetSoundBoxMediasRequest) String() string { return proto.CompactTextString(m) }
func (*GetSoundBoxMediasRequest) ProtoMessage()    {}
func (*GetSoundBoxMediasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{0}
}
func (m *GetSoundBoxMediasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSoundBoxMediasRequest.Unmarshal(m, b)
}
func (m *GetSoundBoxMediasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSoundBoxMediasRequest.Marshal(b, m, deterministic)
}
func (dst *GetSoundBoxMediasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSoundBoxMediasRequest.Merge(dst, src)
}
func (m *GetSoundBoxMediasRequest) XXX_Size() int {
	return xxx_messageInfo_GetSoundBoxMediasRequest.Size(m)
}
func (m *GetSoundBoxMediasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSoundBoxMediasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSoundBoxMediasRequest proto.InternalMessageInfo

func (m *GetSoundBoxMediasRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetSoundBoxMediasRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSoundBoxMediasRequest) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

type SoundBoxMedia struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CustomerID           uint64   `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Duration             string   `protobuf:"bytes,4,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Size                 uint64   `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=Path,proto3" json:"Path,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoundBoxMedia) Reset()         { *m = SoundBoxMedia{} }
func (m *SoundBoxMedia) String() string { return proto.CompactTextString(m) }
func (*SoundBoxMedia) ProtoMessage()    {}
func (*SoundBoxMedia) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{1}
}
func (m *SoundBoxMedia) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoundBoxMedia.Unmarshal(m, b)
}
func (m *SoundBoxMedia) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoundBoxMedia.Marshal(b, m, deterministic)
}
func (dst *SoundBoxMedia) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoundBoxMedia.Merge(dst, src)
}
func (m *SoundBoxMedia) XXX_Size() int {
	return xxx_messageInfo_SoundBoxMedia.Size(m)
}
func (m *SoundBoxMedia) XXX_DiscardUnknown() {
	xxx_messageInfo_SoundBoxMedia.DiscardUnknown(m)
}

var xxx_messageInfo_SoundBoxMedia proto.InternalMessageInfo

func (m *SoundBoxMedia) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SoundBoxMedia) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SoundBoxMedia) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *SoundBoxMedia) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *SoundBoxMedia) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SoundBoxMedia) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SoundBoxMedia) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetSoundBoxMediasResponse struct {
	List                 []*SoundBoxMedia `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64           `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64           `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64           `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64           `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetSoundBoxMediasResponse) Reset()         { *m = GetSoundBoxMediasResponse{} }
func (m *GetSoundBoxMediasResponse) String() string { return proto.CompactTextString(m) }
func (*GetSoundBoxMediasResponse) ProtoMessage()    {}
func (*GetSoundBoxMediasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{2}
}
func (m *GetSoundBoxMediasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSoundBoxMediasResponse.Unmarshal(m, b)
}
func (m *GetSoundBoxMediasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSoundBoxMediasResponse.Marshal(b, m, deterministic)
}
func (dst *GetSoundBoxMediasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSoundBoxMediasResponse.Merge(dst, src)
}
func (m *GetSoundBoxMediasResponse) XXX_Size() int {
	return xxx_messageInfo_GetSoundBoxMediasResponse.Size(m)
}
func (m *GetSoundBoxMediasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSoundBoxMediasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSoundBoxMediasResponse proto.InternalMessageInfo

func (m *GetSoundBoxMediasResponse) GetList() []*SoundBoxMedia {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetSoundBoxMediasResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetSoundBoxMediasResponse) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *GetSoundBoxMediasResponse) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetSoundBoxMediasResponse) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type UpdateSoundBoxMediaRequest struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CustomerID           uint64   `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSoundBoxMediaRequest) Reset()         { *m = UpdateSoundBoxMediaRequest{} }
func (m *UpdateSoundBoxMediaRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSoundBoxMediaRequest) ProtoMessage()    {}
func (*UpdateSoundBoxMediaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{3}
}
func (m *UpdateSoundBoxMediaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSoundBoxMediaRequest.Unmarshal(m, b)
}
func (m *UpdateSoundBoxMediaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSoundBoxMediaRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateSoundBoxMediaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSoundBoxMediaRequest.Merge(dst, src)
}
func (m *UpdateSoundBoxMediaRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSoundBoxMediaRequest.Size(m)
}
func (m *UpdateSoundBoxMediaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSoundBoxMediaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSoundBoxMediaRequest proto.InternalMessageInfo

func (m *UpdateSoundBoxMediaRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UpdateSoundBoxMediaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSoundBoxMediaRequest) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

type GetSoundBoxDeviceGroupsRequest struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	CustomerID           uint64      `protobuf:"varint,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Name                 string      `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetSoundBoxDeviceGroupsRequest) Reset()         { *m = GetSoundBoxDeviceGroupsRequest{} }
func (m *GetSoundBoxDeviceGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSoundBoxDeviceGroupsRequest) ProtoMessage()    {}
func (*GetSoundBoxDeviceGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{4}
}
func (m *GetSoundBoxDeviceGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSoundBoxDeviceGroupsRequest.Unmarshal(m, b)
}
func (m *GetSoundBoxDeviceGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSoundBoxDeviceGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *GetSoundBoxDeviceGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSoundBoxDeviceGroupsRequest.Merge(dst, src)
}
func (m *GetSoundBoxDeviceGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSoundBoxDeviceGroupsRequest.Size(m)
}
func (m *GetSoundBoxDeviceGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSoundBoxDeviceGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSoundBoxDeviceGroupsRequest proto.InternalMessageInfo

func (m *GetSoundBoxDeviceGroupsRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetSoundBoxDeviceGroupsRequest) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *GetSoundBoxDeviceGroupsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SoundBoxDeviceGroup struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CustomerID           uint64   `protobuf:"varint,3,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	DeviceTotal          uint64   `protobuf:"varint,11,opt,name=DeviceTotal,proto3" json:"DeviceTotal,omitempty"`
	DeviceOn             uint64   `protobuf:"varint,12,opt,name=DeviceOn,proto3" json:"DeviceOn,omitempty"`
	DeviceOff            uint64   `protobuf:"varint,13,opt,name=DeviceOff,proto3" json:"DeviceOff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoundBoxDeviceGroup) Reset()         { *m = SoundBoxDeviceGroup{} }
func (m *SoundBoxDeviceGroup) String() string { return proto.CompactTextString(m) }
func (*SoundBoxDeviceGroup) ProtoMessage()    {}
func (*SoundBoxDeviceGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{5}
}
func (m *SoundBoxDeviceGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoundBoxDeviceGroup.Unmarshal(m, b)
}
func (m *SoundBoxDeviceGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoundBoxDeviceGroup.Marshal(b, m, deterministic)
}
func (dst *SoundBoxDeviceGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoundBoxDeviceGroup.Merge(dst, src)
}
func (m *SoundBoxDeviceGroup) XXX_Size() int {
	return xxx_messageInfo_SoundBoxDeviceGroup.Size(m)
}
func (m *SoundBoxDeviceGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SoundBoxDeviceGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SoundBoxDeviceGroup proto.InternalMessageInfo

func (m *SoundBoxDeviceGroup) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SoundBoxDeviceGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SoundBoxDeviceGroup) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *SoundBoxDeviceGroup) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SoundBoxDeviceGroup) GetDeviceTotal() uint64 {
	if m != nil {
		return m.DeviceTotal
	}
	return 0
}

func (m *SoundBoxDeviceGroup) GetDeviceOn() uint64 {
	if m != nil {
		return m.DeviceOn
	}
	return 0
}

func (m *SoundBoxDeviceGroup) GetDeviceOff() uint64 {
	if m != nil {
		return m.DeviceOff
	}
	return 0
}

type SoundBoxDeviceGroupList struct {
	List                 []*SoundBoxDeviceGroup `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64                 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64                 `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64                 `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64                 `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoundBoxDeviceGroupList) Reset()         { *m = SoundBoxDeviceGroupList{} }
func (m *SoundBoxDeviceGroupList) String() string { return proto.CompactTextString(m) }
func (*SoundBoxDeviceGroupList) ProtoMessage()    {}
func (*SoundBoxDeviceGroupList) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{6}
}
func (m *SoundBoxDeviceGroupList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoundBoxDeviceGroupList.Unmarshal(m, b)
}
func (m *SoundBoxDeviceGroupList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoundBoxDeviceGroupList.Marshal(b, m, deterministic)
}
func (dst *SoundBoxDeviceGroupList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoundBoxDeviceGroupList.Merge(dst, src)
}
func (m *SoundBoxDeviceGroupList) XXX_Size() int {
	return xxx_messageInfo_SoundBoxDeviceGroupList.Size(m)
}
func (m *SoundBoxDeviceGroupList) XXX_DiscardUnknown() {
	xxx_messageInfo_SoundBoxDeviceGroupList.DiscardUnknown(m)
}

var xxx_messageInfo_SoundBoxDeviceGroupList proto.InternalMessageInfo

func (m *SoundBoxDeviceGroupList) GetList() []*SoundBoxDeviceGroup {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SoundBoxDeviceGroupList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoundBoxDeviceGroupList) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *SoundBoxDeviceGroupList) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SoundBoxDeviceGroupList) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type SoundBoxDevice struct {
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
	Volume               float64  `protobuf:"fixed64,30,opt,name=Volume,proto3" json:"Volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoundBoxDevice) Reset()         { *m = SoundBoxDevice{} }
func (m *SoundBoxDevice) String() string { return proto.CompactTextString(m) }
func (*SoundBoxDevice) ProtoMessage()    {}
func (*SoundBoxDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{7}
}
func (m *SoundBoxDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoundBoxDevice.Unmarshal(m, b)
}
func (m *SoundBoxDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoundBoxDevice.Marshal(b, m, deterministic)
}
func (dst *SoundBoxDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoundBoxDevice.Merge(dst, src)
}
func (m *SoundBoxDevice) XXX_Size() int {
	return xxx_messageInfo_SoundBoxDevice.Size(m)
}
func (m *SoundBoxDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_SoundBoxDevice.DiscardUnknown(m)
}

var xxx_messageInfo_SoundBoxDevice proto.InternalMessageInfo

func (m *SoundBoxDevice) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SoundBoxDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SoundBoxDevice) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *SoundBoxDevice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SoundBoxDevice) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *SoundBoxDevice) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *SoundBoxDevice) GetGroupID() uint64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *SoundBoxDevice) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *SoundBoxDevice) GetDeviceTypeID() uint64 {
	if m != nil {
		return m.DeviceTypeID
	}
	return 0
}

func (m *SoundBoxDevice) GetDeviceModel() string {
	if m != nil {
		return m.DeviceModel
	}
	return ""
}

func (m *SoundBoxDevice) GetDeviceModelID() uint64 {
	if m != nil {
		return m.DeviceModelID
	}
	return 0
}

func (m *SoundBoxDevice) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *SoundBoxDevice) GetCustomerID() uint64 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *SoundBoxDevice) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *SoundBoxDevice) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *SoundBoxDevice) GetStatusCode() uint64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *SoundBoxDevice) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SoundBoxDevice) GetActived() string {
	if m != nil {
		return m.Actived
	}
	return ""
}

func (m *SoundBoxDevice) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *SoundBoxDevice) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type SoundBoxDeviceList struct {
	List                 []*SoundBoxDevice `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                uint64            `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	CurrentPage          uint64            `protobuf:"varint,10,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	PageSize             uint64            `protobuf:"varint,11,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	TotalCount           uint64            `protobuf:"varint,12,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoundBoxDeviceList) Reset()         { *m = SoundBoxDeviceList{} }
func (m *SoundBoxDeviceList) String() string { return proto.CompactTextString(m) }
func (*SoundBoxDeviceList) ProtoMessage()    {}
func (*SoundBoxDeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_sound_box_443c03cbcacacafe, []int{8}
}
func (m *SoundBoxDeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoundBoxDeviceList.Unmarshal(m, b)
}
func (m *SoundBoxDeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoundBoxDeviceList.Marshal(b, m, deterministic)
}
func (dst *SoundBoxDeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoundBoxDeviceList.Merge(dst, src)
}
func (m *SoundBoxDeviceList) XXX_Size() int {
	return xxx_messageInfo_SoundBoxDeviceList.Size(m)
}
func (m *SoundBoxDeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_SoundBoxDeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_SoundBoxDeviceList proto.InternalMessageInfo

func (m *SoundBoxDeviceList) GetList() []*SoundBoxDevice {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SoundBoxDeviceList) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoundBoxDeviceList) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *SoundBoxDeviceList) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SoundBoxDeviceList) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSoundBoxMediasRequest)(nil), "protos.GetSoundBoxMediasRequest")
	proto.RegisterType((*SoundBoxMedia)(nil), "protos.SoundBoxMedia")
	proto.RegisterType((*GetSoundBoxMediasResponse)(nil), "protos.GetSoundBoxMediasResponse")
	proto.RegisterType((*UpdateSoundBoxMediaRequest)(nil), "protos.UpdateSoundBoxMediaRequest")
	proto.RegisterType((*GetSoundBoxDeviceGroupsRequest)(nil), "protos.GetSoundBoxDeviceGroupsRequest")
	proto.RegisterType((*SoundBoxDeviceGroup)(nil), "protos.SoundBoxDeviceGroup")
	proto.RegisterType((*SoundBoxDeviceGroupList)(nil), "protos.SoundBoxDeviceGroupList")
	proto.RegisterType((*SoundBoxDevice)(nil), "protos.SoundBoxDevice")
	proto.RegisterType((*SoundBoxDeviceList)(nil), "protos.SoundBoxDeviceList")
}

func init() { proto.RegisterFile("sound_box", fileDescriptor_sound_box_443c03cbcacacafe) }

var fileDescriptor_sound_box_443c03cbcacacafe = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x96, 0xb7, 0xe9, 0x96, 0x9d, 0x6d, 0x0b, 0x18, 0x5a, 0xdc, 0x05, 0x95, 0x55, 0xc4, 0x61,
	0xe1, 0x50, 0xa4, 0xf2, 0x04, 0x65, 0x23, 0x55, 0x91, 0xda, 0x52, 0x65, 0x81, 0x2b, 0xb8, 0x1b,
	0xb7, 0x44, 0xda, 0xc4, 0x21, 0x71, 0xaa, 0xc2, 0x91, 0x13, 0x0f, 0x04, 0x17, 0x1e, 0x80, 0x1b,
	0x12, 0x8f, 0x84, 0x3c, 0x76, 0x36, 0x71, 0x59, 0x24, 0x04, 0x95, 0x7a, 0x8a, 0xbf, 0x6f, 0xec,
	0xf1, 0xfc, 0x7c, 0xe3, 0x40, 0xaf, 0x94, 0x55, 0x16, 0xbf, 0x39, 0x91, 0x17, 0xb4, 0x9b, 0x17,
	0x52, 0xc9, 0x72, 0xd0, 0x9d, 0xca, 0x34, 0x95, 0x99, 0xff, 0x89, 0x00, 0xdb, 0x17, 0x6a, 0xa2,
	0x37, 0x3c, 0x97, 0x17, 0x87, 0x22, 0x4e, 0x78, 0x19, 0x89, 0xf7, 0x95, 0x28, 0x15, 0xdd, 0x05,
	0x38, 0xe6, 0x67, 0x49, 0xc6, 0x55, 0x22, 0x33, 0x46, 0x86, 0x64, 0xd4, 0xdf, 0xa5, 0x3b, 0xc6,
	0xc3, 0x4e, 0x63, 0x89, 0x5a, 0xbb, 0x28, 0x05, 0xef, 0x88, 0xa7, 0x82, 0x75, 0x86, 0x64, 0xd4,
	0x8b, 0x70, 0x4d, 0xb7, 0x01, 0xc6, 0x55, 0xa9, 0x64, 0x2a, 0x8a, 0x30, 0x60, 0x4b, 0x43, 0x32,
	0xf2, 0xa2, 0x16, 0xe3, 0x7f, 0x25, 0xb0, 0xe6, 0x44, 0x40, 0xd7, 0xa1, 0x13, 0x06, 0x78, 0xa3,
	0x17, 0x75, 0xc2, 0xe0, 0x5f, 0xbc, 0xd2, 0x01, 0xdc, 0x08, 0xaa, 0xc2, 0xc4, 0xee, 0xe1, 0xb9,
	0x39, 0xd6, 0xfe, 0x26, 0xc9, 0x47, 0xc1, 0x96, 0xf1, 0x14, 0xae, 0x35, 0x77, 0xcc, 0xd5, 0x3b,
	0xd6, 0x35, 0x77, 0xe8, 0x35, 0x7d, 0x00, 0xbd, 0x71, 0x21, 0xb8, 0x12, 0xf1, 0x9e, 0x62, 0x7d,
	0x34, 0x34, 0x84, 0xff, 0x8d, 0xc0, 0xd6, 0x82, 0xe2, 0x95, 0xb9, 0xcc, 0x4a, 0x41, 0x1f, 0x83,
	0x77, 0x90, 0x94, 0x8a, 0x91, 0xe1, 0xd2, 0xa8, 0xbf, 0xbb, 0x51, 0xd7, 0xcd, 0xd9, 0x1d, 0xe1,
	0x16, 0x7a, 0x17, 0x96, 0xc7, 0xb2, 0xca, 0x14, 0xe6, 0xe7, 0x45, 0x06, 0xd0, 0x21, 0xf4, 0xc7,
	0x55, 0x51, 0x88, 0x4c, 0x1d, 0xf3, 0x33, 0xc1, 0x00, 0x6d, 0x6d, 0x4a, 0xa7, 0xa8, 0xbf, 0x98,
	0x4a, 0x1f, 0xcd, 0x73, 0xac, 0xcb, 0xf3, 0x52, 0x2a, 0x3e, 0x33, 0x8e, 0x57, 0x4d, 0x79, 0x1a,
	0xc6, 0x7f, 0x0b, 0x83, 0x57, 0x79, 0xcc, 0x95, 0x70, 0x03, 0xb2, 0xad, 0xbf, 0x82, 0x06, 0xf8,
	0x9f, 0x09, 0x6c, 0xb7, 0xca, 0x13, 0x88, 0xf3, 0x64, 0x2a, 0xf6, 0x0b, 0x59, 0xe5, 0xff, 0xa5,
	0x30, 0xf7, 0xda, 0xce, 0x6f, 0x7d, 0xaf, 0x43, 0x5d, 0x6a, 0x42, 0xf5, 0x7f, 0x10, 0xb8, 0xb3,
	0x20, 0x8e, 0x2b, 0xd1, 0xd9, 0x26, 0x74, 0x27, 0x8a, 0xab, 0xaa, 0xb4, 0x2a, 0xb3, 0x48, 0xb7,
	0xcf, 0x5c, 0x85, 0x45, 0xb7, 0xfd, 0x69, 0x53, 0xa8, 0x50, 0x84, 0x2f, 0x32, 0xdb, 0xa0, 0x39,
	0xd6, 0xca, 0xb3, 0xeb, 0xd3, 0x53, 0xb6, 0x86, 0xc6, 0x86, 0xf0, 0xbf, 0x13, 0x78, 0xf8, 0xc7,
	0xd2, 0x5a, 0xfd, 0x3d, 0x75, 0xf4, 0x77, 0xff, 0xb2, 0xfe, 0x5a, 0x67, 0xae, 0x51, 0x85, 0x3f,
	0x3d, 0x58, 0x77, 0x23, 0xfa, 0xab, 0x9e, 0xf8, 0xb0, 0x3a, 0x11, 0x45, 0xc2, 0x67, 0x47, 0x55,
	0x7a, 0x22, 0x0a, 0xdb, 0x6b, 0x87, 0x33, 0xf5, 0x2f, 0xa7, 0x45, 0x92, 0xb7, 0x9e, 0x80, 0x36,
	0x85, 0x81, 0xcb, 0x32, 0x41, 0xf3, 0xb2, 0x79, 0x21, 0x6a, 0xac, 0x8b, 0x81, 0xb5, 0xb1, 0x53,
	0x6f, 0x00, 0x65, 0xb0, 0x82, 0x8b, 0x30, 0xb0, 0xb9, 0xd4, 0x50, 0x27, 0x6a, 0x5b, 0xfb, 0x21,
	0x17, 0xd8, 0xb0, 0x5e, 0xd4, 0x62, 0x74, 0xc4, 0x0d, 0x0a, 0x03, 0xb6, 0x8e, 0xc7, 0x1d, 0xae,
	0x51, 0xcc, 0xa1, 0x8c, 0xc5, 0x8c, 0xdd, 0xac, 0x23, 0x9e, 0x53, 0xf4, 0x11, 0xac, 0xb5, 0x60,
	0x18, 0xb0, 0x5b, 0xe8, 0xc6, 0x25, 0x75, 0x5e, 0xb5, 0x3e, 0xd9, 0x6d, 0x93, 0x57, 0x8d, 0x2f,
	0xa9, 0x99, 0x2e, 0x7a, 0x35, 0x0f, 0xb8, 0x4a, 0x54, 0x15, 0x0b, 0xb6, 0x31, 0x24, 0x23, 0x12,
	0xcd, 0xb1, 0xd6, 0xe4, 0x81, 0xcc, 0xce, 0x8c, 0x71, 0x13, 0x8d, 0x0d, 0xa1, 0x3d, 0x1b, 0xe5,
	0x8f, 0x65, 0x2c, 0xd8, 0x3d, 0xe3, 0xb9, 0x61, 0x5a, 0x73, 0xc2, 0x9c, 0x39, 0x61, 0xb0, 0xb2,
	0x37, 0x55, 0xc9, 0xb9, 0x88, 0xd9, 0x16, 0x1a, 0x6a, 0xe8, 0xbe, 0xbe, 0x83, 0x4b, 0xaf, 0xaf,
	0xf6, 0xf7, 0x5a, 0xce, 0xaa, 0x54, 0xb0, 0x6d, 0x0c, 0xc5, 0x22, 0xff, 0x0b, 0x01, 0xea, 0x4a,
	0x0a, 0xd5, 0xfd, 0xc4, 0x19, 0x87, 0xcd, 0xc5, 0xe3, 0x70, 0x7d, 0x93, 0x70, 0x62, 0xfe, 0xcc,
	0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x79, 0x73, 0x24, 0x8a, 0xad, 0x07, 0x00, 0x00,
}
