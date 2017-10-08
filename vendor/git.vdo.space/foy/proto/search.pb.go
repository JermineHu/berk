// Code generated by protoc-gen-go.
// source: search.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Followee struct {
	UUID           string `protobuf:"bytes,1,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	BelongUserUUID string `protobuf:"bytes,2,opt,name=BelongUserUUID,json=belongUserUUID" json:"BelongUserUUID,omitempty"`
	FUserModel     *User  `protobuf:"bytes,3,opt,name=FUserModel,json=fUserModel" json:"FUserModel,omitempty"`
}

func (m *Followee) Reset()                    { *m = Followee{} }
func (m *Followee) String() string            { return proto1.CompactTextString(m) }
func (*Followee) ProtoMessage()               {}
func (*Followee) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Followee) GetFUserModel() *User {
	if m != nil {
		return m.FUserModel
	}
	return nil
}

type Follower struct {
	UUID           string `protobuf:"bytes,1,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	BelongUserUUID string `protobuf:"bytes,2,opt,name=BelongUserUUID,json=belongUserUUID" json:"BelongUserUUID,omitempty"`
	FUserModel     *User  `protobuf:"bytes,3,opt,name=FUserModel,json=fUserModel" json:"FUserModel,omitempty"`
}

func (m *Follower) Reset()                    { *m = Follower{} }
func (m *Follower) String() string            { return proto1.CompactTextString(m) }
func (*Follower) ProtoMessage()               {}
func (*Follower) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *Follower) GetFUserModel() *User {
	if m != nil {
		return m.FUserModel
	}
	return nil
}

type FollowUser struct {
	Followee *Followee `protobuf:"bytes,1,opt,name=Followee,json=followee" json:"Followee,omitempty"`
	Follower *Follower `protobuf:"bytes,2,opt,name=Follower,json=follower" json:"Follower,omitempty"`
}

func (m *FollowUser) Reset()                    { *m = FollowUser{} }
func (m *FollowUser) String() string            { return proto1.CompactTextString(m) }
func (*FollowUser) ProtoMessage()               {}
func (*FollowUser) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *FollowUser) GetFollowee() *Followee {
	if m != nil {
		return m.Followee
	}
	return nil
}

func (m *FollowUser) GetFollower() *Follower {
	if m != nil {
		return m.Follower
	}
	return nil
}

type FollowUUIDS struct {
	FolloweeUUID string `protobuf:"bytes,1,opt,name=FolloweeUUID,json=followeeUUID" json:"FolloweeUUID,omitempty"`
	FollowerUUID string `protobuf:"bytes,2,opt,name=FollowerUUID,json=followerUUID" json:"FollowerUUID,omitempty"`
}

func (m *FollowUUIDS) Reset()                    { *m = FollowUUIDS{} }
func (m *FollowUUIDS) String() string            { return proto1.CompactTextString(m) }
func (*FollowUUIDS) ProtoMessage()               {}
func (*FollowUUIDS) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

type Location struct {
	Country  string `protobuf:"bytes,1,opt,name=Country,json=country" json:"Country,omitempty"`
	State    string `protobuf:"bytes,2,opt,name=State,json=state" json:"State,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=City,json=city" json:"City,omitempty"`
	District string `protobuf:"bytes,4,opt,name=District,json=district" json:"District,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto1.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func init() {
	proto1.RegisterType((*Followee)(nil), "proto.Followee")
	proto1.RegisterType((*Follower)(nil), "proto.Follower")
	proto1.RegisterType((*FollowUser)(nil), "proto.FollowUser")
	proto1.RegisterType((*FollowUUIDS)(nil), "proto.FollowUUIDS")
	proto1.RegisterType((*Location)(nil), "proto.Location")
}

func init() { proto1.RegisterFile("search.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x90, 0x41, 0x6b, 0x83, 0x40,
	0x10, 0x85, 0xb1, 0xd5, 0xc6, 0x8c, 0x92, 0xc2, 0xd0, 0x83, 0xe4, 0x14, 0x3c, 0x94, 0x40, 0x20,
	0x87, 0xf4, 0x1f, 0x34, 0x12, 0x10, 0xda, 0x8b, 0x61, 0x7f, 0x80, 0x31, 0x6b, 0x6b, 0x11, 0xb7,
	0xcc, 0xae, 0x14, 0xff, 0x7d, 0xd9, 0x75, 0x4d, 0xb6, 0xf4, 0x9e, 0x93, 0xbc, 0x6f, 0x9e, 0xf3,
	0xf6, 0x0d, 0xc4, 0x92, 0x97, 0x54, 0x7d, 0x6e, 0xbf, 0x49, 0x28, 0x81, 0x81, 0xf9, 0x2c, 0xa1,
	0x97, 0x9c, 0x46, 0x94, 0x4a, 0x08, 0x0f, 0xa2, 0x6d, 0xc5, 0x0f, 0xe7, 0x88, 0xe0, 0x33, 0x96,
	0x67, 0x89, 0xb7, 0xf2, 0xd6, 0xf3, 0xc2, 0xef, 0x59, 0x9e, 0xe1, 0x33, 0x2c, 0x5e, 0x79, 0x2b,
	0xba, 0x0f, 0x26, 0x39, 0x99, 0xe9, 0x9d, 0x99, 0x2e, 0x4e, 0x7f, 0x28, 0x6e, 0x00, 0x0e, 0x5a,
	0xbc, 0x8b, 0x33, 0x6f, 0x93, 0xfb, 0x95, 0xb7, 0x8e, 0x76, 0xd1, 0x98, 0xb1, 0xd5, 0xbc, 0x80,
	0xfa, 0x32, 0x76, 0x42, 0xe9, 0x76, 0xa1, 0x35, 0xc0, 0x18, 0xaa, 0x11, 0x6e, 0xae, 0xbd, 0x4d,
	0x74, 0xb4, 0x7b, 0xb4, 0x3f, 0x4e, 0xb8, 0x08, 0xeb, 0xe9, 0x30, 0x57, 0x33, 0x99, 0x97, 0xfc,
	0x33, 0xd3, 0xc5, 0x4c, 0x29, 0x83, 0xc8, 0xe6, 0xb0, 0x3c, 0x3b, 0x62, 0x0a, 0xf1, 0xb4, 0xd1,
	0xe9, 0x19, 0xd7, 0x0e, 0x73, 0x3c, 0x6e, 0xdb, 0xc9, 0x63, 0x58, 0xfa, 0x05, 0xe1, 0x9b, 0xa8,
	0x4a, 0xd5, 0x88, 0x0e, 0x13, 0x98, 0xed, 0x45, 0xdf, 0x29, 0x1a, 0xec, 0xba, 0x59, 0x35, 0x4a,
	0x7c, 0x82, 0xe0, 0xa8, 0x4a, 0xc5, 0xed, 0x8a, 0x40, 0x6a, 0xa1, 0x6f, 0xbc, 0x6f, 0xd4, 0x60,
	0x2e, 0x34, 0x2f, 0xfc, 0xaa, 0x51, 0x03, 0x2e, 0x21, 0xcc, 0x1a, 0xa9, 0xa8, 0xa9, 0x54, 0xe2,
	0x1b, 0x1e, 0x9e, 0xad, 0x3e, 0x3d, 0x98, 0x72, 0x2f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07,
	0xba, 0xcc, 0xff, 0x3e, 0x02, 0x00, 0x00,
}