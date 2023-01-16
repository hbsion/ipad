// Code generated by protoc-gen-go. DO NOT EDIT.
// source: massSend.proto

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

type MassSendRequest struct {
	BaseRequest *BaseRequest `protobuf:"bytes,1,opt,name=baseRequest" json:"baseRequest,omitempty"`
	/*CameraType           *uint64           `protobuf:"varint,13,opt,name=cameraType" json:"cameraType,omitempty"`
	ClientId             *string           `protobuf:"bytes,4,opt,name=clientId" json:"clientId,omitempty"`
	CompressType         *uint64           `protobuf:"varint,17,opt,name=compressType" json:"compressType,omitempty"`
	DataBuffer           *SKBuiltinString_ `protobuf:"bytes,7,opt,name=dataBuffer" json:"dataBuffer,omitempty"`
	DataStartPos         *uint64           `protobuf:"varint,8,opt,name=dataStartPos" json:"dataStartPos,omitempty"`
	DataTotalLen         *uint64           `protobuf:"varint,9,opt,name=dataTotalLen" json:"dataTotalLen,omitempty"`
	IsSendAgain          *uint64           `protobuf:"varint,16,opt,name=isSendAgain" json:"isSendAgain,omitempty"`
	MediaTime            *uint64           `protobuf:"varint,6,opt,name=mediaTime" json:"mediaTime,omitempty"`
	MsgType              *uint64           `protobuf:"varint,5,opt,name=msgType" json:"msgType,omitempty"`
	ThumbAeskey          *string           `protobuf:"bytes,23,opt,name=thumbAeskey" json:"thumbAeskey,omitempty"`
	ThumbData            *SKBuiltinString_ `protobuf:"bytes,12,opt,name=thumbData" json:"thumbData,omitempty"`
	ThumbHeight          *uint64           `protobuf:"varint,22,opt,name=thumbHeight" json:"thumbHeight,omitempty"`
	ThumbStartPos        *uint64           `protobuf:"varint,11,opt,name=thumbStartPos" json:"thumbStartPos,omitempty"`
	ThumbTotalLen        *uint64           `protobuf:"varint,10,opt,name=thumbTotalLen" json:"thumbTotalLen,omitempty"`
	ThumbUrl             *string           `protobuf:"bytes,20,opt,name=thumbUrl" json:"thumbUrl,omitempty"`
	ThumbWith            *uint64           `protobuf:"varint,21,opt,name=thumbWith" json:"thumbWith,omitempty"`
	ToList               *string           `protobuf:"bytes,2,opt,name=toList" json:"toList,omitempty"`
	ToListCount          *uint64           `protobuf:"varint,15,opt,name=toListCount" json:"toListCount,omitempty"`
	ToListMd5            *string           `protobuf:"bytes,3,opt,name=toListMd5" json:"toListMd5,omitempty"`
	VideoAeskey          *string           `protobuf:"bytes,24,opt,name=videoAeskey" json:"videoAeskey,omitempty"`
	VideoSource          *uint64           `protobuf:"varint,14,opt,name=videoSource" json:"videoSource,omitempty"`
	VideoUrl             *string           `protobuf:"bytes,19,opt,name=videoUrl" json:"videoUrl,omitempty"`
	VoiceFormat          *uint64           `protobuf:"varint,18,opt,name=voiceFormat" json:"voiceFormat,omitempty"`*/

	ToList        *string           `protobuf:"bytes,2,opt,name=ToList" json:"ToList,omitempty"`
	ToListMd5     *string           `protobuf:"bytes,3,opt,name=ToListMd5" json:"ToListMd5,omitempty"`
	ClientId      *string           `protobuf:"bytes,4,opt,name=ClientId" json:"ClientId,omitempty"`
	MsgType       *uint64           `protobuf:"varint,5,opt,name=MsgType" json:"MsgType,omitempty"`
	MediaTime     *uint64           `protobuf:"varint,6,opt,name=MediaTime" json:"MediaTime,omitempty"`
	DataBuffer    *SKBuiltinString_ `protobuf:"bytes,7,opt,name=DataBuffer" json:"DataBuffer,omitempty"`
	DataStartPos  *uint64           `protobuf:"varint,8,opt,name=DataStartPos" json:"DataStartPos,omitempty"`
	DataTotalLen  *uint64           `protobuf:"varint,9,opt,name=DataTotalLen" json:"DataTotalLen,omitempty"`
	ThumbTotalLen *uint64           `protobuf:"varint,10,opt,name=ThumbTotalLen" json:"ThumbTotalLen,omitempty"`
	ThumbStartPos *uint64           `protobuf:"varint,11,opt,name=ThumbStartPos" json:"ThumbStartPos,omitempty"`
	ThumbData     *SKBuiltinString_ `protobuf:"bytes,12,opt,name=ThumbData" json:"ThumbData,omitempty"`
	CameraType    *uint64           `protobuf:"varint,13,opt,name=Cameratype" json:"Cameratype,omitempty"`
	VideoSource   *uint64           `protobuf:"varint,14,opt,name=VideoSource" json:"VideoSource,omitempty"`
	ToListCount   *uint64           `protobuf:"varint,15,opt,name=ToListCount" json:"ToListCount,omitempty"`
	IsSendAgain   *uint64           `protobuf:"varint,16,opt,name=IsSendAgain" json:"IsSendAgain,omitempty"`
	CompressType  *uint64           `protobuf:"varint,17,opt,name=CompressType" json:"CompressType,omitempty"`
	VoiceFormat   *uint64           `protobuf:"varint,18,opt,name=VoiceFormat" json:"VoiceFormat,omitempty"`
	VideoUrl      *string           `protobuf:"bytes,19,opt,name=VideoUrl" json:"VideoUrl,omitempty"`
	ThumbUrl      *string           `protobuf:"bytes,20,opt,name=ThumbUrl" json:"ThumbUrl,omitempty"`
	ThumbWith     *uint64           `protobuf:"varint,21,opt,name=ThumbWidth" json:"ThumbWidth,omitempty"`
	ThumbHeight   *uint64           `protobuf:"varint,22,opt,name=ThumbHeight" json:"ThumbHeight,omitempty"`
	ThumbAeskey   *string           `protobuf:"bytes,23,opt,name=ThumbAESKey" json:"ThumbAESKey,omitempty"`
	VideoAeskey   *string           `protobuf:"bytes,24,opt,name=VideoAESKey" json:"VideoAESKey,omitempty"`
	MD5           *string           `protobuf:"bytes,25,opt,name=MD5" json:"MD5,omitempty"`

	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MassSendRequest) Reset()         { *m = MassSendRequest{} }
