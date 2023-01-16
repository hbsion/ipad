// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trustInfo.proto

package mmproto

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

type TrustDeviceInfo struct {
	Key                  *string  `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Val                  *string  `protobuf:"bytes,2,req,name=val" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrustDeviceInfo) Reset()         { *m = TrustDeviceInfo{} }
func (m *TrustDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*TrustDeviceInfo) ProtoMessage()    {}
func (*TrustDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{0}
}

func (m *TrustDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustDeviceInfo.Unmarshal(m, b)
}
func (m *TrustDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustDeviceInfo.Marshal(b, m, deterministic)
}
func (m *TrustDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustDeviceInfo.Merge(m, src)
}
func (m *TrustDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_TrustDeviceInfo.Size(m)
}
func (m *TrustDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TrustDeviceInfo proto.InternalMessageInfo

func (m *TrustDeviceInfo) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *TrustDeviceInfo) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

type TrustData struct {
	Tdi                  []*TrustDeviceInfo `protobuf:"bytes,1,rep,name=tdi" json:"tdi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrustData) Reset()         { *m = TrustData{} }
func (m *TrustData) String() string { return proto.CompactTextString(m) }
func (*TrustData) ProtoMessage()    {}
func (*TrustData) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{1}
}

func (m *TrustData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustData.Unmarshal(m, b)
}
func (m *TrustData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustData.Marshal(b, m, deterministic)
}
func (m *TrustData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustData.Merge(m, src)
}
func (m *TrustData) XXX_Size() int {
	return xxx_messageInfo_TrustData.Size(m)
}
func (m *TrustData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustData.DiscardUnknown(m)
}

var xxx_messageInfo_TrustData proto.InternalMessageInfo

func (m *TrustData) GetTdi() []*TrustDeviceInfo {
	if m != nil {
		return m.Tdi
	}
	return nil
}

type TrustReq struct {
	Td                   *TrustData `protobuf:"bytes,1,req,name=td" json:"td,omitempty"`
	Md                   *string    `protobuf:"bytes,2,opt,name=md" json:"md,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TrustReq) Reset()         { *m = TrustReq{} }
func (m *TrustReq) String() string { return proto.CompactTextString(m) }
func (*TrustReq) ProtoMessage()    {}
func (*TrustReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{2}
}

func (m *TrustReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustReq.Unmarshal(m, b)
}
func (m *TrustReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustReq.Marshal(b, m, deterministic)
}
func (m *TrustReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustReq.Merge(m, src)
}
func (m *TrustReq) XXX_Size() int {
	return xxx_messageInfo_TrustReq.Size(m)
}
func (m *TrustReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustReq.DiscardUnknown(m)
}

var xxx_messageInfo_TrustReq proto.InternalMessageInfo

func (m *TrustReq) GetTd() *TrustData {
	if m != nil {
		return m.Td
	}
	return nil
}

func (m *TrustReq) GetMd() string {
	if m != nil && m.Md != nil {
		return *m.Md
	}
	return ""
}

type TrustSoftData struct {
	SoftConfig           []byte   `protobuf:"bytes,1,req,name=softConfig" json:"softConfig,omitempty"`
	SoftData             []byte   `protobuf:"bytes,2,req,name=softData" json:"softData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrustSoftData) Reset()         { *m = TrustSoftData{} }
func (m *TrustSoftData) String() string { return proto.CompactTextString(m) }
func (*TrustSoftData) ProtoMessage()    {}
func (*TrustSoftData) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{3}
}

func (m *TrustSoftData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustSoftData.Unmarshal(m, b)
}
func (m *TrustSoftData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustSoftData.Marshal(b, m, deterministic)
}
func (m *TrustSoftData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustSoftData.Merge(m, src)
}
func (m *TrustSoftData) XXX_Size() int {
	return xxx_messageInfo_TrustSoftData.Size(m)
}
func (m *TrustSoftData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustSoftData.DiscardUnknown(m)
}

var xxx_messageInfo_TrustSoftData proto.InternalMessageInfo

func (m *TrustSoftData) GetSoftConfig() []byte {
	if m != nil {
		return m.SoftConfig
	}
	return nil
}

func (m *TrustSoftData) GetSoftData() []byte {
	if m != nil {
		return m.SoftData
	}
	return nil
}

type TrustResponseData struct {
	SoftData             *TrustSoftData `protobuf:"bytes,2,req,name=softData" json:"softData,omitempty"`
	DeviceToken          *string        `protobuf:"bytes,3,req,name=deviceToken" json:"deviceToken,omitempty"`
	TimeStamp            *uint32        `protobuf:"varint,4,req,name=timeStamp" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TrustResponseData) Reset()         { *m = TrustResponseData{} }
func (m *TrustResponseData) String() string { return proto.CompactTextString(m) }
func (*TrustResponseData) ProtoMessage()    {}
func (*TrustResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{4}
}

func (m *TrustResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustResponseData.Unmarshal(m, b)
}
func (m *TrustResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustResponseData.Marshal(b, m, deterministic)
}
func (m *TrustResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustResponseData.Merge(m, src)
}
func (m *TrustResponseData) XXX_Size() int {
	return xxx_messageInfo_TrustResponseData.Size(m)
}
func (m *TrustResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_TrustResponseData proto.InternalMessageInfo

func (m *TrustResponseData) GetSoftData() *TrustSoftData {
	if m != nil {
		return m.SoftData
	}
	return nil
}

func (m *TrustResponseData) GetDeviceToken() string {
	if m != nil && m.DeviceToken != nil {
		return *m.DeviceToken
	}
	return ""
}

func (m *TrustResponseData) GetTimeStamp() uint32 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

type TrustResponse struct {
	BaseResponse         *BaseResponse      `protobuf:"bytes,1,req,name=BaseResponse" json:"BaseResponse,omitempty"`
	TrustResponseData    *TrustResponseData `protobuf:"bytes,2,req,name=TrustResponseData" json:"TrustResponseData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrustResponse) Reset()         { *m = TrustResponse{} }
func (m *TrustResponse) String() string { return proto.CompactTextString(m) }
func (*TrustResponse) ProtoMessage()    {}
func (*TrustResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efeafd76de4e3385, []int{5}
}

func (m *TrustResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustResponse.Unmarshal(m, b)
}
func (m *TrustResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustResponse.Marshal(b, m, deterministic)
}
func (m *TrustResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustResponse.Merge(m, src)
}
func (m *TrustResponse) XXX_Size() int {
	return xxx_messageInfo_TrustResponse.Size(m)
}
func (m *TrustResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrustResponse proto.InternalMessageInfo

func (m *TrustResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *TrustResponse) GetTrustResponseData() *TrustResponseData {
	if m != nil {
		return m.TrustResponseData
	}
	return nil
}

func init() {
	proto.RegisterType((*TrustDeviceInfo)(nil), "mmproto.TrustDeviceInfo")
	proto.RegisterType((*TrustData)(nil), "mmproto.TrustData")
	proto.RegisterType((*TrustReq)(nil), "mmproto.TrustReq")
	proto.RegisterType((*TrustSoftData)(nil), "mmproto.TrustSoftData")
	proto.RegisterType((*TrustResponseData)(nil), "mmproto.TrustResponseData")
	proto.RegisterType((*TrustResponse)(nil), "mmproto.TrustResponse")
}

func init() {
	proto.RegisterFile("trustInfo.proto", fileDescriptor_efeafd76de4e3385)
}

var fileDescriptor_efeafd76de4e3385 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcf, 0x4a, 0xfb, 0x30,
	0x1c, 0xa7, 0xe9, 0x0f, 0x7e, 0xeb, 0xb7, 0x9b, 0xd3, 0x80, 0x12, 0x86, 0x48, 0xc9, 0xa9, 0x78,
	0xd8, 0xa1, 0x20, 0xe2, 0xc5, 0x83, 0x7a, 0x50, 0x76, 0xcb, 0xf6, 0x02, 0xd1, 0xa4, 0x52, 0x5c,
	0x9a, 0xba, 0xc4, 0x82, 0x4f, 0xe0, 0x0b, 0xf8, 0xc0, 0xd2, 0x34, 0xdd, 0xda, 0xea, 0x2d, 0xf9,
	0xe4, 0xf3, 0x37, 0x30, 0xb7, 0xbb, 0x0f, 0x63, 0x9f, 0xca, 0x5c, 0x2f, 0xab, 0x9d, 0xb6, 0x1a,
	0xff, 0x57, 0xca, 0x1d, 0x16, 0xf0, 0xcc, 0x8d, 0x6c, 0x41, 0x7a, 0x05, 0xf3, 0x4d, 0xc3, 0x7b,
	0x90, 0x75, 0xf1, 0x22, 0x1b, 0x36, 0x3e, 0x86, 0x70, 0x25, 0x3f, 0x49, 0x90, 0xa0, 0x34, 0x62,
	0xcd, 0xb1, 0x41, 0x6a, 0xbe, 0x25, 0xa8, 0x45, 0x6a, 0xbe, 0xa5, 0xd7, 0x10, 0xb5, 0x32, 0x6e,
	0x39, 0xbe, 0x84, 0xd0, 0x8a, 0x82, 0x04, 0x49, 0x98, 0xc6, 0x19, 0x59, 0xfa, 0x98, 0xe5, 0xc8,
	0x97, 0x35, 0x24, 0x7a, 0x0b, 0x13, 0x87, 0x33, 0xf9, 0x8e, 0x29, 0x20, 0x2b, 0x5c, 0x4e, 0x9c,
	0xe1, 0x91, 0x8c, 0x5b, 0xce, 0x90, 0x15, 0xf8, 0x08, 0x90, 0x12, 0x04, 0x25, 0x41, 0x1a, 0x31,
	0xa4, 0x04, 0x5d, 0xc1, 0xcc, 0x11, 0xd6, 0x3a, 0x6f, 0xc3, 0x2f, 0x00, 0x8c, 0xce, 0xed, 0xbd,
	0x2e, 0xf3, 0xe2, 0xd5, 0x99, 0x4d, 0x59, 0x0f, 0xc1, 0x0b, 0x98, 0x18, 0xcf, 0x75, 0x03, 0xa6,
	0x6c, 0x7f, 0xa7, 0x5f, 0x01, 0x9c, 0xf8, 0x36, 0xa6, 0xd2, 0xa5, 0x91, 0xce, 0x31, 0x1b, 0x29,
	0xe2, 0xec, 0x6c, 0x58, 0xae, 0xcb, 0x3e, 0x38, 0xe1, 0x04, 0x62, 0xe1, 0x96, 0x6e, 0xf4, 0x9b,
	0x2c, 0x49, 0xe8, 0x7e, 0xaa, 0x0f, 0xe1, 0x73, 0x88, 0x6c, 0xa1, 0xe4, 0xda, 0x72, 0x55, 0x91,
	0x7f, 0x09, 0x4a, 0x67, 0xec, 0x00, 0xd0, 0xef, 0xc0, 0xef, 0xea, 0x9a, 0xe0, 0x1b, 0x98, 0xde,
	0x71, 0x23, 0xbb, 0xbb, 0xff, 0xa6, 0xd3, 0x7d, 0x93, 0xfe, 0x23, 0x1b, 0x50, 0xf1, 0xe3, 0x1f,
	0xab, 0xfc, 0x92, 0xc5, 0x70, 0x49, 0x9f, 0xc1, 0x7e, 0x8b, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x48, 0x36, 0x78, 0x61, 0x44, 0x02, 0x00, 0x00,
}