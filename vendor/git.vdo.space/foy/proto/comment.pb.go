// Code generated by protoc-gen-go.
// source: comment.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommentType int32

const (
	CommentType_comment_type_null  CommentType = 0
	CommentType_comment_type_news  CommentType = 1
	CommentType_comment_type_image CommentType = 2
	CommentType_comment_type_diary CommentType = 3
)

var CommentType_name = map[int32]string{
	0: "comment_type_null",
	1: "comment_type_news",
	2: "comment_type_image",
	3: "comment_type_diary",
}
var CommentType_value = map[string]int32{
	"comment_type_null":  0,
	"comment_type_news":  1,
	"comment_type_image": 2,
	"comment_type_diary": 3,
}

func (x CommentType) String() string {
	return proto1.EnumName(CommentType_name, int32(x))
}
func (CommentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Comment struct {
	ID          int64       `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID        string      `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	SourceId    int64       `protobuf:"varint,3,opt,name=SourceId,json=sourceId" json:"SourceId,omitempty"`
	User        *User       `protobuf:"bytes,4,opt,name=User,json=user" json:"User,omitempty"`
	UserUUID    string      `protobuf:"bytes,5,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	ReplyTo     *Comment    `protobuf:"bytes,6,opt,name=ReplyTo,json=replyTo" json:"ReplyTo,omitempty"`
	ReplyToUUID string      `protobuf:"bytes,7,opt,name=ReplyToUUID,json=replyToUUID" json:"ReplyToUUID,omitempty"`
	Content     string      `protobuf:"bytes,8,opt,name=Content,json=content" json:"Content,omitempty"`
	Type        CommentType `protobuf:"varint,9,opt,name=Type,json=type,enum=proto.CommentType" json:"Type,omitempty"`
	Timestamp   string      `protobuf:"bytes,10,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto1.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Comment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Comment) GetReplyTo() *Comment {
	if m != nil {
		return m.ReplyTo
	}
	return nil
}

type NewsComments struct {
	Data []*NewsComment `protobuf:"bytes,1,rep,name=Data,json=data" json:"Data,omitempty"`
}

func (m *NewsComments) Reset()                    { *m = NewsComments{} }
func (m *NewsComments) String() string            { return proto1.CompactTextString(m) }
func (*NewsComments) ProtoMessage()               {}
func (*NewsComments) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *NewsComments) GetData() []*NewsComment {
	if m != nil {
		return m.Data
	}
	return nil
}

