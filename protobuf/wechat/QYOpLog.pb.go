// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qyopog.proto

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

type QYOpLogRequest struct {
	Type                 *int64   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	V                    []byte   `protobuf:"bytes,2,rep,name=V" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QYOpLogRequest) Reset()         { *m = QYOpLogRequest{} }
func (m *QYOpLogRequest) String() string { return proto.CompactTextString(m) }
func (*QYOpLogRequest) ProtoMessage()    {}
func (*QYOpLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70024e181f871431, []int{0}
}

func (m *QYOpLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QYOpLogRequest.Unmarshal(m, b)
}
func (m *QYOpLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QYOpLogRequest.Marshal(b, m, deterministic)
}
func (m *QYOpLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QYOpLogRequest.Merge(m, src)
}
func (m *QYOpLogRequest) XXX_Size() int {
	return xxx_messageInfo_QYOpLogRequest.Size(m)
}
func (m *QYOpLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QYOpLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QYOpLogRequest proto.InternalMessageInfo

func (m *QYOpLogRequest) GetType() int64 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *QYOpLogRequest) GetV() []byte {
	if m != nil {
		return m.V
	}
	return nil
}

type ModInfo struct {
	Cmd                  *uint32  `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	Value                *string  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModInfo) Reset()         { *m = ModInfo{} }
func (m *ModInfo) String() string { return proto.CompactTextString(m) }
func (*ModInfo) ProtoMessage()    {}
func (*ModInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_70024e181f871431, []int{1}
}

func (m *ModInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModInfo.Unmarshal(m, b)
}
func (m *ModInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModInfo.Marshal(b, m, deterministic)
}
func (m *ModInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModInfo.Merge(m, src)
}
func (m *ModInfo) XXX_Size() int {
	return xxx_messageInfo_ModInfo.Size(m)
}
func (m *ModInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModInfo proto.InternalMessageInfo

func (m *ModInfo) GetCmd() uint32 {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return 0
}

func (m *ModInfo) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*QYOpLogRequest)(nil), "wechat_proto.QYOpLogRequest")
	proto.RegisterType((*ModInfo)(nil), "wechat_proto.ModInfo")
}

func init() { proto.RegisterFile("qyopog.proto", fileDescriptor_70024e181f871431) }

var fileDescriptor_70024e181f871431 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xac, 0xcc, 0x2f,
	0xc8, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x4f, 0x4d, 0xce, 0x48, 0x2c,
	0x89, 0x07, 0xf3, 0xa4, 0xa0, 0x3c, 0x88, 0x9c, 0x92, 0x11, 0x17, 0x5f, 0x60, 0xa4, 0x7f, 0x81,
	0x4f, 0x7e, 0x7a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65,
	0x41, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x98, 0x2d, 0xc4, 0xc3, 0xc5, 0x18, 0x26,
	0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x13, 0xc4, 0x18, 0xa6, 0x64, 0xc8, 0xc5, 0xee, 0x9b, 0x9f, 0xe2,
	0x99, 0x97, 0x96, 0x2f, 0x24, 0xc0, 0xc5, 0x9c, 0x9c, 0x9b, 0x02, 0x56, 0xcb, 0x1b, 0x04, 0x62,
	0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x38, 0x4e, 0xdc, 0x51, 0x9c, 0x7a, 0x7a, 0xfa, 0x10, 0x9b, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0xcc, 0x8e, 0x47, 0x9e, 0x00, 0x00, 0x00,
}