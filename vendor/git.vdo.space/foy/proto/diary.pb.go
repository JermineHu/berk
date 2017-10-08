// Code generated by protoc-gen-go.
// source: diary.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DiaryPrivacy int32

const (
	DiaryPrivacy_diary_privacy_null    DiaryPrivacy = 0
	DiaryPrivacy_diary_privacy_public  DiaryPrivacy = 1
	DiaryPrivacy_diary_privacy_private DiaryPrivacy = 2
)

var DiaryPrivacy_name = map[int32]string{
	0: "diary_privacy_null",
	1: "diary_privacy_public",
	2: "diary_privacy_private",
}
var DiaryPrivacy_value = map[string]int32{
	"diary_privacy_null":    0,
	"diary_privacy_public":  1,
	"diary_privacy_private": 2,
}

func (x DiaryPrivacy) String() string {
	return proto1.EnumName(DiaryPrivacy_name, int32(x))
}
func (DiaryPrivacy) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type DiaryStatus int32

const (
	DiaryStatus_diary_status_none      DiaryStatus = 0
	DiaryStatus_diary_status_null      DiaryStatus = 1
	DiaryStatus_diary_status_saved     DiaryStatus = 2
	DiaryStatus_diary_status_published DiaryStatus = 3
)

var DiaryStatus_name = map[int32]string{
	0: "diary_status_none",
	1: "diary_status_null",
	2: "diary_status_saved",
	3: "diary_status_published",
}
var DiaryStatus_value = map[string]int32{
	"diary_status_none":      0,
	"diary_status_null":      1,
	"diary_status_saved":     2,
	"diary_status_published": 3,
}

