// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Gender int32

const (
	Gender_user_gender_null   Gender = 0
	Gender_user_gender_female Gender = 1
	Gender_user_gender_male   Gender = 2
)

var Gender_name = map[int32]string{
	0: "user_gender_null",
	1: "user_gender_female",
	2: "user_gender_male",
}
var Gender_value = map[string]int32{
	"user_gender_null":   0,
	"user_gender_female": 1,
	"user_gender_male":   2,
}

func (x Gender) String() string {
	return proto1.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

type Status int32

const (
	Status_USER_STATUS_NULL        Status = 0
	Status_USER_STATUS_ACTIVATED   Status = 1
	Status_USER_STATUS_DEACTIVATED Status = 2
	Status_USER_STATUS_CLOSED      Status = 3
)

var Status_name = map[int32]string{
	0: "USER_STATUS_NULL",
	1: "USER_STATUS_ACTIVATED",
	2: "USER_STATUS_DEACTIVATED",
	3: "USER_STATUS_CLOSED",
}
var Status_value = map[string]int32{
	"USER_STATUS_NULL":        0,
	"USER_STATUS_ACTIVATED":   1,
	"USER_STATUS_DEACTIVATED": 2,
	"USER_STATUS_CLOSED":      3,
}

func (x Status) String() string {
	return proto1.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

type UserType int32

const (
	UserType_NULL          UserType = 0
	UserType_NORMAL        UserType = 1
	UserType_PUP           UserType = 2
	UserType_OPERATIONS    UserType = 3
	UserType_ADMINISTRATOR UserType = 4
	UserType_MERCHANT      UserType = 5
)

var UserType_name = map[int32]string{
	0: "NULL",
	1: "NORMAL",
	2: "PUP",
	3: "OPERATIONS",
	4: "ADMINISTRATOR",
	5: "MERCHANT",
}
var UserType_value = map[string]int32{
	"NULL":          0,
	"NORMAL":        1,
	"PUP":           2,
	"OPERATIONS":    3,
	"ADMINISTRATOR": 4,
	"MERCHANT":      5,
}

func (x UserType) String() string {
	return proto1.EnumName(UserType_name, int32(x))
}
func (UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

type UserRelatedType int32

const (
	UserRelatedType_related_type_null          UserRelatedType = 0
	UserRelatedType_related_type_none          UserRelatedType = 1
	UserRelatedType_related_type_followee      UserRelatedType = 2
	UserRelatedType_related_type_follower      UserRelatedType = 3
	UserRelatedType_related_type_mutual_follow UserRelatedType = 4
	UserRelatedType_related_type_friend        UserRelatedType = 5
	UserRelatedType_related_type_Self          UserRelatedType = 100
)

var UserRelatedType_name = map[int32]string{
	0:   "related_type_null",
	1:   "related_type_none",
	2:   "related_type_followee",
	3:   "related_type_follower",
	4:   "related_type_mutual_follow",
	5:   "related_type_friend",
	100: "related_type_Self",
}
var UserRelatedType_value = map[string]int32{
	"related_type_null":          0,
	"related_type_none":          1,
	"related_type_followee":      2,
	"related_type_follower":      3,
	"related_type_mutual_follow": 4,
	"related_type_friend":        5,
	"related_type_Self":          100,
}

func (x UserRelatedType) String() string {
	return proto1.EnumName(UserRelatedType_name, int32(x))
}
func (UserRelatedType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

type User struct {
	ID             int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID           string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt      string       `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	TusoID         *StringType  `protobuf:"bytes,4,opt,name=TusoID,json=tusoID" json:"TusoID,omitempty"`
	Email          *StringType  `protobuf:"bytes,5,opt,name=Email,json=email" json:"Email,omitempty"`
	MobileNumber   *StringType  `protobuf:"bytes,6,opt,name=MobileNumber,json=mobileNumber" json:"MobileNumber,omitempty"`
	Password       *StringType  `protobuf:"bytes,7,opt,name=Password,json=password" json:"Password,omitempty"`
	Birthday       string       `protobuf:"bytes,8,opt,name=Birthday,json=birthday" json:"Birthday,omitempty"`
	Salt           *StringType  `protobuf:"bytes,9,opt,name=Salt,json=salt" json:"Salt,omitempty"`
	Token          *StringType  `protobuf:"bytes,10,opt,name=Token,json=token" json:"Token,omitempty"`
	Nickname       *StringType  `protobuf:"bytes,11,opt,name=Nickname,json=nickname" json:"Nickname,omitempty"`
	RealName       *StringType  `protobuf:"bytes,12,opt,name=RealName,json=realName" json:"RealName,omitempty"`
	Gender         Gender       `protobuf:"varint,13,opt,name=Gender,json=gender,enum=proto.Gender" json:"Gender,omitempty"`
	Location       *StringType  `protobuf:"bytes,14,opt,name=Location,json=location" json:"Location,omitempty"`
	FolloweesCount *IntegerType `protobuf:"bytes,15,opt,name=FolloweesCount,json=followeesCount" json:"FolloweesCount,omitempty"`
	FollowersCount *IntegerType `protobuf:"bytes,16,opt,name=FollowersCount,json=followersCount" json:"FollowersCount,omitempty"`
	FriendsCount   *IntegerType `protobuf:"bytes,17,opt,name=FriendsCount,json=friendsCount" json:"FriendsCount,omitempty"`
	ImagesCount    *IntegerType `protobuf:"bytes,18,opt,name=ImagesCount,json=imagesCount" json:"ImagesCount,omitempty"`
	TusosCount     *IntegerType `protobuf:"bytes,19,opt,name=TusosCount,json=tusosCount" json:"TusosCount,omitempty"`
	NuclearKey     *StringType  `protobuf:"bytes,20,opt,name=NuclearKey,json=nuclearKey" json:"NuclearKey,omitempty"`
	Secrets        *StringType  `protobuf:"bytes,21,opt,name=Secrets,json=secrets" json:"Secrets,omitempty"`
	Status         Status       `protobuf:"varint,22,opt,name=Status,json=status,enum=proto.Status" json:"Status,omitempty"`
	Type           UserType     `protobuf:"varint,23,opt,name=Type,json=type,enum=proto.UserType" json:"Type,omitempty"`
	FirstPhoto     string       `protobuf:"bytes,25,opt,name=FirstPhoto,json=firstPhoto" json:"FirstPhoto,omitempty"`
	FirstTuso      string       `protobuf:"bytes,26,opt,name=FirstTuso,json=firstTuso" json:"FirstTuso,omitempty"`
	UpdatedAt      string       `protobuf:"bytes,27,opt,name=UpdatedAt,json=updatedAt" json:"UpdatedAt,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *User) GetTusoID() *StringType {
	if m != nil {
		return m.TusoID
	}
	return nil
}

func (m *User) GetEmail() *StringType {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetMobileNumber() *StringType {
	if m != nil {
		return m.MobileNumber
	}
	return nil
}

func (m *User) GetPassword() *StringType {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *User) GetSalt() *StringType {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *User) GetToken() *StringType {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *User) GetNickname() *StringType {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *User) GetRealName() *StringType {
	if m != nil {
		return m.RealName
	}
	return nil
}

func (m *User) GetLocation() *StringType {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *User) GetFolloweesCount() *IntegerType {
	if m != nil {
		return m.FolloweesCount
	}
	return nil
}

func (m *User) GetFollowersCount() *IntegerType {
	if m != nil {
		return m.FollowersCount
	}
	return nil
}

func (m *User) GetFriendsCount() *IntegerType {
	if m != nil {
		return m.FriendsCount
	}
	return nil
}

func (m *User) GetImagesCount() *IntegerType {
	if m != nil {
		return m.ImagesCount
	}
	return nil
}

func (m *User) GetTusosCount() *IntegerType {
	if m != nil {
		return m.TusosCount
	}
	return nil
}

func (m *User) GetNuclearKey() *StringType {
	if m != nil {
		return m.NuclearKey
	}
	return nil
}

func (m *User) GetSecrets() *StringType {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type UserRelation struct {
	ID              int64           `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID            string          `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	FromID          *IntegerType    `protobuf:"bytes,3,opt,name=FromID,json=fromID" json:"FromID,omitempty"`
	ToID            *IntegerType    `protobuf:"bytes,4,opt,name=ToID,json=toID" json:"ToID,omitempty"`
	RelatedType     UserRelatedType `protobuf:"varint,5,opt,name=RelatedType,json=relatedType,enum=proto.UserRelatedType" json:"RelatedType,omitempty"`
	ApplyingFriends *BooleanType    `protobuf:"bytes,6,opt,name=ApplyingFriends,json=applyingFriends" json:"ApplyingFriends,omitempty"`
	ApplyAt         string          `protobuf:"bytes,7,opt,name=ApplyAt,json=applyAt" json:"ApplyAt,omitempty"`
	Remark          *StringType     `protobuf:"bytes,8,opt,name=Remark,json=remark" json:"Remark,omitempty"`
	ToUser          *User           `protobuf:"bytes,9,opt,name=ToUser,json=toUser" json:"ToUser,omitempty"`
	FromUser        *User           `protobuf:"bytes,10,opt,name=FromUser,json=fromUser" json:"FromUser,omitempty"`
}

func (m *UserRelation) Reset()                    { *m = UserRelation{} }
func (m *UserRelation) String() string            { return proto1.CompactTextString(m) }
func (*UserRelation) ProtoMessage()               {}
func (*UserRelation) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *UserRelation) GetFromID() *IntegerType {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *UserRelation) GetToID() *IntegerType {
	if m != nil {
		return m.ToID
	}
	return nil
}

func (m *UserRelation) GetApplyingFriends() *BooleanType {
	if m != nil {
		return m.ApplyingFriends
	}
	return nil
}

func (m *UserRelation) GetRemark() *StringType {
	if m != nil {
		return m.Remark
	}
	return nil
}

func (m *UserRelation) GetToUser() *User {
	if m != nil {
		return m.ToUser
	}
	return nil
}

func (m *UserRelation) GetFromUser() *User {
	if m != nil {
		return m.FromUser
	}
	return nil
}

type FolloweeUser struct {
	FromUserID int64 `protobuf:"varint,1,opt,name=FromUserID,json=fromUserID" json:"FromUserID,omitempty"`
	ToUser     *User `protobuf:"bytes,2,opt,name=ToUser,json=toUser" json:"ToUser,omitempty"`
	IsFollow   bool  `protobuf:"varint,3,opt,name=IsFollow,json=isFollow" json:"IsFollow,omitempty"`
}

func (m *FolloweeUser) Reset()                    { *m = FolloweeUser{} }
func (m *FolloweeUser) String() string            { return proto1.CompactTextString(m) }
func (*FolloweeUser) ProtoMessage()               {}
func (*FolloweeUser) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *FolloweeUser) GetToUser() *User {
	if m != nil {
		return m.ToUser
	}
	return nil
}

func init() {
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*UserRelation)(nil), "proto.UserRelation")
	proto1.RegisterType((*FolloweeUser)(nil), "proto.FolloweeUser")
	proto1.RegisterEnum("proto.Gender", Gender_name, Gender_value)
	proto1.RegisterEnum("proto.Status", Status_name, Status_value)
	proto1.RegisterEnum("proto.UserType", UserType_name, UserType_value)
	proto1.RegisterEnum("proto.UserRelatedType", UserRelatedType_name, UserRelatedType_value)
}

func init() { proto1.RegisterFile("user.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x95, 0x6f, 0x6f, 0xe2, 0x46,
	0x10, 0xc6, 0x0f, 0x30, 0xc6, 0x19, 0x38, 0xe2, 0xcc, 0x5d, 0x92, 0x4d, 0xae, 0x3a, 0x45, 0x39,
	0x5d, 0x2f, 0xa5, 0xea, 0x49, 0x4d, 0xff, 0xa8, 0xaa, 0xfa, 0xc6, 0x87, 0x49, 0x6b, 0x95, 0x00,
	0xb2, 0x4d, 0x5f, 0xf4, 0x0d, 0x72, 0x60, 0xe1, 0xac, 0x18, 0x1b, 0xad, 0xd7, 0x3a, 0xf1, 0x31,
	0xfa, 0x9d, 0xda, 0xef, 0x55, 0xed, 0x7a, 0x4d, 0xcc, 0x25, 0x56, 0xfb, 0xca, 0xf2, 0x3c, 0xbf,
	0x67, 0x77, 0x3d, 0x3b, 0x33, 0x06, 0xc8, 0x52, 0xca, 0xde, 0x6f, 0x58, 0xc2, 0x13, 0x6c, 0xca,
	0xc7, 0x79, 0x37, 0xce, 0xa2, 0x88, 0x6f, 0x37, 0x34, 0x0f, 0x5f, 0xfe, 0x6d, 0x80, 0x36, 0x4d,
	0x29, 0xc3, 0x2e, 0xd4, 0x1d, 0x9b, 0xd4, 0x2e, 0x6a, 0x57, 0x0d, 0xb7, 0x1e, 0xda, 0x88, 0xa0,
	0x4d, 0xa7, 0x8e, 0x4d, 0xea, 0x17, 0xb5, 0xab, 0x03, 0x57, 0xcb, 0xa6, 0x8e, 0x8d, 0x5f, 0xc0,
	0x41, 0x9f, 0xd1, 0x80, 0xd3, 0x85, 0xc5, 0x49, 0x43, 0x0a, 0x07, 0xf3, 0x22, 0x80, 0x5f, 0x81,
	0xee, 0x67, 0x69, 0xe2, 0xd8, 0x44, 0xbb, 0xa8, 0x5d, 0xb5, 0xaf, 0x8f, 0xf2, 0x2d, 0xde, 0x7b,
	0x9c, 0x85, 0xf1, 0xca, 0xdf, 0x6e, 0xa8, 0xab, 0x73, 0x09, 0xe0, 0x3b, 0x68, 0x0e, 0xd6, 0x41,
	0x18, 0x91, 0x66, 0x15, 0xd9, 0xa4, 0x42, 0xc7, 0x1f, 0xa0, 0x73, 0x9b, 0xdc, 0x85, 0x11, 0x1d,
	0x65, 0xeb, 0x3b, 0xca, 0x88, 0x5e, 0xc5, 0x77, 0xd6, 0x25, 0x0c, 0xbf, 0x01, 0x63, 0x12, 0xa4,
	0xe9, 0xa7, 0x84, 0x2d, 0x48, 0xab, 0xca, 0x62, 0x6c, 0x14, 0x82, 0xe7, 0x60, 0x7c, 0x08, 0x19,
	0xff, 0xb8, 0x08, 0xb6, 0xc4, 0x90, 0x9f, 0x65, 0xdc, 0xa9, 0x77, 0x7c, 0x0b, 0x9a, 0x17, 0x44,
	0x9c, 0x1c, 0x54, 0x2d, 0xa3, 0xa5, 0x41, 0xc4, 0xc5, 0x17, 0xf9, 0xc9, 0x3d, 0x8d, 0x09, 0x54,
	0x7e, 0x11, 0x17, 0xba, 0x38, 0xda, 0x28, 0x9c, 0xdf, 0xc7, 0xc1, 0x9a, 0x92, 0x76, 0xe5, 0xd1,
	0x62, 0x85, 0x08, 0xdc, 0xa5, 0x41, 0x34, 0x12, 0x78, 0xa7, 0x12, 0x67, 0x0a, 0xc1, 0xb7, 0xa0,
	0xff, 0x4a, 0xe3, 0x05, 0x65, 0xe4, 0xf9, 0x45, 0xed, 0xaa, 0x7b, 0xfd, 0x5c, 0xc1, 0x79, 0xd0,
	0xd5, 0x57, 0xf2, 0x29, 0x56, 0x1d, 0x26, 0xf3, 0x80, 0x87, 0x49, 0x4c, 0xba, 0x95, 0xab, 0x46,
	0x0a, 0xc1, 0x9f, 0xa1, 0x7b, 0x93, 0x44, 0x51, 0xf2, 0x89, 0xd2, 0xb4, 0x9f, 0x64, 0x31, 0x27,
	0x87, 0xd2, 0x84, 0xca, 0xe4, 0xc4, 0x9c, 0xae, 0x28, 0x93, 0xae, 0xee, 0x72, 0x8f, 0x2c, 0x79,
	0x99, 0xf2, 0x9a, 0xff, 0xe9, 0x55, 0x24, 0xfe, 0x08, 0x9d, 0x1b, 0x16, 0xd2, 0x78, 0xa1, 0x9c,
	0x47, 0x95, 0xce, 0xce, 0xb2, 0xc4, 0xe1, 0xf7, 0xd0, 0x76, 0xd6, 0xc1, 0xaa, 0x38, 0x2c, 0x56,
	0xda, 0xda, 0xe1, 0x03, 0x86, 0xd7, 0x00, 0xa2, 0x7e, 0x95, 0xe9, 0x45, 0xa5, 0x09, 0xf8, 0x8e,
	0xc2, 0x6f, 0x01, 0x46, 0xd9, 0x3c, 0xa2, 0x01, 0xfb, 0x9d, 0x6e, 0xc9, 0xcb, 0xaa, 0x54, 0x42,
	0xbc, 0x83, 0xf0, 0x6b, 0x68, 0x79, 0x74, 0xce, 0x28, 0x4f, 0xc9, 0x71, 0x15, 0xdf, 0x4a, 0x73,
	0x42, 0xdc, 0xa7, 0xc7, 0x03, 0x9e, 0xa5, 0xe4, 0x64, 0xef, 0x3e, 0xf3, 0xa0, 0xab, 0xa7, 0xf2,
	0x89, 0x6f, 0x40, 0x13, 0x3e, 0x72, 0x2a, 0xa1, 0x43, 0x05, 0x89, 0xbe, 0xce, 0x4b, 0x54, 0x34,
	0x3c, 0xbe, 0x06, 0xb8, 0x09, 0x59, 0xca, 0x27, 0x1f, 0x13, 0x9e, 0x90, 0x33, 0x59, 0xe7, 0xb0,
	0xdc, 0x45, 0x44, 0x77, 0x4b, 0x5d, 0x24, 0x81, 0x9c, 0xe7, 0xdd, 0xbd, 0x2c, 0x02, 0x42, 0x9d,
	0x6e, 0x16, 0xaa, 0xf7, 0x5f, 0xe5, 0x6a, 0x56, 0x04, 0x2e, 0xff, 0x6a, 0x40, 0x47, 0x6c, 0xe7,
	0xd2, 0x28, 0x2f, 0x99, 0xff, 0x33, 0x4e, 0x7a, 0xa0, 0xdf, 0xb0, 0x64, 0xed, 0xd8, 0x72, 0x96,
	0x3c, 0x9d, 0x6c, 0x7d, 0x29, 0x09, 0xfc, 0x12, 0x34, 0xff, 0x61, 0xb4, 0x3c, 0x45, 0x6a, 0x5c,
	0x4c, 0x96, 0x9f, 0xa0, 0x2d, 0xcf, 0x40, 0x17, 0x32, 0x21, 0x4d, 0x99, 0x90, 0x93, 0x52, 0x42,
	0x4a, 0xaa, 0xdb, 0x66, 0x0f, 0x2f, 0xf8, 0x0b, 0x1c, 0x5a, 0x9b, 0x4d, 0xb4, 0x0d, 0xe3, 0x95,
	0x2a, 0x3a, 0x35, 0x6d, 0x8a, 0xcd, 0x3e, 0x24, 0x49, 0x44, 0x83, 0x58, 0x3a, 0x0f, 0x83, 0x7d,
	0x14, 0x09, 0xb4, 0xa4, 0xdb, 0xe2, 0x72, 0xe0, 0x1c, 0xb8, 0xad, 0x20, 0x7f, 0x15, 0x63, 0xd1,
	0xa5, 0xeb, 0x80, 0xdd, 0xcb, 0xd1, 0xf2, 0xf4, 0x58, 0x64, 0x12, 0xc0, 0x37, 0xa0, 0xfb, 0x89,
	0x38, 0xa4, 0x9a, 0x36, 0xed, 0xf2, 0xb9, 0x75, 0x2e, 0x25, 0x7c, 0x07, 0x86, 0xc8, 0x9a, 0xc4,
	0xe0, 0x31, 0x66, 0x2c, 0x95, 0x78, 0x99, 0x40, 0xa7, 0xe8, 0x5a, 0x69, 0x14, 0xf7, 0xaf, 0xb4,
	0xdd, 0xd5, 0xc0, 0x72, 0x17, 0x29, 0xed, 0x5e, 0xaf, 0xde, 0xfd, 0x1c, 0x0c, 0x27, 0xcd, 0x97,
	0x95, 0xb7, 0x66, 0xb8, 0x46, 0xa8, 0xde, 0x7b, 0xc3, 0x62, 0xf8, 0xe0, 0x4b, 0x30, 0xc5, 0xaf,
	0x67, 0x96, 0x8f, 0x9b, 0x99, 0xf8, 0xe7, 0x98, 0xcf, 0xf0, 0x04, 0xb0, 0x1c, 0x5d, 0xd2, 0x75,
	0x10, 0x51, 0xb3, 0xf6, 0x39, 0x2d, 0xa3, 0xf5, 0x5e, 0x5c, 0x94, 0xbe, 0xd0, 0xa7, 0xde, 0xc0,
	0x9d, 0x79, 0xbe, 0xe5, 0x4f, 0xbd, 0xd9, 0x68, 0x3a, 0x1c, 0x9a, 0xcf, 0xf0, 0x0c, 0x8e, 0xcb,
	0x51, 0xab, 0xef, 0x3b, 0x7f, 0x58, 0xfe, 0xc0, 0x36, 0x6b, 0xf8, 0x0a, 0x4e, 0xcb, 0x92, 0x3d,
	0x78, 0x10, 0xeb, 0xe2, 0x14, 0x65, 0xb1, 0x3f, 0x1c, 0x7b, 0x03, 0xdb, 0x6c, 0xf4, 0xfe, 0x04,
	0xa3, 0x68, 0x18, 0x34, 0x40, 0x53, 0xbb, 0x00, 0xe8, 0xa3, 0xb1, 0x7b, 0x6b, 0x0d, 0xcd, 0x1a,
	0xb6, 0xa0, 0x31, 0x99, 0x4e, 0xcc, 0x3a, 0x76, 0x01, 0xc6, 0x93, 0x81, 0x6b, 0xf9, 0xce, 0x78,
	0xe4, 0x99, 0x0d, 0x3c, 0x82, 0xe7, 0x96, 0x7d, 0xeb, 0x8c, 0x1c, 0xcf, 0x77, 0x2d, 0x7f, 0xec,
	0x9a, 0x1a, 0x76, 0xc0, 0xb8, 0x1d, 0xb8, 0xfd, 0xdf, 0xac, 0x91, 0x6f, 0x36, 0x7b, 0xff, 0xd4,
	0xe0, 0xf0, 0xb3, 0xe2, 0xc3, 0x63, 0x38, 0x52, 0xe5, 0x37, 0x13, 0xed, 0x59, 0x24, 0xe9, 0x51,
	0x38, 0x89, 0x45, 0x8e, 0xce, 0xe0, 0x78, 0x2f, 0x5c, 0x4c, 0x59, 0xb3, 0x5e, 0x25, 0x31, 0xb3,
	0x81, 0xaf, 0xe1, 0x7c, 0x4f, 0x5a, 0x67, 0x3c, 0x0b, 0x22, 0x45, 0x98, 0x1a, 0x9e, 0xc2, 0x8b,
	0x7d, 0xab, 0xac, 0x66, 0xb3, 0xf9, 0xe8, 0x14, 0x1e, 0x8d, 0x96, 0xe6, 0xe2, 0x4e, 0x97, 0x15,
	0xf1, 0xdd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x36, 0xa3, 0xed, 0x59, 0x08, 0x00, 0x00,
}
