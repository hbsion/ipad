// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fav.proto

package wechat

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

type ShareFavRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,req,name=BaseRequest" json:"BaseRequest,omitempty"`
	ToUser               *string      `protobuf:"bytes,2,opt,name=ToUser" json:"ToUser,omitempty"`
	Scene                *uint32      `protobuf:"varint,3,req,name=Scene" json:"Scene,omitempty"`
	Count                *uint32      `protobuf:"varint,4,req,name=Count" json:"Count,omitempty"`
	FavIdList            []uint32     `protobuf:"varint,5,rep,packed,name=FavIdList" json:"FavIdList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ShareFavRequest) Reset()         { *m = ShareFavRequest{} }
func (m *ShareFavRequest) String() string { return proto.CompactTextString(m) }
func (*ShareFavRequest) ProtoMessage()    {}
func (*ShareFavRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{9}
}
func (m *ShareFavRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareFavRequest.Unmarshal(m, b)
}
func (m *ShareFavRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareFavRequest.Marshal(b, m, deterministic)
}
func (dst *ShareFavRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareFavRequest.Merge(dst, src)
}
func (m *ShareFavRequest) XXX_Size() int {
	return xxx_messageInfo_ShareFavRequest.Size(m)
}
func (m *ShareFavRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareFavRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareFavRequest proto.InternalMessageInfo

func (m *ShareFavRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *ShareFavRequest) GetToUser() string {
	if m != nil && m.ToUser != nil {
		return *m.ToUser
	}
	return ""
}

func (m *ShareFavRequest) GetScene() uint32 {
	if m != nil && m.Scene != nil {
		return *m.Scene
	}
	return 0
}

func (m *ShareFavRequest) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *ShareFavRequest) GetFavIdList() []uint32 {
	if m != nil {
		return m.FavIdList
	}
	return nil
}

type ShareFavResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,req,name=BaseResponse" json:"BaseResponse,omitempty"`
	UrlCount             *uint32       `protobuf:"varint,2,req,name=UrlCount" json:"UrlCount,omitempty"`
	UrlList              []*StringT    `protobuf:"bytes,3,rep,name=UrlList" json:"UrlList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShareFavResponse) Reset()         { *m = ShareFavResponse{} }
func (m *ShareFavResponse) String() string { return proto.CompactTextString(m) }
func (*ShareFavResponse) ProtoMessage()    {}
func (*ShareFavResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{10}
}
func (m *ShareFavResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareFavResponse.Unmarshal(m, b)
}
func (m *ShareFavResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareFavResponse.Marshal(b, m, deterministic)
}
func (dst *ShareFavResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareFavResponse.Merge(dst, src)
}
func (m *ShareFavResponse) XXX_Size() int {
	return xxx_messageInfo_ShareFavResponse.Size(m)
}
func (m *ShareFavResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareFavResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareFavResponse proto.InternalMessageInfo

func (m *ShareFavResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *ShareFavResponse) GetUrlCount() uint32 {
	if m != nil && m.UrlCount != nil {
		return *m.UrlCount
	}
	return 0
}

func (m *ShareFavResponse) GetUrlList() []*StringT {
	if m != nil {
		return m.UrlList
	}
	return nil
}

type CheckCDN struct {
	DataId               *string  `protobuf:"bytes,1,opt,name=DataId" json:"DataId,omitempty"`
	FullMd5              *string  `protobuf:"bytes,2,opt,name=FullMd5" json:"FullMd5,omitempty"`
	Head256Md5           *string  `protobuf:"bytes,3,opt,name=Head256Md5" json:"Head256Md5,omitempty"`
	FullSize             *uint32  `protobuf:"varint,4,req,name=FullSize" json:"FullSize,omitempty"`
	DataSourceType       *uint32  `protobuf:"varint,5,req,name=DataSourceType" json:"DataSourceType,omitempty"`
	DataSourceId         *string  `protobuf:"bytes,6,opt,name=DataSourceId" json:"DataSourceId,omitempty"`
	IsThumb              *uint32  `protobuf:"varint,7,req,name=IsThumb" json:"IsThumb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckCDN) Reset()         { *m = CheckCDN{} }