func (x DiaryStatus) String() string {
	return proto1.EnumName(DiaryStatus_name, int32(x))
}
func (DiaryStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type Diary struct {
	ID           int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID         string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt    string       `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	UserUUID     string       `protobuf:"bytes,4,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	Title        *StringType  `protobuf:"bytes,5,opt,name=Title,json=title" json:"Title,omitempty"`
	Content      *StringType  `protobuf:"bytes,6,opt,name=Content,json=content" json:"Content,omitempty"`
	Style        *StringType  `protobuf:"bytes,7,opt,name=Style,json=style" json:"Style,omitempty"`
	DiaryPrivacy DiaryPrivacy `protobuf:"varint,8,opt,name=DiaryPrivacy,json=diaryPrivacy,enum=proto.DiaryPrivacy" json:"DiaryPrivacy,omitempty"`
	DiaryStatus  DiaryStatus  `protobuf:"varint,9,opt,name=DiaryStatus,json=diaryStatus,enum=proto.DiaryStatus" json:"DiaryStatus,omitempty"`
	Timestamp    string       `protobuf:"bytes,10,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
}

func (m *Diary) Reset()                    { *m = Diary{} }
func (m *Diary) String() string            { return proto1.CompactTextString(m) }
func (*Diary) ProtoMessage()               {}
func (*Diary) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Diary) GetTitle() *StringType {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *Diary) GetContent() *StringType {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Diary) GetStyle() *StringType {
	if m != nil {
		return m.Style
	}
	return nil
}

func init() {
	proto1.RegisterType((*Diary)(nil), "proto.Diary")
	proto1.RegisterEnum("proto.DiaryPrivacy", DiaryPrivacy_name, DiaryPrivacy_value)
	proto1.RegisterEnum("proto.DiaryStatus", DiaryStatus_name, DiaryStatus_value)
}

func init() { proto1.RegisterFile("diary.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x6e, 0x82, 0x40,
	0x10, 0x86, 0x05, 0x45, 0x65, 0x34, 0x06, 0xa7, 0xd5, 0x6c, 0x4d, 0x0f, 0xa4, 0x97, 0x12, 0x9b,
	0x78, 0xb0, 0x4d, 0x7a, 0x6e, 0xe4, 0xe2, 0xad, 0x41, 0x39, 0xf5, 0x60, 0x10, 0x36, 0x2d, 0x09,
	0x02, 0x81, 0xc1, 0x84, 0xc7, 0xee, 0x1b, 0x34, 0xbb, 0xab, 0xa9, 0xb4, 0xf1, 0xb4, 0x99, 0xef,
	0xfb, 0x67, 0x77, 0x18, 0x60, 0x10, 0xc5, 0x41, 0x51, 0x2f, 0xf2, 0x22, 0xa3, 0x0c, 0x0d, 0x79,
	0xcc, 0x46, 0x69, 0x95, 0x24, 0x54, 0xe7, 0x5c, 0xe1, 0x87, 0x6f, 0x1d, 0x0c, 0x57, 0xc4, 0x70,
	0x04, 0xfa, 0xda, 0x65, 0x9a, 0xad, 0x39, 0x6d, 0x4f, 0x8f, 0x5d, 0x44, 0xe8, 0xf8, 0xfe, 0xda,
	0x65, 0xba, 0xad, 0x39, 0xa6, 0xd7, 0xa9, 0xfc, 0xb5, 0x8b, 0xf7, 0x60, 0xae, 0x0a, 0x1e, 0x10,
	0x8f, 0xde, 0x88, 0xb5, 0xa5, 0x30, 0xc3, 0x33, 0xc0, 0x19, 0xf4, 0xfd, 0x92, 0x17, 0xb2, 0xab,
	0x23, 0x65, 0xbf, 0x3a, 0xd5, 0xf8, 0x08, 0xc6, 0x36, 0xa6, 0x84, 0x33, 0xc3, 0xd6, 0x9c, 0xc1,
	0x72, 0xac, 0x9e, 0x5f, 0x6c, 0xa8, 0x88, 0xd3, 0xcf, 0x6d, 0x9d, 0x73, 0xcf, 0x20, 0xe1, 0xf1,
	0x09, 0x7a, 0xab, 0x2c, 0x25, 0x9e, 0x12, 0xeb, 0x5e, 0x8b, 0xf6, 0x42, 0x95, 0x10, 0xb7, 0x6e,
	0xa8, 0x4e, 0x38, 0xeb, 0x5d, 0xbd, 0xb5, 0x14, 0x1e, 0x5f, 0x61, 0x28, 0xbf, 0xf2, 0xbd, 0x88,
	0x8f, 0x41, 0x58, 0xb3, 0xbe, 0xad, 0x39, 0xa3, 0xe5, 0xcd, 0x29, 0x7f, 0xa9, 0xbc, 0x61, 0x74,
	0x51, 0xe1, 0x0b, 0x0c, 0xa4, 0xdd, 0x50, 0x40, 0x55, 0xc9, 0x4c, 0xd9, 0x87, 0x97, 0x7d, 0xca,
	0x78, 0x6a, 0xd9, 0xaa, 0x10, 0x7b, 0xda, 0xc6, 0x07, 0x5e, 0x52, 0x70, 0xc8, 0x19, 0xa8, 0x3d,
	0xd1, 0x19, 0xcc, 0x3f, 0x9a, 0xc3, 0xe0, 0x14, 0x50, 0x36, 0xef, 0x72, 0x05, 0x76, 0xe2, 0x1f,
	0x59, 0x2d, 0x64, 0x70, 0xdb, 0xe4, 0x79, 0xb5, 0x4f, 0xe2, 0xd0, 0xd2, 0xf0, 0x0e, 0x26, 0x7f,
	0x8c, 0x38, 0x89, 0x5b, 0xfa, 0x3c, 0x6b, 0x0c, 0x8c, 0x13, 0x18, 0xab, 0x64, 0x29, 0xeb, 0x5d,
	0x9a, 0xa5, 0xdc, 0x6a, 0xfd, 0xc7, 0xe2, 0x45, 0xed, 0x77, 0x92, 0x13, 0x2e, 0x83, 0x23, 0x8f,
	0x2c, 0x1d, 0x67, 0x30, 0x6d, 0x70, 0x39, 0x48, 0xf9, 0xc5, 0x23, 0xab, 0xbd, 0xef, 0xca, 0x5d,
	0x3c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x31, 0x3c, 0xca, 0x83, 0x6e, 0x02, 0x00, 0x00,
}