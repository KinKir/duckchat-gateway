// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_friend_delete.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

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

// *
//
// action: api.friend.delete
//
type ApiFriendDeleteRequest struct {
	ToUserId             string   `protobuf:"bytes,1,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiFriendDeleteRequest) Reset()         { *m = ApiFriendDeleteRequest{} }
func (m *ApiFriendDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ApiFriendDeleteRequest) ProtoMessage()    {}
func (*ApiFriendDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_friend_delete_0dabc450153c23f8, []int{0}
}
func (m *ApiFriendDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFriendDeleteRequest.Unmarshal(m, b)
}
func (m *ApiFriendDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFriendDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *ApiFriendDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFriendDeleteRequest.Merge(dst, src)
}
func (m *ApiFriendDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ApiFriendDeleteRequest.Size(m)
}
func (m *ApiFriendDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFriendDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFriendDeleteRequest proto.InternalMessageInfo

func (m *ApiFriendDeleteRequest) GetToUserId() string {
	if m != nil {
		return m.ToUserId
	}
	return ""
}

type ApiFriendDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiFriendDeleteResponse) Reset()         { *m = ApiFriendDeleteResponse{} }
func (m *ApiFriendDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ApiFriendDeleteResponse) ProtoMessage()    {}
func (*ApiFriendDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_friend_delete_0dabc450153c23f8, []int{1}
}
func (m *ApiFriendDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFriendDeleteResponse.Unmarshal(m, b)
}
func (m *ApiFriendDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFriendDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *ApiFriendDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFriendDeleteResponse.Merge(dst, src)
}
func (m *ApiFriendDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ApiFriendDeleteResponse.Size(m)
}
func (m *ApiFriendDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFriendDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFriendDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApiFriendDeleteRequest)(nil), "site.ApiFriendDeleteRequest")
	proto.RegisterType((*ApiFriendDeleteResponse)(nil), "site.ApiFriendDeleteResponse")
}

func init() {
	proto.RegisterFile("site/api_friend_delete.proto", fileDescriptor_api_friend_delete_0dabc450153c23f8)
}

var fileDescriptor_api_friend_delete_0dabc450153c23f8 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xce, 0x2c, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0x8c, 0x4f, 0x2b, 0xca, 0x4c, 0xcd, 0x4b, 0x89, 0x4f, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9, 0x2a, 0x99, 0x70, 0x89,
	0x39, 0x16, 0x64, 0xba, 0x81, 0xe5, 0x5d, 0xc0, 0xd2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x52, 0x5c, 0x1c, 0x25, 0xf9, 0xa1, 0xc5, 0xa9, 0x45, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x92, 0x24, 0x97, 0x38, 0x86, 0xae, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0xa7, 0x08, 0x2e, 0xe1, 0xe4, 0xfc, 0x5c, 0xbd, 0xaa, 0xc4, 0x9c, 0x4a, 0x88, 0x45, 0x7a,
	0x20, 0x7b, 0xa2, 0xf4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53,
	0x4a, 0x93, 0xb3, 0x93, 0x33, 0x12, 0x4b, 0xe0, 0x0c, 0xdd, 0xf4, 0xc4, 0x92, 0xd4, 0xf2, 0xc4,
	0x4a, 0x7d, 0xb0, 0x06, 0x7d, 0x90, 0x86, 0x53, 0x4c, 0xfc, 0x51, 0x89, 0x39, 0x95, 0x31, 0x01,
	0x20, 0x91, 0x98, 0xe0, 0xcc, 0x92, 0xd4, 0x24, 0x36, 0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x80, 0x2f, 0x2c, 0xd7, 0x00, 0x00, 0x00,
}