func (m *CheckCDN) String() string { return proto.CompactTextString(m) }
func (*CheckCDN) ProtoMessage()    {}
func (*CheckCDN) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{19}
}
func (m *CheckCDN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckCDN.Unmarshal(m, b)
}
func (m *CheckCDN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckCDN.Marshal(b, m, deterministic)
}
func (dst *CheckCDN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCDN.Merge(dst, src)
}
func (m *CheckCDN) XXX_Size() int {
	return xxx_messageInfo_CheckCDN.Size(m)
}
func (m *CheckCDN) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCDN.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCDN proto.InternalMessageInfo

func (m *CheckCDN) GetDataId() string {
	if m != nil && m.DataId != nil {
		return *m.DataId
	}
	return ""
}

func (m *CheckCDN) GetFullMd5() string {
	if m != nil && m.FullMd5 != nil {
		return *m.FullMd5
	}
	return ""
}

func (m *CheckCDN) GetHead256Md5() string {
	if m != nil && m.Head256Md5 != nil {
		return *m.Head256Md5
	}
	return ""
}

func (m *CheckCDN) GetFullSize() uint32 {
	if m != nil && m.FullSize != nil {
		return *m.FullSize
	}
	return 0
}

func (m *CheckCDN) GetDataSourceType() uint32 {
	if m != nil && m.DataSourceType != nil {
		return *m.DataSourceType
	}
	return 0
}

func (m *CheckCDN) GetDataSourceId() string {
	if m != nil && m.DataSourceId != nil {
		return *m.DataSourceId
	}
	return ""
}

func (m *CheckCDN) GetIsThumb() uint32 {
	if m != nil && m.IsThumb != nil {
		return *m.IsThumb
	}
	return 0
}

