// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorName           string   `protobuf:"bytes,2,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	LikeCount            uint32   `protobuf:"varint,5,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetLikeCount() uint32 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

type PostId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostId) Reset()         { *m = PostId{} }
func (m *PostId) String() string { return proto.CompactTextString(m) }
func (*PostId) ProtoMessage()    {}
func (*PostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}

func (m *PostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostId.Unmarshal(m, b)
}
func (m *PostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostId.Marshal(b, m, deterministic)
}
func (m *PostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostId.Merge(m, src)
}
func (m *PostId) XXX_Size() int {
	return xxx_messageInfo_PostId.Size(m)
}
func (m *PostId) XXX_DiscardUnknown() {
	xxx_messageInfo_PostId.DiscardUnknown(m)
}

var xxx_messageInfo_PostId proto.InternalMessageInfo

func (m *PostId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "proto.Post")
	proto.RegisterType((*PostId)(nil), "proto.PostId")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xb1, 0xad, 0x76, 0x4a, 0x3d, 0x2c, 0x22, 0x4b, 0x45, 0x2c, 0x3d, 0x48, 0x2f, 0x4d,
	0xb4, 0xc5, 0x1f, 0x68, 0xf5, 0x50, 0x10, 0x29, 0x91, 0x5e, 0xbc, 0x94, 0x34, 0x3b, 0xa6, 0x8b,
	0xc9, 0x6e, 0x48, 0x26, 0x05, 0xfd, 0x00, 0xbf, 0xc7, 0x4f, 0x94, 0xdd, 0xb4, 0xd0, 0x16, 0x0b,
	0x9e, 0x86, 0xf7, 0x66, 0xde, 0xcc, 0xdb, 0xc7, 0x02, 0x64, 0xba, 0x20, 0x2f, 0xcb, 0x35, 0x69,
	0x56, 0xb7, 0xa5, 0x73, 0x15, 0x6b, 0x1d, 0x27, 0xe8, 0x5b, 0xb4, 0x2c, 0xdf, 0x7d, 0x4c, 0x33,
	0xfa, 0xac, 0x66, 0x7a, 0xdf, 0x0e, 0xd4, 0x66, 0xba, 0x20, 0x76, 0x0e, 0xae, 0x14, 0xdc, 0xe9,
	0x3a, 0xfd, 0x66, 0xe0, 0x4a, 0xc1, 0x6e, 0xa0, 0x15, 0x96, 0xb4, 0xd2, 0xf9, 0x42, 0x85, 0x29,
	0x72, 0xd7, 0x36, 0xa0, 0xa2, 0x5e, 0xc2, 0x14, 0xd9, 0x05, 0xd4, 0x49, 0x52, 0x82, 0xfc, 0xc4,
	0xb6, 0x2a, 0xc0, 0x38, 0x9c, 0x46, 0x5a, 0x11, 0x2a, 0xe2, 0x35, 0xcb, 0x6f, 0x21, 0xbb, 0x06,
	0x48, 0xe4, 0x07, 0x2e, 0x22, 0x5d, 0x2a, 0xe2, 0xf5, 0xae, 0xd3, 0x6f, 0x07, 0x4d, 0xc3, 0x4c,
	0x0c, 0xd1, 0xe3, 0xd0, 0x30, 0x3e, 0xa6, 0xe2, 0xd0, 0xc9, 0xf0, 0xc7, 0x85, 0xd6, 0x38, 0xd1,
	0xf1, 0x2b, 0xe6, 0x6b, 0x19, 0x21, 0xeb, 0x03, 0x4c, 0x72, 0x0c, 0x09, 0x0d, 0xc9, 0x5a, 0xd5,
	0x43, 0x3c, 0x23, 0xee, 0xb4, 0x77, 0xc0, 0x54, 0xb0, 0x5b, 0x38, 0x0b, 0x30, 0x14, 0x76, 0x6e,
	0xbf, 0xd5, 0xd9, 0x95, 0xb1, 0x7b, 0x80, 0x79, 0x26, 0xfe, 0xdc, 0x78, 0xe9, 0x55, 0xe9, 0x79,
	0xdb, 0xf4, 0xbc, 0x27, 0x93, 0x1e, 0x1b, 0x01, 0x3c, 0x62, 0x82, 0x1b, 0xc9, 0xc1, 0xf2, 0x63,
	0xa2, 0x21, 0x34, 0x9f, 0x65, 0x41, 0x46, 0x52, 0xb0, 0x23, 0x43, 0x7b, 0xce, 0xee, 0x1c, 0x73,
	0x68, 0x9e, 0xad, 0x35, 0xa1, 0x75, 0xfa, 0xbf, 0x43, 0xe3, 0x87, 0xb7, 0x51, 0x2c, 0x69, 0x55,
	0x2e, 0xbd, 0x48, 0xa7, 0xbe, 0xc2, 0x50, 0x09, 0xcc, 0x53, 0x54, 0xf8, 0x85, 0x85, 0x1f, 0x07,
	0xb3, 0xc9, 0xa0, 0xda, 0x39, 0xd8, 0x64, 0xba, 0xf9, 0x1b, 0x0d, 0x5b, 0x46, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0xc4, 0x96, 0x1c, 0x45, 0x02, 0x00, 0x00,
}
