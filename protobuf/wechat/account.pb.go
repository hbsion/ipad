// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package wechat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetSafetyInfoRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,req,name=BaseRequest" json:"BaseRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetSafetyInfoRequest) Reset()         { *m = GetSafetyInfoRequest{} }
func (m *GetSafetyInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetSafetyInfoRequest) ProtoMessage()    {}
func (*GetSafetyInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *GetSafetyInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSafetyInfoRequest.Unmarshal(m, b)
}
func (m *GetSafetyInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSafetyInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetSafetyInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSafetyInfoRequest.Merge(m, src)
}
func (m *GetSafetyInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetSafetyInfoRequest.Size(m)
}
func (m *GetSafetyInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSafetyInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSafetyInfoRequest proto.InternalMessageInfo

func (m *GetSafetyInfoRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

type GetSafetyInfoResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,req,name=BaseResponse" json:"BaseResponse,omitempty"`
	ContactUsernameList  *DeviceInfo   `protobuf:"bytes,2,req,name=ContactUsernameList" json:"ContactUsernameList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetSafetyInfoResponse) Reset()         { *m = GetSafetyInfoResponse{} }
func (m *GetSafetyInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetSafetyInfoResponse) ProtoMessage()    {}
func (*GetSafetyInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *GetSafetyInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSafetyInfoResponse.Unmarshal(m, b)
}
func (m *GetSafetyInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSafetyInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetSafetyInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSafetyInfoResponse.Merge(m, src)
}
func (m *GetSafetyInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetSafetyInfoResponse.Size(m)
}
func (m *GetSafetyInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSafetyInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSafetyInfoResponse proto.InternalMessageInfo

func (m *GetSafetyInfoResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *GetSafetyInfoResponse) GetContactUsernameList() *DeviceInfo {
	if m != nil {
		return m.ContactUsernameList
	}
	return nil
}

type DeviceInfo struct {
	SafeDevices          []*SafeDevice `protobuf:"bytes,1,rep,name=SafeDevices" json:"SafeDevices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceInfo.Unmarshal(m, b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceInfo.Size(m)
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetSafeDevices() []*SafeDevice {
	if m != nil {
		return m.SafeDevices
	}
	return nil
}

type SafeDevice struct {
	Uuid                 *string  `protobuf:"bytes,1,opt,name=QrUuid" json:"QrUuid,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	DeviceType           *string  `protobuf:"bytes,3,opt,name=DeviceType" json:"DeviceType,omitempty"`
	CreateTime           *uint32  `protobuf:"varint,4,req,name=CreateTime" json:"CreateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SafeDevice) Reset()         { *m = SafeDevice{} }
func (m *SafeDevice) String() string { return proto.CompactTextString(m) }
func (*SafeDevice) ProtoMessage()    {}
func (*SafeDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *SafeDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SafeDevice.Unmarshal(m, b)
}
func (m *SafeDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SafeDevice.Marshal(b, m, deterministic)
}
func (m *SafeDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SafeDevice.Merge(m, src)
}
func (m *SafeDevice) XXX_Size() int {
	return xxx_messageInfo_SafeDevice.Size(m)
}
func (m *SafeDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_SafeDevice.DiscardUnknown(m)
}

var xxx_messageInfo_SafeDevice proto.InternalMessageInfo

func (m *SafeDevice) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SafeDevice) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *SafeDevice) GetDeviceType() string {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return ""
}

func (m *SafeDevice) GetCreateTime() uint32 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

type DelSafeDeviceRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,req,name=BaseRequest" json:"BaseRequest,omitempty"`
	Uuid                 *string      `protobuf:"bytes,2,opt,name=QrUuid" json:"QrUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DelSafeDeviceRequest) Reset()         { *m = DelSafeDeviceRequest{} }
func (m *DelSafeDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DelSafeDeviceRequest) ProtoMessage()    {}
func (*DelSafeDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *DelSafeDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSafeDeviceRequest.Unmarshal(m, b)
}
func (m *DelSafeDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSafeDeviceRequest.Marshal(b, m, deterministic)
}
func (m *DelSafeDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSafeDeviceRequest.Merge(m, src)
}
func (m *DelSafeDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DelSafeDeviceRequest.Size(m)
}
func (m *DelSafeDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSafeDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelSafeDeviceRequest proto.InternalMessageInfo

func (m *DelSafeDeviceRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *DelSafeDeviceRequest) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

type DelSafeDeviceResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,req,name=BaseResponse" json:"BaseResponse,omitempty"`
	SafeDevice           *uint32       `protobuf:"varint,2,opt,name=SafeDevice" json:"SafeDevice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DelSafeDeviceResponse) Reset()         { *m = DelSafeDeviceResponse{} }
func (m *DelSafeDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*DelSafeDeviceResponse) ProtoMessage()    {}
func (*DelSafeDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *DelSafeDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSafeDeviceResponse.Unmarshal(m, b)
}
func (m *DelSafeDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSafeDeviceResponse.Marshal(b, m, deterministic)
}
func (m *DelSafeDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSafeDeviceResponse.Merge(m, src)
}
func (m *DelSafeDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_DelSafeDeviceResponse.Size(m)
}
func (m *DelSafeDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSafeDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelSafeDeviceResponse proto.InternalMessageInfo

func (m *DelSafeDeviceResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *DelSafeDeviceResponse) GetSafeDevice() uint32 {
	if m != nil && m.SafeDevice != nil {
		return *m.SafeDevice
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSafetyInfoRequest)(nil), "wechat_proto.GetSafetyInfoRequest")
	proto.RegisterType((*GetSafetyInfoResponse)(nil), "wechat_proto.GetSafetyInfoResponse")
	proto.RegisterType((*DeviceInfo)(nil), "wechat_proto.DeviceInfo")
	proto.RegisterType((*SafeDevice)(nil), "wechat_proto.SafeDevice")
	proto.RegisterType((*DelSafeDeviceRequest)(nil), "wechat_proto.DelSafeDeviceRequest")
	proto.RegisterType((*DelSafeDeviceResponse)(nil), "wechat_proto.DelSafeDeviceResponse")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x95, 0xb4, 0x4b, 0x6f, 0x9a, 0xc5, 0x5f, 0x2b, 0xf9, 0xeb, 0x50, 0x45, 0x9e, 0x32,
	0x05, 0x89, 0x11, 0x24, 0x86, 0xb6, 0x12, 0x7f, 0x84, 0x18, 0xdc, 0x76, 0x61, 0x41, 0x56, 0xb8,
	0x85, 0x48, 0x24, 0x0e, 0xb1, 0x43, 0xd5, 0xe7, 0xe1, 0x45, 0x51, 0xec, 0xa0, 0xb8, 0x25, 0x63,
	0x37, 0x9f, 0x73, 0x4f, 0x7e, 0xe7, 0xc6, 0x86, 0x50, 0xa4, 0xa9, 0xac, 0x0b, 0x9d, 0x94, 0x95,
	0xd4, 0x92, 0x8c, 0xf7, 0x98, 0xbe, 0x0b, 0xfd, 0x62, 0xd4, 0xac, 0x55, 0x76, 0xc6, 0xd6, 0x30,
	0xb9, 0x45, 0xbd, 0x16, 0x3b, 0xd4, 0x87, 0xfb, 0x62, 0x27, 0x39, 0x7e, 0xd6, 0xa8, 0x34, 0xb9,
	0x86, 0x60, 0x21, 0x14, 0xb6, 0x92, 0x7a, 0x91, 0x1f, 0x07, 0x97, 0xff, 0x13, 0x97, 0x94, 0x38,
	0x01, 0xee, 0xa6, 0xd9, 0xb7, 0x07, 0xd3, 0x13, 0xaa, 0x2a, 0x65, 0xa1, 0x90, 0xdc, 0xc0, 0xd8,
	0x06, 0xad, 0x6e, 0xb9, 0xb3, 0x3e, 0xae, 0x4d, 0xf0, 0xa3, 0x3c, 0x79, 0x80, 0x7f, 0x4b, 0x59,
	0x68, 0x91, 0xea, 0xad, 0xc2, 0xaa, 0x10, 0x39, 0x3e, 0x66, 0x4a, 0x53, 0xdf, 0x60, 0xe8, 0x31,
	0x66, 0x85, 0x5f, 0x59, 0x8a, 0xa6, 0xbe, 0xef, 0x23, 0x76, 0x07, 0xd0, 0x45, 0xc8, 0x15, 0x04,
	0xcd, 0xbe, 0xd6, 0x51, 0xd4, 0x8b, 0x06, 0x7f, 0x89, 0x5d, 0x80, 0xbb, 0x61, 0xa6, 0x01, 0x3a,
	0x49, 0x08, 0x0c, 0x9f, 0x44, 0xde, 0xfc, 0x9b, 0x17, 0x8f, 0xb8, 0x39, 0x37, 0xde, 0xb6, 0xce,
	0x5e, 0xa9, 0x6f, 0xbd, 0xe6, 0x4c, 0xe6, 0xbf, 0xfd, 0x9b, 0x43, 0x89, 0x74, 0x60, 0x26, 0x8e,
	0xd3, 0xcc, 0x97, 0x15, 0x0a, 0x8d, 0x9b, 0x2c, 0x47, 0x3a, 0x8c, 0xfc, 0x38, 0xe4, 0x8e, 0xc3,
	0xde, 0x60, 0xb2, 0xc2, 0x0f, 0x67, 0xa7, 0x33, 0x3c, 0x5d, 0xdf, 0xa2, 0x6c, 0x0f, 0xd3, 0x93,
	0xa2, 0x33, 0xbd, 0xe6, 0xdc, 0xbd, 0x37, 0x53, 0x19, 0x72, 0xc7, 0x59, 0x04, 0xcf, 0xa3, 0x24,
	0xb9, 0xb0, 0xb4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xae, 0x59, 0x66, 0xd5, 0x02, 0x00,
	0x00,
}