type NewsComment struct {
	ID        int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID      string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt string       `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	UpdatedAt string       `protobuf:"bytes,4,opt,name=UpdatedAt,json=updatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt string       `protobuf:"bytes,5,opt,name=DeletedAt,json=deletedAt" json:"DeletedAt,omitempty"`
	SourceId  int64        `protobuf:"varint,6,opt,name=SourceId,json=sourceId" json:"SourceId,omitempty"`
	User      *User        `protobuf:"bytes,7,opt,name=User,json=user" json:"User,omitempty"`
	UserID    int64        `protobuf:"varint,8,opt,name=UserID,json=userID" json:"UserID,omitempty"`
	ReplyTo   *NewsComment `protobuf:"bytes,9,opt,name=ReplyTo,json=replyTo" json:"ReplyTo,omitempty"`
	ReplyToID int64        `protobuf:"varint,10,opt,name=ReplyToID,json=replyToID" json:"ReplyToID,omitempty"`
	Content   string       `protobuf:"bytes,11,opt,name=Content,json=content" json:"Content,omitempty"`
	Type      CommentType  `protobuf:"varint,12,opt,name=Type,json=type,enum=proto.CommentType" json:"Type,omitempty"`
	Timestamp string       `protobuf:"bytes,13,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
}

func (m *NewsComment) Reset()                    { *m = NewsComment{} }
func (m *NewsComment) String() string            { return proto1.CompactTextString(m) }
func (*NewsComment) ProtoMessage()               {}
func (*NewsComment) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *NewsComment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *NewsComment) GetReplyTo() *NewsComment {
	if m != nil {
		return m.ReplyTo
	}
	return nil
}

func init() {
	proto1.RegisterType((*Comment)(nil), "proto.Comment")
	proto1.RegisterType((*NewsComments)(nil), "proto.NewsComments")
	proto1.RegisterType((*NewsComment)(nil), "proto.NewsComment")
	proto1.RegisterEnum("proto.CommentType", CommentType_name, CommentType_value)
}

func init() { proto1.RegisterFile("comment.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xa4, 0x8d, 0xe4, 0x1d, 0x25, 0xc6, 0x1d, 0x68, 0x58, 0x82, 0xa1, 0xc2, 0x87, 0x22,
	0x4a, 0xc9, 0x21, 0x85, 0xde, 0x4b, 0x74, 0xd1, 0xa5, 0x87, 0xad, 0x75, 0x0e, 0xaa, 0x35, 0x14,
	0x53, 0xcb, 0x12, 0xab, 0x15, 0xc1, 0xff, 0xa1, 0xbf, 0xa6, 0xbf, 0xb0, 0xec, 0x87, 0x12, 0xbb,
	0x49, 0xa0, 0x39, 0x99, 0x79, 0x6f, 0xde, 0x9b, 0xf5, 0x7b, 0x08, 0x2e, 0x36, 0x5d, 0xdb, 0xd2,
	0x5e, 0x5f, 0xf7, 0xaa, 0xd3, 0x1d, 0x9e, 0xd9, 0x9f, 0x2b, 0x18, 0x07, 0x52, 0x0e, 0x5a, 0xfd,
	0x09, 0x21, 0xb9, 0x75, 0x4b, 0x38, 0x87, 0xb0, 0x2c, 0x44, 0x90, 0x05, 0x79, 0x24, 0xc3, 0x6d,
	0x81, 0x08, 0xac, 0xaa, 0xca, 0x42, 0x84, 0x59, 0x90, 0x73, 0xc9, 0xc6, 0xaa, 0x2c, 0xf0, 0x0a,
	0x66, 0xdf, 0xbb, 0x51, 0x6d, 0xa8, 0x6c, 0x44, 0x64, 0x37, 0x67, 0x83, 0x9f, 0xf1, 0x3d, 0xb0,
	0x6a, 0x20, 0x25, 0x58, 0x16, 0xe4, 0xe9, 0x4d, 0xea, 0x2e, 0x5c, 0x1b, 0x48, 0x32, 0x73, 0xd2,
	0x88, 0xcd, 0x64, 0x4d, 0xcf, 0xac, 0xe9, 0x6c, 0xf4, 0x33, 0xe6, 0x90, 0x48, 0xea, 0x77, 0x87,
	0x75, 0x27, 0x62, 0xab, 0x9f, 0x7b, 0xbd, 0x7f, 0x9d, 0x4c, 0x94, 0xa3, 0x31, 0x83, 0xd4, 0x6f,
	0x5a, 0xa3, 0xc4, 0x1a, 0xa5, 0xea, 0x11, 0x42, 0x61, 0xfe, 0xd3, 0x5e, 0xd3, 0x5e, 0x8b, 0x99,
	0x65, 0x93, 0x8d, 0x1b, 0xf1, 0x03, 0xb0, 0xf5, 0xa1, 0x27, 0xc1, 0xb3, 0x20, 0x9f, 0xdf, 0xe0,
	0xe9, 0x09, 0xc3, 0x48, 0xa6, 0x0f, 0x3d, 0xe1, 0x12, 0xf8, 0x7a, 0xdb, 0xd2, 0xa0, 0xeb, 0xb6,
	0x17, 0x60, 0x3d, 0xb8, 0x9e, 0x80, 0xd5, 0x17, 0x38, 0xff, 0x46, 0xf7, 0x83, 0x97, 0x0d, 0xc6,
	0xb5, 0xa8, 0x75, 0x2d, 0x82, 0x2c, 0xca, 0xd3, 0x07, 0xd7, 0xa3, 0x15, 0xc9, 0x9a, 0x5a, 0xd7,
	0xab, 0xdf, 0x11, 0xa4, 0x47, 0xe8, 0x7f, 0x05, 0xbe, 0x04, 0x7e, 0xab, 0xa8, 0xd6, 0xd4, 0x7c,
	0xd5, 0x36, 0x71, 0x2e, 0xf9, 0x66, 0x02, 0x0c, 0x5b, 0xf5, 0x8d, 0x67, 0x99, 0x63, 0xc7, 0x09,
	0x30, 0x6c, 0x41, 0x3b, 0x72, 0xac, 0x0b, 0x9c, 0x37, 0x13, 0x70, 0x52, 0x65, 0xfc, 0x42, 0x95,
	0xc9, 0x4b, 0x55, 0x5e, 0x42, 0x6c, 0xa6, 0xb2, 0xb0, 0x09, 0x47, 0x32, 0x1e, 0xed, 0x84, 0x9f,
	0x1e, 0x6b, 0xe4, 0x56, 0xfb, 0x5c, 0x1a, 0x0f, 0x55, 0x2e, 0x81, 0xfb, 0xed, 0xb2, 0xb0, 0x31,
	0x47, 0x92, 0xab, 0x09, 0x38, 0xae, 0x31, 0x7d, 0xbe, 0xc6, 0xf3, 0xd7, 0xd4, 0x78, 0xf1, 0x4f,
	0x8d, 0x1f, 0x7f, 0x41, 0x7a, 0x24, 0xc1, 0x77, 0xf0, 0xd6, 0x7f, 0x2e, 0x77, 0x46, 0x7c, 0xb7,
	0x1f, 0x77, 0xbb, 0xc5, 0x9b, 0xa7, 0x30, 0xdd, 0x0f, 0x8b, 0x00, 0x2f, 0x01, 0x4f, 0xe0, 0x6d,
	0x5b, 0xff, 0xa4, 0x45, 0xf8, 0x04, 0x6f, 0xb6, 0xb5, 0x3a, 0x2c, 0xa2, 0x1f, 0xb1, 0x7d, 0xe3,
	0xe7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x08, 0x16, 0x00, 0x93, 0x03, 0x00, 0x00,
}