func (m *MassSendRequest) String() string { return proto.CompactTextString(m) }
func (*MassSendRequest) ProtoMessage()    {}
func (*MassSendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20167eb4258e0ea3, []int{0}
}

func (m *MassSendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MassSendRequest.Unmarshal(m, b)
}
func (m *MassSendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MassSendRequest.Marshal(b, m, deterministic)
}
func (m *MassSendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MassSendRequest.Merge(m, src)
}
func (m *MassSendRequest) XXX_Size() int {
	return xxx_messageInfo_MassSendRequest.Size(m)
}
func (m *MassSendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MassSendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MassSendRequest proto.InternalMessageInfo

func (m *MassSendRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *MassSendRequest) GetCameraType() uint64 {
	if m != nil && m.CameraType != nil {
		return *m.CameraType
	}
	return 0
}

func (m *MassSendRequest) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *MassSendRequest) GetCompressType() uint64 {
	if m != nil && m.CompressType != nil {
		return *m.CompressType
	}
	return 0
}

func (m *MassSendRequest) GetDataBuffer() *SKBuiltinString_ {
	if m != nil {
		return m.DataBuffer
	}
	return nil
}

func (m *MassSendRequest) GetDataStartPos() uint64 {
	if m != nil && m.DataStartPos != nil {
		return *m.DataStartPos
	}
	return 0
}