type FavCDNItem struct {
	DataId               *string  `protobuf:"bytes,1,opt,name=DataId" json:"DataId,omitempty"`
	FullMd5              *string  `protobuf:"bytes,2,opt,name=FullMd5" json:"FullMd5,omitempty"`
	Head256Md5           *string  `protobuf:"bytes,3,opt,name=Head256Md5" json:"Head256Md5,omitempty"`
	FullSize             *uint32  `protobuf:"varint,4,req,name=FullSize" json:"FullSize,omitempty"`
	CDNURL               *string  `protobuf:"bytes,5,opt,name=CDNURL" json:"CDNURL,omitempty"`
	AESKey               *string  `protobuf:"bytes,6,opt,name=AESKey" json:"AESKey,omitempty"`
	EncryVer             *int32   `protobuf:"varint,7,req,name=EncryVer" json:"EncryVer,omitempty"`
	VideoId              *string  `protobuf:"bytes,8,opt,name=VideoId" json:"VideoId,omitempty"`
	Status               *int32   `protobuf:"varint,9,req,name=Status" json:"Status,omitempty"`
	DataStatus           *int32   `protobuf:"varint,10,req,name=DataStatus" json:"DataStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavCDNItem) Reset()         { *m = FavCDNItem{} }
func (m *FavCDNItem) String() string { return proto.CompactTextString(m) }
func (*FavCDNItem) ProtoMessage()    {}
func (*FavCDNItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{20}
}
func (m *FavCDNItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavCDNItem.Unmarshal(m, b)
}
func (m *FavCDNItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavCDNItem.Marshal(b, m, deterministic)
}
func (dst *FavCDNItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavCDNItem.Merge(dst, src)
}
func (m *FavCDNItem) XXX_Size() int {
	return xxx_messageInfo_FavCDNItem.Size(m)
}
func (m *FavCDNItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FavCDNItem.DiscardUnknown(m)
}

var xxx_messageInfo_FavCDNItem proto.InternalMessageInfo

func (m *FavCDNItem) GetDataId() string {
	if m != nil && m.DataId != nil {
		return *m.DataId
	}
	return ""
}

func (m *FavCDNItem) GetFullMd5() string {
	if m != nil && m.FullMd5 != nil {
		return *m.FullMd5
	}
	return ""
}

func (m *FavCDNItem) GetHead256Md5() string {
	if m != nil && m.Head256Md5 != nil {
		return *m.Head256Md5
	}
	return ""
}

func (m *FavCDNItem) GetFullSize() uint32 {
	if m != nil && m.FullSize != nil {
		return *m.FullSize
	}
	return 0
}

func (m *FavCDNItem) GetCDNURL() string {
	if m != nil && m.CDNURL != nil {
		return *m.CDNURL
	}
	return ""
}

func (m *FavCDNItem) GetAESKey() string {
	if m != nil && m.AESKey != nil {
		return *m.AESKey
	}
	return ""
}

func (m *FavCDNItem) GetEncryVer() int32 {
	if m != nil && m.EncryVer != nil {
		return *m.EncryVer
	}
	return 0
}

func (m *FavCDNItem) GetVideoId() string {
	if m != nil && m.VideoId != nil {
		return *m.VideoId
	}
	return ""
}

func (m *FavCDNItem) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *FavCDNItem) GetDataStatus() int32 {
	if m != nil && m.DataStatus != nil {
		return *m.DataStatus
	}
	return 0
}

type CheckCDNRequest struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,req,name=BaseRequest" json:"BaseRequest,omitempty"`
	Count                *uint32      `protobuf:"varint,2,req,name=Count" json:"Count,omitempty"`
	List                 []*CheckCDN  `protobuf:"bytes,3,rep,name=List" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CheckCDNRequest) Reset()         { *m = CheckCDNRequest{} }
func (m *CheckCDNRequest) String() string { return proto.CompactTextString(m) }
func (*CheckCDNRequest) ProtoMessage()    {}
func (*CheckCDNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{23}
}
func (m *CheckCDNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckCDNRequest.Unmarshal(m, b)
}
func (m *CheckCDNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckCDNRequest.Marshal(b, m, deterministic)
}
func (dst *CheckCDNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCDNRequest.Merge(dst, src)
}
func (m *CheckCDNRequest) XXX_Size() int {
	return xxx_messageInfo_CheckCDNRequest.Size(m)
}
func (m *CheckCDNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCDNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCDNRequest proto.InternalMessageInfo

func (m *CheckCDNRequest) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *CheckCDNRequest) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CheckCDNRequest) GetList() []*CheckCDN {
	if m != nil {
		return m.List
	}
	return nil
}

type CheckCDNResponse struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,req,name=BaseResponse" json:"BaseResponse,omitempty"`
	Count                *uint32       `protobuf:"varint,2,req,name=Count" json:"Count,omitempty"`
	List                 []*FavCDNItem `protobuf:"bytes,3,rep,name=List" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckCDNResponse) Reset()         { *m = CheckCDNResponse{} }
func (m *CheckCDNResponse) String() string { return proto.CompactTextString(m) }
func (*CheckCDNResponse) ProtoMessage()    {}
func (*CheckCDNResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fav_079e142ae0efdd4b, []int{24}
}
func (m *CheckCDNResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckCDNResponse.Unmarshal(m, b)
}
func (m *CheckCDNResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckCDNResponse.Marshal(b, m, deterministic)
}
func (dst *CheckCDNResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCDNResponse.Merge(dst, src)
}
func (m *CheckCDNResponse) XXX_Size() int {
	return xxx_messageInfo_CheckCDNResponse.Size(m)
}
func (m *CheckCDNResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCDNResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCDNResponse proto.InternalMessageInfo

func (m *CheckCDNResponse) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *CheckCDNResponse) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CheckCDNResponse) GetList() []*FavCDNItem {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {

	proto.RegisterType((*ShareFavRequest)(nil), "pb.ShareFavRequest")
	proto.RegisterType((*ShareFavResponse)(nil), "pb.ShareFavResponse")
	proto.RegisterType((*CheckCDN)(nil), "pb.CheckCDN")
	proto.RegisterType((*FavCDNItem)(nil), "pb.FavCDNItem")
	proto.RegisterType((*CheckCDNRequest)(nil), "pb.CheckCDNRequest")
	proto.RegisterType((*CheckCDNResponse)(nil), "pb.CheckCDNResponse")

}

func init() { proto.RegisterFile("fav.proto", fileDescriptor_fav_079e142ae0efdd4b) }

var fileDescriptor_fav_079e142ae0efdd4b = []byte{
	// 1285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0x78, 0x6c, 0x27, 0x2e, 0xc7, 0x1b, 0xa7, 0x59, 0x82, 0x15, 0xa1, 0x95, 0xd5, 0xda,
	0x5d, 0x45, 0x48, 0x1b, 0xc1, 0x8a, 0x85, 0x73, 0x62, 0xc7, 0x10, 0x36, 0x09, 0x68, 0x6c, 0x47,
	0xdc, 0x56, 0x1d, 0x4f, 0x3b, 0x31, 0x8c, 0x67, 0xbc, 0x33, 0x6d, 0x13, 0x23, 0x4e, 0xfc, 0x09,
	0x24, 0x0e, 0x1c, 0x38, 0x20, 0xae, 0xdc, 0x78, 0x14, 0x24, 0x5e, 0x81, 0x57, 0x01, 0x75, 0x57,
	0xf7, 0x4c, 0x8f, 0xe3, 0xac, 0x56, 0x98, 0xc0, 0x6d, 0xbe, 0xaa, 0x76, 0x4d, 0xcd, 0xf7, 0x55,
	0x55, 0x97, 0xa1, 0x32, 0x64, 0xb3, 0xbd, 0x49, 0x1c, 0x89, 0x88, 0x14, 0x26, 0xe7, 0x3b, 0x55,
	0x31, 0x9f, 0xf0, 0x04, 0x0d, 0x3b, 0x70, 0xce, 0x12, 0x8e, 0xcf, 0xf4, 0x3b, 0x07, 0x60, 0xdf,
	0xf7, 0x3b, 0x6c, 0x76, 0x24, 0xf8, 0x98, 0xdc, 0x85, 0x92, 0x7c, 0xf4, 0x1b, 0x4e, 0xb3, 0xb0,
	0x5b, 0xf2, 0x10, 0x10, 0x02, 0xc5, 0xde, 0x7c, 0xc2, 0x1b, 0x05, 0x65, 0x54, 0xcf, 0xd2, 0xd6,
	0x09, 0xd8, 0x45, 0xc3, 0x6d, 0x16, 0x76, 0x6b, 0x9e, 0x7a, 0x26, 0xf7, 0x00, 0xfa, 0x13, 0x9f,
	0x09, 0xde, 0x1b, 0x8d, 0x79, 0xa3, 0xa8, 0x3c, 0x96, 0x85, 0xbc, 0x0e, 0x15, 0x44, 0x5d, 0xfe,
	0xbc, 0x51, 0x52, 0xee, 0xcc, 0x40, 0xbf, 0x76, 0xe0, 0x95, 0xd6, 0x25, 0x1f, 0x7c, 0xaa, 0x93,
	0xf1, 0xf8, 0xf3, 0x29, 0x4f, 0x04, 0x79, 0x0b, 0xaa, 0x07, 0x2c, 0xe1, 0x1a, 0xaa, 0xcc, 0xaa,
	0x8f, 0x37, 0xf7, 0x26, 0xe7, 0x7b, 0x96, 0xd9, 0xb3, 0xcf, 0xc8, 0x44, 0xba, 0xd1, 0x34, 0x1e,
	0xf0, 0x34, 0xed, 0x9a, 0x67, 0x59, 0xc8, 0x0e, 0xac, 0x23, 0x3a, 0xf2, 0x1b, 0x6e, 0xd3, 0xd9,
	0xad, 0x78, 0x29, 0xa6, 0x1f, 0xc3, 0x46, 0x87, 0xcd, 0xce, 0x46, 0x3e, 0x8f, 0x8e, 0xc2, 0x61,
	0x44, 0xb6, 0xa1, 0xdc, 0x19, 0x05, 0x5c, 0x71, 0x22, 0x4f, 0x6a, 0x24, 0xed, 0xfb, 0x3c, 0x79,
	0xca, 0xe7, 0x8d, 0x02, 0xda, 0x11, 0x91, 0x06, 0xac, 0xe1, 0x8f, 0x4d, 0x68, 0x03, 0xe9, 0x1f,
	0x0e, 0x54, 0xbb, 0x97, 0x2c, 0xe6, 0x1d, 0x36, 0xd3, 0x27, 0x3b, 0x71, 0x34, 0xee, 0x8f, 0x42,
	0xf5, 0x51, 0x35, 0xcf, 0x40, 0x19, 0xbb, 0x17, 0xf5, 0x13, 0x1e, 0x9b, 0xd8, 0x88, 0x24, 0x81,
	0x2a, 0x80, 0xe2, 0x17, 0x99, 0xcf, 0x0c, 0x52, 0xbc, 0xee, 0x80, 0x87, 0x86, 0x79, 0x04, 0x99,
	0xa4, 0x48, 0xb8, 0x96, 0xf4, 0x1e, 0xc0, 0x47, 0x2c, 0x66, 0x63, 0x45, 0x78, 0xa3, 0xac, 0xde,
	0x62, 0x59, 0x08, 0x55, 0x2c, 0x48, 0x19, 0xf0, 0xc4, 0x9a, 0x3a, 0x91, 0xb3, 0xd1, 0x63, 0xb8,
	0x9b, 0xd7, 0x2b, 0x99, 0x44, 0x61, 0xc2, 0xc9, 0xdb, 0xb0, 0x81, 0x62, 0x20, 0xd6, 0x8a, 0xd5,
	0x33, 0xc5, 0xd0, 0xee, 0xe5, 0x4e, 0xd1, 0x43, 0xa8, 0x9e, 0x44, 0xaa, 0x10, 0x43, 0x9f, 0x5f,
	0x49, 0x89, 0x24, 0x1c, 0xf1, 0xc0, 0xd7, 0xec, 0xa4, 0x58, 0xfb, 0xce, 0x58, 0x30, 0x35, 0xe2,
	0xa6, 0x98, 0x86, 0xb0, 0x81, 0x61, 0x3e, 0x3c, 0xff, 0x84, 0x0f, 0x84, 0x24, 0xb9, 0xc7, 0x2e,
	0x4e, 0xd9, 0x98, 0x6b, 0xfd, 0x0c, 0x94, 0x51, 0xf6, 0x85, 0x88, 0x95, 0x0b, 0x69, 0x4e, 0xb1,
	0x24, 0x0d, 0xc3, 0xa3, 0x84, 0x08, 0xd2, 0x3e, 0x28, 0x36, 0x1d, 0x59, 0xf3, 0xf2, 0x99, 0xfe,
	0xe5, 0xc0, 0x96, 0xce, 0x7b, 0xb5, 0x9a, 0x4d, 0x75, 0x2a, 0x2c, 0xe8, 0xa4, 0xf8, 0x68, 0x45,
	0xd3, 0x50, 0x68, 0xc9, 0x2d, 0x0b, 0x79, 0x04, 0x15, 0x85, 0x8e, 0x47, 0x89, 0x68, 0x14, 0x9b,
	0xae, 0x79, 0x8d, 0x45, 0xa5, 0x97, 0x9d, 0x20, 0x4d, 0xa8, 0x22, 0x2f, 0x18, 0x0f, 0x4b, 0xc2,
	0x36, 0x91, 0x37, 0x01, 0x10, 0xaa, 0x88, 0x65, 0x15, 0xb1, 0x9e, 0x45, 0x44, 0x9f, 0x67, 0x9d,
	0xa1, 0x1f, 0x00, 0xb1, 0x09, 0x58, 0xa9, 0x08, 0x7e, 0x73, 0x60, 0xd3, 0xb4, 0xc8, 0x0a, 0x5c,
	0xde, 0xd4, 0x3f, 0x69, 0x87, 0xb8, 0x0b, 0x1d, 0x82, 0x74, 0xe8, 0xbe, 0x41, 0x22, 0x9a, 0x50,
	0x51, 0x12, 0x28, 0x1e, 0x4a, 0x4d, 0x77, 0xb7, 0x76, 0x50, 0xa8, 0x3b, 0x5e, 0x66, 0xa4, 0x3f,
	0x38, 0x50, 0xcf, 0x92, 0x5d, 0xe5, 0xbb, 0x65, 0x2d, 0xf6, 0xe3, 0x00, 0xb3, 0xd0, 0x15, 0x6d,
	0x30, 0x79, 0x08, 0x6b, 0xfd, 0x38, 0x50, 0x69, 0xb8, 0x4a, 0x8e, 0x0d, 0x19, 0xac, 0x2b, 0xe2,
	0x51, 0x78, 0xf1, 0x4c, 0x78, 0xc6, 0x49, 0x7f, 0x77, 0x60, 0x2b, 0x1b, 0xe5, 0x2b, 0xb0, 0xb7,
	0x03, 0xeb, 0xad, 0x60, 0xc4, 0x43, 0xa1, 0x8a, 0x51, 0x35, 0x86, 0xc1, 0x69, 0x0b, 0xe8, 0xb1,
	0xaf, 0xa6, 0x69, 0x7e, 0xda, 0x16, 0x5f, 0x38, 0x6d, 0x4b, 0xf9, 0x69, 0x2b, 0x95, 0xc2, 0x52,
	0xd2, 0x33, 0x48, 0x23, 0xfa, 0x8b, 0x03, 0xc4, 0xfe, 0x98, 0x95, 0xd8, 0x5d, 0xde, 0x5a, 0xb9,
	0xdb, 0xc8, 0x5d, 0xb8, 0x8d, 0x94, 0x22, 0x09, 0xf7, 0xbb, 0xa3, 0xcf, 0xb1, 0xdf, 0x8b, 0x5e,
	0x8a, 0xe9, 0xbb, 0x50, 0x6b, 0xf3, 0xc0, 0xe4, 0x96, 0x4c, 0x48, 0x1d, 0x5c, 0x8f, 0x0b, 0x7d,
	0x69, 0xca, 0xc7, 0xe5, 0xaf, 0xa4, 0x5f, 0x39, 0xb0, 0x7d, 0xc0, 0xc4, 0xe0, 0xd2, 0xfa, 0xf9,
	0x4a, 0x13, 0xc3, 0xae, 0x98, 0x65, 0x75, 0xeb, 0xde, 0x50, 0xb7, 0xaf, 0x5d, 0xcb, 0x62, 0x55,
	0x82, 0x97, 0x64, 0xf2, 0x00, 0x8a, 0x56, 0xd5, 0x6e, 0xc9, 0x18, 0x39, 0xda, 0x3c, 0xe5, 0xa6,
	0xbf, 0x3a, 0x2a, 0x63, 0x3d, 0xaf, 0x73, 0x1b, 0x48, 0xaa, 0xd5, 0x36, 0x94, 0xbb, 0x82, 0x89,
	0x69, 0xa2, 0x77, 0x10, 0x8d, 0xac, 0xf2, 0x71, 0xed, 0xf2, 0x49, 0xb7, 0x93, 0xe2, 0x8d, 0xdb,
	0x49, 0xe9, 0xc5, 0xdb, 0x49, 0x79, 0x71, 0x3b, 0x49, 0xa5, 0x7b, 0x8f, 0x8b, 0xff, 0x4f, 0xba,
	0x9f, 0x8c, 0x74, 0x76, 0x16, 0xb7, 0x20, 0xdd, 0xa3, 0xdc, 0x2d, 0x80, 0x02, 0xd6, 0x64, 0xa4,
	0xe5, 0x57, 0xc0, 0x9f, 0x0e, 0xac, 0xab, 0x55, 0xa0, 0xd5, 0x3e, 0x95, 0x9a, 0xb4, 0x99, 0x60,
	0xd9, 0xc2, 0x84, 0x48, 0xad, 0x3b, 0xd3, 0x20, 0x38, 0xf1, 0x9f, 0xe8, 0xa9, 0x62, 0xa0, 0x54,
	0xe6, 0x7d, 0xce, 0xfc, 0xc7, 0x4f, 0xde, 0x91, 0x4e, 0x54, 0xd2, 0xb2, 0xc8, 0x5e, 0x94, 0x47,
	0x75, 0x2f, 0xaa, 0xe9, 0x68, 0x30, 0x79, 0x08, 0x77, 0x64, 0x7c, 0x6b, 0x00, 0xa1, 0xb2, 0x0b,
	0x56, 0xb9, 0xd0, 0x64, 0x96, 0x23, 0x5f, 0x8f, 0x9b, 0x9c, 0x4d, 0x66, 0x78, 0x94, 0xf4, 0x2e,
	0xa7, 0xe3, 0xf3, 0xc6, 0x1a, 0x2e, 0x64, 0x1a, 0xd2, 0x9f, 0x0b, 0x00, 0x1d, 0x36, 0x6b, 0xb5,
	0x4f, 0xd5, 0x9a, 0xfc, 0xdf, 0x7e, 0xe2, 0x36, 0x94, 0x5b, 0xed, 0xd3, 0xbe, 0x77, 0xac, 0xa7,
	0xa7, 0x46, 0x6a, 0x03, 0x3d, 0xec, 0xca, 0x0d, 0x54, 0xcf, 0x4e, 0x44, 0x32, 0xd6, 0x61, 0x38,
	0x88, 0xe7, 0x67, 0x3c, 0x56, 0xdf, 0x51, 0xf2, 0x52, 0x6c, 0x6f, 0xa7, 0xeb, 0xb9, 0xed, 0xd4,
	0x6a, 0xb1, 0x4a, 0xae, 0xc5, 0xee, 0x01, 0x28, 0x92, 0xd0, 0x07, 0xca, 0x67, 0x59, 0x68, 0x07,
	0xb6, 0x74, 0x31, 0x86, 0xc3, 0xe8, 0x9f, 0xb7, 0x04, 0xfd, 0xb6, 0x00, 0xc4, 0x0e, 0xb4, 0xf2,
	0x7d, 0x6a, 0xa6, 0xb7, 0x2c, 0x6c, 0x6b, 0x7a, 0xcb, 0x3e, 0xef, 0x45, 0x82, 0x21, 0xd7, 0xae,
	0x72, 0x66, 0x06, 0x72, 0x1f, 0x6a, 0x27, 0x57, 0x1d, 0x36, 0x93, 0x5b, 0xbe, 0xa5, 0x46, 0xde,
	0x48, 0xde, 0x80, 0xfa, 0xc9, 0xd5, 0xfe, 0x54, 0x44, 0xfd, 0x49, 0x10, 0x31, 0x7c, 0x0f, 0xd6,
	0xdd, 0x35, 0x3b, 0xd9, 0x03, 0x82, 0xb6, 0x76, 0xf4, 0x59, 0x98, 0x9e, 0xc6, 0x01, 0xb3, 0xc4,
	0x43, 0xbf, 0x80, 0x4d, 0xd3, 0x4b, 0xb7, 0x30, 0x61, 0x8a, 0x8b, 0x8b, 0x44, 0xfa, 0x2e, 0x9c,
	0xc6, 0x5f, 0x3a, 0x50, 0xcf, 0x5e, 0x7f, 0x0b, 0xa3, 0x85, 0xe6, 0x52, 0xb8, 0xa3, 0x87, 0x8a,
	0xee, 0x2c, 0x9d, 0x84, 0x07, 0x77, 0x3a, 0x6c, 0xd6, 0x9d, 0x87, 0x03, 0x6b, 0x27, 0xe9, 0xf2,
	0x80, 0x0f, 0x44, 0x14, 0x9b, 0xbf, 0x03, 0x06, 0x93, 0xfb, 0x50, 0x7e, 0xca, 0xe7, 0x07, 0xd3,
	0xa1, 0x52, 0x53, 0x7f, 0xd6, 0xc1, 0x74, 0x38, 0xe4, 0xf1, 0x33, 0xe1, 0x69, 0x9f, 0x1c, 0x9d,
	0x9b, 0x69, 0x50, 0x9d, 0xe1, 0xf5, 0x7b, 0xfb, 0x01, 0xac, 0xb5, 0xc6, 0x38, 0x80, 0x0b, 0x2a,
	0x58, 0x55, 0x71, 0x84, 0x26, 0xcf, 0xf8, 0x5e, 0xee, 0x95, 0x72, 0xe6, 0xb4, 0xa2, 0x50, 0x8c,
	0xc2, 0x29, 0xb7, 0x6e, 0xa3, 0x9c, 0x8d, 0xfe, 0xe8, 0xc0, 0x46, 0x2f, 0x66, 0x61, 0x62, 0x66,
	0x8b, 0xbd, 0x7d, 0xe9, 0x2f, 0x4d, 0xb7, 0xaf, 0xec, 0xbf, 0x68, 0xe1, 0x86, 0xff, 0xa2, 0x6e,
	0xee, 0xbf, 0x68, 0x13, 0xaa, 0x1d, 0x36, 0x93, 0xcd, 0x6a, 0xad, 0x66, 0xb6, 0x49, 0x5e, 0x94,
	0x69, 0xf1, 0x16, 0x3d, 0xf5, 0x4c, 0xbf, 0x71, 0xe0, 0xae, 0xba, 0x64, 0x4c, 0x5e, 0xff, 0x7a,
	0x19, 0xde, 0xcf, 0xd5, 0x80, 0xaa, 0x23, 0x9b, 0x03, 0x5d, 0x05, 0xdf, 0x3b, 0xf0, 0xea, 0x42,
	0x1e, 0xb7, 0x50, 0x8f, 0x2f, 0x95, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x7a, 0x0d,
	0x7a, 0x50, 0x11, 0x00, 0x00,
}
