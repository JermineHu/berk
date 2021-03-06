// Code generated by protoc-gen-go.
// source: pupbatch.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PupBatchOperationRequest struct {
	SockpuppetAccountIDs []int64  `protobuf:"varint,1,rep,packed,name=SockpuppetAccountIDs,json=sockpuppetAccountIDs" json:"SockpuppetAccountIDs,omitempty"`
	SockpuppetAccountAll bool     `protobuf:"varint,2,opt,name=SockpuppetAccountAll,json=sockpuppetAccountAll" json:"SockpuppetAccountAll,omitempty"`
	SockpuppetAccountPwd string   `protobuf:"bytes,3,opt,name=SockpuppetAccountPwd,json=sockpuppetAccountPwd" json:"SockpuppetAccountPwd,omitempty"`
	UserAccountIDs       []int64  `protobuf:"varint,4,rep,packed,name=UserAccountIDs,json=userAccountIDs" json:"UserAccountIDs,omitempty"`
	NewsUUIDs            []string `protobuf:"bytes,5,rep,name=NewsUUIDs,json=newsUUIDs" json:"NewsUUIDs,omitempty"`
	Tags                 []string `protobuf:"bytes,6,rep,name=Tags,json=tags" json:"Tags,omitempty"`
}

func (m *PupBatchOperationRequest) Reset()                    { *m = PupBatchOperationRequest{} }
func (m *PupBatchOperationRequest) String() string            { return proto1.CompactTextString(m) }
func (*PupBatchOperationRequest) ProtoMessage()               {}
func (*PupBatchOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func init() {
	proto1.RegisterType((*PupBatchOperationRequest)(nil), "proto.PupBatchOperationRequest")
}

func init() { proto1.RegisterFile("pupbatch.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0x2d, 0x48,
	0x4a, 0x2c, 0x49, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x13,
	0x98, 0xb8, 0x24, 0x02, 0x4a, 0x0b, 0x9c, 0x40, 0x32, 0xfe, 0x05, 0xa9, 0x45, 0x89, 0x25, 0x99,
	0xf9, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x46, 0x5c, 0x22, 0xc1, 0xf9, 0xc9,
	0xd9, 0x05, 0xa5, 0x05, 0x05, 0xa9, 0x25, 0x8e, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x9e, 0x2e,
	0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xcc, 0x41, 0x22, 0xc5, 0x58, 0xe4, 0xb0, 0xea, 0x71, 0xcc,
	0xc9, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0xc0, 0xa2, 0xc7, 0x31, 0x27, 0x07, 0xab, 0x9e, 0x80,
	0xf2, 0x14, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x4e, 0x2c, 0x7a, 0x02, 0xca, 0x53, 0x84, 0xd4, 0xb8,
	0xf8, 0x42, 0x8b, 0x53, 0x8b, 0x90, 0x5c, 0xc5, 0x02, 0x76, 0x15, 0x5f, 0x29, 0x8a, 0xa8, 0x90,
	0x0c, 0x17, 0xa7, 0x5f, 0x6a, 0x79, 0x71, 0x68, 0x28, 0x48, 0x09, 0xab, 0x02, 0xb3, 0x06, 0x67,
	0x10, 0x67, 0x1e, 0x4c, 0x40, 0x48, 0x88, 0x8b, 0x25, 0x24, 0x31, 0xbd, 0x58, 0x82, 0x0d, 0x2c,
	0xc1, 0x52, 0x92, 0x98, 0x5e, 0x9c, 0xc4, 0x06, 0x0e, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xd0, 0x39, 0xa7, 0x32, 0x01, 0x00, 0x00,
}