func (m *MassSendRequest) GetDataTotalLen() uint64 {
	if m != nil && m.DataTotalLen != nil {
		return *m.DataTotalLen
	}
	return 0
}

func (m *MassSendRequest) GetIsSendAgain() uint64 {
	if m != nil && m.IsSendAgain != nil {
		return *m.IsSendAgain
	}
	return 0
}

func (m *MassSendRequest) GetMediaTime() uint64 {
	if m != nil && m.MediaTime != nil {
		return *m.MediaTime
	}
	return 0
}

func (m *MassSendRequest) GetMsgType() uint64 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *MassSendRequest) GetThumbAeskey() string {
	if m != nil && m.ThumbAeskey != nil {
		return *m.ThumbAeskey
	}
	return ""
}

func (m *MassSendRequest) GetThumbData() *SKBuiltinString_ {
	if m != nil {
		return m.ThumbData
	}
	return nil
}

func (m *MassSendRequest) GetThumbHeight() uint64 {
	if m != nil && m.ThumbHeight != nil {
		return *m.ThumbHeight
	}
	return 0
}

func (m *MassSendRequest) GetThumbStartPos() uint64 {
	if m != nil && m.ThumbStartPos != nil {
		return *m.ThumbStartPos
	}
	return 0
}

func (m *MassSendRequest) GetThumbTotalLen() uint64 {
	if m != nil && m.ThumbTotalLen != nil {
		return *m.ThumbTotalLen
	}
	return 0
}

func (m *MassSendRequest) GetThumbUrl() string {
	if m != nil && m.ThumbUrl != nil {
		return *m.ThumbUrl
	}
	return ""
}

func (m *MassSendRequest) GetThumbWith() uint64 {
	if m != nil && m.ThumbWith != nil {
		return *m.ThumbWith
	}
	return 0
}

func (m *MassSendRequest) GetToList() string {
	if m != nil && m.ToList != nil {
		return *m.ToList
	}
	return ""
}

func (m *MassSendRequest) GetToListCount() uint64 {
	if m != nil && m.ToListCount != nil {
		return *m.ToListCount
	}
	return 0
}

func (m *MassSendRequest) GetToListMd5() string {
	if m != nil && m.ToListMd5 != nil {
		return *m.ToListMd5
	}
	return ""
}

func (m *MassSendRequest) GetVideoAeskey() string {
	if m != nil && m.VideoAeskey != nil {
		return *m.VideoAeskey
	}
	return ""
}

func (m *MassSendRequest) GetVideoSource() uint64 {
	if m != nil && m.VideoSource != nil {
		return *m.VideoSource
	}
	return 0
}

func (m *MassSendRequest) GetVideoUrl() string {
	if m != nil && m.VideoUrl != nil {
		return *m.VideoUrl
	}
	return ""
}

func (m *MassSendRequest) GetVoiceFormat() uint64 {
	if m != nil && m.VoiceFormat != nil {
		return *m.VoiceFormat
	}
	return 0
}

type MassSendResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,opt,name=baseResponse" json:"baseResponse,omitempty"`
	DataStartPos         *uint64       `protobuf:"varint,2,opt,name=dataStartPos" json:"dataStartPos,omitempty"`
	MaxSupport           *uint64       `protobuf:"varint,4,opt,name=maxSupport" json:"maxSupport,omitempty"`
	ThumbStartPos        *uint64       `protobuf:"varint,3,opt,name=thumbStartPos" json:"thumbStartPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MassSendResponse) Reset()         { *m = MassSendResponse{} }
func (m *MassSendResponse) String() string { return proto.CompactTextString(m) }
func (*MassSendResponse) ProtoMessage()    {}
func (*MassSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20167eb4258e0ea3, []int{1}
}

func (m *MassSendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MassSendResponse.Unmarshal(m, b)
}
func (m *MassSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MassSendResponse.Marshal(b, m, deterministic)
}
func (m *MassSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MassSendResponse.Merge(m, src)
}
func (m *MassSendResponse) XXX_Size() int {
	return xxx_messageInfo_MassSendResponse.Size(m)
}
func (m *MassSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MassSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MassSendResponse proto.InternalMessageInfo

func (m *MassSendResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *MassSendResponse) GetDataStartPos() uint64 {
	if m != nil && m.DataStartPos != nil {
		return *m.DataStartPos
	}
	return 0
}

func (m *MassSendResponse) GetMaxSupport() uint64 {
	if m != nil && m.MaxSupport != nil {
		return *m.MaxSupport
	}
	return 0
}

func (m *MassSendResponse) GetThumbStartPos() uint64 {
	if m != nil && m.ThumbStartPos != nil {
		return *m.ThumbStartPos
	}
	return 0
}

func init() {
	proto.RegisterType((*MassSendRequest)(nil), "wechat_proto.MassSendRequest")
	proto.RegisterType((*MassSendResponse)(nil), "wechat_proto.MassSendResponse")
}

func init() { proto.RegisterFile("massSend.proto", fileDescriptor_20167eb4258e0ea3) }

var fileDescriptor_20167eb4258e0ea3 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0xa5, 0xbb, 0x6b, 0xb7, 0xbd, 0xd3, 0xfd, 0x30, 0xea, 0x7a, 0x2d, 0x52, 0x4a, 0xf1, 0xa1,
	0x4f, 0x23, 0x08, 0x3e, 0x29, 0xc2, 0x56, 0x11, 0xc5, 0x5d, 0x90, 0x99, 0x8a, 0xe0, 0xcb, 0x92,
	0x9d, 0xc9, 0xb6, 0xc1, 0xce, 0x64, 0x4c, 0x32, 0xea, 0xfe, 0x37, 0x7f, 0x95, 0xbf, 0x40, 0xe6,
	0x66, 0x3e, 0x52, 0xb6, 0x0f, 0xbe, 0xe5, 0x9c, 0x9c, 0x93, 0xdc, 0x7b, 0x92, 0x0b, 0xc7, 0x19,
	0x37, 0x26, 0x16, 0x79, 0x1a, 0x16, 0x5a, 0x59, 0xc5, 0x46, 0xbf, 0x44, 0xb2, 0xe6, 0xf6, 0x8a,
	0xd0, 0xb8, 0x46, 0x6e, 0x6f, 0xf6, 0xb7, 0x0f, 0x27, 0x97, 0xb5, 0x3c, 0x12, 0x3f, 0x4a, 0x61,
	0x2c, 0x7b, 0x05, 0xc1, 0x35, 0x37, 0xa2, 0x86, 0xd8, 0x9b, 0xf6, 0xe6, 0xc1, 0x8b, 0x27, 0xa1,
	0x7f, 0x4a, 0xb8, 0xe8, 0x04, 0x91, 0xaf, 0x66, 0x13, 0x80, 0x84, 0x67, 0x42, 0xf3, 0xe5, 0x6d,
	0x21, 0xf0, 0x68, 0xda, 0x9b, 0x1f, 0x44, 0x1e, 0xc3, 0xc6, 0x30, 0x48, 0x36, 0x52, 0xe4, 0xf6,
	0x63, 0x8a, 0x07, 0xd3, 0xde, 0x7c, 0x18, 0xb5, 0x98, 0xcd, 0x60, 0x94, 0xa8, 0xac, 0xd0, 0xc2,
	0x18, 0x72, 0xdf, 0x27, 0xf7, 0x16, 0xc7, 0xde, 0x00, 0xa4, 0xdc, 0xf2, 0x45, 0x79, 0x73, 0x23,
	0x34, 0x1e, 0x52, 0x6d, 0x93, 0xed, 0xda, 0xe2, 0x4f, 0x8b, 0x52, 0x6e, 0xac, 0xcc, 0x63, 0xab,
	0x65, 0xbe, 0xba, 0x8a, 0x3c, 0x47, 0x75, 0x47, 0x85, 0x62, 0xcb, 0xb5, 0xfd, 0xac, 0x0c, 0x0e,
	0xdc, 0x1d, 0x3e, 0xd7, 0x68, 0x96, 0xca, 0xf2, 0xcd, 0x85, 0xc8, 0x71, 0xd8, 0x69, 0x1a, 0x8e,
	0x4d, 0x21, 0x90, 0x94, 0xda, 0xf9, 0x8a, 0xcb, 0x1c, 0x4f, 0x49, 0xe2, 0x53, 0xec, 0x29, 0x0c,
	0x33, 0x91, 0x4a, 0xbe, 0x94, 0x99, 0xc0, 0x3e, 0xed, 0x77, 0x04, 0x43, 0x38, 0xcc, 0xcc, 0x8a,
	0xda, 0xbc, 0x47, 0x7b, 0x0d, 0xac, 0x4e, 0xb6, 0xeb, 0x32, 0xbb, 0x3e, 0x17, 0xe6, 0xbb, 0xb8,
	0xc5, 0xc7, 0x14, 0x92, 0x4f, 0xb1, 0xd7, 0x30, 0x24, 0xf8, 0x8e, 0x5b, 0x8e, 0xa3, 0xff, 0x8a,
	0xa0, 0x33, 0xb4, 0xe7, 0x7f, 0x10, 0x72, 0xb5, 0xb6, 0x78, 0xe6, 0x2a, 0xf7, 0x28, 0xf6, 0x0c,
	0x8e, 0x08, 0xb6, 0x21, 0x05, 0xa4, 0xd9, 0x26, 0x5b, 0x55, 0x1b, 0x13, 0x78, 0xaa, 0x36, 0xa7,
	0x31, 0x0c, 0x88, 0xf8, 0xa2, 0x37, 0xf8, 0xd0, 0xbd, 0x77, 0x83, 0xab, 0x84, 0x68, 0xfd, 0x55,
	0xda, 0x35, 0x3e, 0x72, 0x09, 0xb5, 0x04, 0x3b, 0x83, 0xbe, 0x55, 0x17, 0xd2, 0x58, 0xdc, 0x23,
	0x5f, 0x8d, 0xa8, 0x7e, 0x5a, 0xbd, 0x55, 0x65, 0x6e, 0xf1, 0xa4, 0xae, 0xbf, 0xa3, 0xe8, 0x5c,
	0x82, 0x97, 0xe9, 0x4b, 0xdc, 0x27, 0x73, 0x47, 0x54, 0xfe, 0x9f, 0x32, 0x15, 0xaa, 0xce, 0x17,
	0x5d, 0xbe, 0x1e, 0xd5, 0x2a, 0x62, 0x55, 0xea, 0x44, 0xe0, 0xb1, 0xbb, 0xc1, 0xa3, 0xaa, 0xae,
	0x08, 0x56, 0x5d, 0x3d, 0x70, 0x5d, 0x35, 0x98, 0xdc, 0x4a, 0x26, 0xe2, 0xbd, 0xd2, 0x19, 0xb7,
	0xc8, 0x6a, 0x77, 0x47, 0xcd, 0xfe, 0xf4, 0xe0, 0xb4, 0x1b, 0x3a, 0x53, 0xa8, 0xdc, 0x54, 0x1f,
	0x7b, 0xe4, 0xe6, 0xc8, 0xe1, 0x7a, 0xec, 0xc6, 0xbb, 0xc6, 0xce, 0x29, 0xa2, 0x2d, 0xfd, 0x9d,
	0x8f, 0xbd, 0xb7, 0xe3, 0x63, 0x4f, 0x00, 0x32, 0xfe, 0x3b, 0x2e, 0x8b, 0x42, 0x69, 0x4b, 0xe3,
	0x77, 0x10, 0x79, 0xcc, 0xdd, 0x87, 0xdf, 0xdf, 0xf1, 0xf0, 0x8b, 0xe0, 0xdb, 0x30, 0x0c, 0x9f,
	0xbb, 0xba, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x98, 0x9a, 0x8e, 0xa6, 0x6d, 0x04, 0x00, 0x00,
}
