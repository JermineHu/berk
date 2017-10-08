// Code generated by protoc-gen-go.
// source: nulltype.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StringType struct {
	String_ string `protobuf:"bytes,1,opt,name=String,json=string" json:"String,omitempty"`
	Null    bool   `protobuf:"varint,2,opt,name=Null,json=null" json:"Null,omitempty"`
}

func (m *StringType) Reset()                    { *m = StringType{} }
func (m *StringType) String() string            { return proto1.CompactTextString(m) }
func (*StringType) ProtoMessage()               {}
func (*StringType) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type IntegerType struct {
	Int  int64 `protobuf:"varint,1,opt,name=Int,json=int" json:"Int,omitempty"`
	Null bool  `protobuf:"varint,2,opt,name=Null,json=null" json:"Null,omitempty"`
}

func (m *IntegerType) Reset()                    { *m = IntegerType{} }
func (m *IntegerType) String() string            { return proto1.CompactTextString(m) }
func (*IntegerType) ProtoMessage()               {}
func (*IntegerType) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type BooleanType struct {
	Bool bool `protobuf:"varint,1,opt,name=Bool,json=bool" json:"Bool,omitempty"`
	Null bool `protobuf:"varint,2,opt,name=Null,json=null" json:"Null,omitempty"`
}

func (m *BooleanType) Reset()                    { *m = BooleanType{} }
func (m *BooleanType) String() string            { return proto1.CompactTextString(m) }
func (*BooleanType) ProtoMessage()               {}
func (*BooleanType) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type FloatType struct {
	Float float32 `protobuf:"fixed32,1,opt,name=Float,json=float" json:"Float,omitempty"`
	Null  bool    `protobuf:"varint,2,opt,name=Null,json=null" json:"Null,omitempty"`
}

func (m *FloatType) Reset()                    { *m = FloatType{} }
func (m *FloatType) String() string            { return proto1.CompactTextString(m) }
func (*FloatType) ProtoMessage()               {}
func (*FloatType) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func init() {
	proto1.RegisterType((*StringType)(nil), "proto.StringType")
	proto1.RegisterType((*IntegerType)(nil), "proto.IntegerType")
	proto1.RegisterType((*BooleanType)(nil), "proto.BooleanType")
	proto1.RegisterType((*FloatType)(nil), "proto.FloatType")
}

func init() { proto1.RegisterFile("nulltype.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0xcf, 0x31, 0xcb, 0x83, 0x40,
	0x0c, 0xc6, 0x71, 0xd4, 0x53, 0x34, 0xc2, 0xcb, 0xcb, 0x51, 0x8a, 0xa3, 0x38, 0x39, 0x75, 0x91,
	0x42, 0xe7, 0x0e, 0x05, 0x97, 0x0e, 0xd7, 0x7e, 0x01, 0x85, 0xab, 0x08, 0x21, 0x11, 0x1b, 0x07,
	0xbf, 0x7d, 0x31, 0x5d, 0xed, 0x14, 0xfe, 0xc3, 0x2f, 0xf0, 0xc0, 0x1f, 0x2d, 0x88, 0xb2, 0x4e,
	0xfe, 0x34, 0xcd, 0x2c, 0x6c, 0x63, 0x3d, 0xd5, 0x05, 0xe0, 0x21, 0xf3, 0x48, 0xc3, 0x73, 0x9d,
	0xbc, 0x3d, 0x42, 0xf2, 0xad, 0x22, 0x28, 0x83, 0x3a, 0x73, 0xc9, 0x5b, 0xcb, 0x5a, 0x30, 0xf7,
	0x05, 0xb1, 0x08, 0xcb, 0xa0, 0x4e, 0x9d, 0xd9, 0x5e, 0x55, 0x0d, 0xe4, 0x2d, 0x89, 0x1f, 0xfc,
	0xac, 0xf4, 0x1f, 0xa2, 0x96, 0x44, 0x5d, 0xe4, 0xa2, 0x91, 0x64, 0x17, 0x9d, 0x21, 0xbf, 0x32,
	0xa3, 0xef, 0x48, 0x91, 0x05, 0xb3, 0xa5, 0xaa, 0xd4, 0x99, 0x9e, 0x19, 0x7f, 0xb0, 0xec, 0x86,
	0xdc, 0x89, 0xa2, 0x03, 0xc4, 0x1a, 0xaa, 0x42, 0x17, 0xbf, 0xb6, 0xd8, 0x63, 0x7d, 0xa2, 0x1b,
	0x9b, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x31, 0x95, 0xf9, 0xfc, 0x00, 0x00, 0x00,
}
