// Code generated by protoc-gen-go.
// source: photo.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PhotoType int32

const (
	PhotoType_photo_type_null   PhotoType = 0
	PhotoType_photo_type_system PhotoType = 1
	PhotoType_photo_type_edited PhotoType = 2
)

var PhotoType_name = map[int32]string{
	0: "photo_type_null",
	1: "photo_type_system",
	2: "photo_type_edited",
}
var PhotoType_value = map[string]int32{
	"photo_type_null":   0,
	"photo_type_system": 1,
	"photo_type_edited": 2,
}

func (x PhotoType) String() string {
	return proto1.EnumName(PhotoType_name, int32(x))
}
func (PhotoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type PhotoPrivacy int32

const (
	PhotoPrivacy_photo_privacy_null    PhotoPrivacy = 0
	PhotoPrivacy_photo_privacy_public  PhotoPrivacy = 1
	PhotoPrivacy_photo_privacy_private PhotoPrivacy = 2
)

var PhotoPrivacy_name = map[int32]string{
	0: "photo_privacy_null",
	1: "photo_privacy_public",
	2: "photo_privacy_private",
}
var PhotoPrivacy_value = map[string]int32{
	"photo_privacy_null":    0,
	"photo_privacy_public":  1,
	"photo_privacy_private": 2,
}

func (x PhotoPrivacy) String() string {
	return proto1.EnumName(PhotoPrivacy_name, int32(x))
}
func (PhotoPrivacy) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

type PersistentType int32

const (
	PersistentType_persistent_type_null             PersistentType = 0
	PersistentType_persistent_type_new_photo        PersistentType = 1
	PersistentType_persistent_type_new_edited_photo PersistentType = 2
	PersistentType_persistent_type_edited_photo     PersistentType = 3
	PersistentType_persistent_type_edited_avatar    PersistentType = 4
)

var PersistentType_name = map[int32]string{
	0: "persistent_type_null",
	1: "persistent_type_new_photo",
	2: "persistent_type_new_edited_photo",
	3: "persistent_type_edited_photo",
	4: "persistent_type_edited_avatar",
}
var PersistentType_value = map[string]int32{
	"persistent_type_null":             0,
	"persistent_type_new_photo":        1,
	"persistent_type_new_edited_photo": 2,
	"persistent_type_edited_photo":     3,
	"persistent_type_edited_avatar":    4,
}

func (x PersistentType) String() string {
	return proto1.EnumName(PersistentType_name, int32(x))
}
func (PersistentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

type BasePhoto struct {
	ID           int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID         string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	Width        *IntegerType `protobuf:"bytes,3,opt,name=Width,json=width" json:"Width,omitempty"`
	Height       *IntegerType `protobuf:"bytes,4,opt,name=Height,json=height" json:"Height,omitempty"`
	FileURL      *StringType  `protobuf:"bytes,5,opt,name=FileURL,json=fileURL" json:"FileURL,omitempty"`
	FileUUID     string       `protobuf:"bytes,6,opt,name=FileUUID,json=fileUUID" json:"FileUUID,omitempty"`
	FileSize     *IntegerType `protobuf:"bytes,7,opt,name=FileSize,json=fileSize" json:"FileSize,omitempty"`
	PrimaryColor *StringType  `protobuf:"bytes,8,opt,name=PrimaryColor,json=primaryColor" json:"PrimaryColor,omitempty"`
}

func (m *BasePhoto) Reset()                    { *m = BasePhoto{} }
func (m *BasePhoto) String() string            { return proto1.CompactTextString(m) }
func (*BasePhoto) ProtoMessage()               {}
func (*BasePhoto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *BasePhoto) GetWidth() *IntegerType {
	if m != nil {
		return m.Width
	}
	return nil
}

func (m *BasePhoto) GetHeight() *IntegerType {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *BasePhoto) GetFileURL() *StringType {
	if m != nil {
		return m.FileURL
	}
	return nil
}

func (m *BasePhoto) GetFileSize() *IntegerType {
	if m != nil {
		return m.FileSize
	}
	return nil
}

func (m *BasePhoto) GetPrimaryColor() *StringType {
	if m != nil {
		return m.PrimaryColor
	}
	return nil
}

type Photo struct {
	ID             int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID           string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt      string       `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	User           *User        `protobuf:"bytes,4,opt,name=User,json=user" json:"User,omitempty"`
	UserUUID       string       `protobuf:"bytes,5,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	RawPhoto       *BasePhoto   `protobuf:"bytes,6,opt,name=RawPhoto,json=rawPhoto" json:"RawPhoto,omitempty"`
	Width          *IntegerType `protobuf:"bytes,7,opt,name=Width,json=width" json:"Width,omitempty"`
	Height         *IntegerType `protobuf:"bytes,8,opt,name=Height,json=height" json:"Height,omitempty"`
	InPipeline     *BooleanType `protobuf:"bytes,9,opt,name=InPipeline,json=inPipeline" json:"InPipeline,omitempty"`
	PhotoType      PhotoType    `protobuf:"varint,10,opt,name=PhotoType,json=photoType,enum=proto.PhotoType" json:"PhotoType,omitempty"`
	FileUUID       string       `protobuf:"bytes,11,opt,name=FileUUID,json=fileUUID" json:"FileUUID,omitempty"`
	FileSize       *IntegerType `protobuf:"bytes,12,opt,name=FileSize,json=fileSize" json:"FileSize,omitempty"`
	FileURL        *StringType  `protobuf:"bytes,13,opt,name=FileURL,json=fileURL" json:"FileURL,omitempty"`
	Identifier     *StringType  `protobuf:"bytes,14,opt,name=Identifier,json=identifier" json:"Identifier,omitempty"`
	PhotoPrivacy   PhotoPrivacy `protobuf:"varint,15,opt,name=PhotoPrivacy,json=photoPrivacy,enum=proto.PhotoPrivacy" json:"PhotoPrivacy,omitempty"`
	PrimaryColor   *StringType  `protobuf:"bytes,16,opt,name=PrimaryColor,json=primaryColor" json:"PrimaryColor,omitempty"`
	GEOLocation    *StringType  `protobuf:"bytes,17,opt,name=GEOLocation,json=gEOLocation" json:"GEOLocation,omitempty"`
	Note           *Note        `protobuf:"bytes,18,opt,name=Note,json=note" json:"Note,omitempty"`
	NoteUUID       string       `protobuf:"bytes,19,opt,name=NoteUUID,json=noteUUID" json:"NoteUUID,omitempty"`
	URL            *StringType  `protobuf:"bytes,20,opt,name=URL,json=uRL" json:"URL,omitempty"`
	DisplayVersion *IntegerType `protobuf:"bytes,21,opt,name=DisplayVersion,json=displayVersion" json:"DisplayVersion,omitempty"`
	EditParam      *StringType  `protobuf:"bytes,22,opt,name=EditParam,json=editParam" json:"EditParam,omitempty"`
	IsAvatar       *BooleanType `protobuf:"bytes,23,opt,name=IsAvatar,json=isAvatar" json:"IsAvatar,omitempty"`
	IsTuso         *BooleanType `protobuf:"bytes,24,opt,name=IsTuso,json=isTuso" json:"IsTuso,omitempty"`
	CommentsCount  *IntegerType `protobuf:"bytes,25,opt,name=CommentsCount,json=commentsCount" json:"CommentsCount,omitempty"`
	Timestamp      string       `protobuf:"bytes,26,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	Exif           *StringType  `protobuf:"bytes,27,opt,name=Exif,json=exif" json:"Exif,omitempty"`
	IsIllegal      *BooleanType `protobuf:"bytes,28,opt,name=IsIllegal,json=isIllegal" json:"IsIllegal,omitempty"`
}

func (m *Photo) Reset()                    { *m = Photo{} }
func (m *Photo) String() string            { return proto1.CompactTextString(m) }
func (*Photo) ProtoMessage()               {}
func (*Photo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *Photo) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Photo) GetRawPhoto() *BasePhoto {
	if m != nil {
		return m.RawPhoto
	}
	return nil
}

func (m *Photo) GetWidth() *IntegerType {
	if m != nil {
		return m.Width
	}
	return nil
}

func (m *Photo) GetHeight() *IntegerType {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *Photo) GetInPipeline() *BooleanType {
	if m != nil {
		return m.InPipeline
	}
	return nil
}

func (m *Photo) GetFileSize() *IntegerType {
	if m != nil {
		return m.FileSize
	}
	return nil
}

func (m *Photo) GetFileURL() *StringType {
	if m != nil {
		return m.FileURL
	}
	return nil
}

func (m *Photo) GetIdentifier() *StringType {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Photo) GetPrimaryColor() *StringType {
	if m != nil {
		return m.PrimaryColor
	}
	return nil
}

func (m *Photo) GetGEOLocation() *StringType {
	if m != nil {
		return m.GEOLocation
	}
	return nil
}

func (m *Photo) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *Photo) GetURL() *StringType {
	if m != nil {
		return m.URL
	}
	return nil
}

func (m *Photo) GetDisplayVersion() *IntegerType {
	if m != nil {
		return m.DisplayVersion
	}
	return nil
}

func (m *Photo) GetEditParam() *StringType {
	if m != nil {
		return m.EditParam
	}
	return nil
}

func (m *Photo) GetIsAvatar() *BooleanType {
	if m != nil {
		return m.IsAvatar
	}
	return nil
}

func (m *Photo) GetIsTuso() *BooleanType {
	if m != nil {
		return m.IsTuso
	}
	return nil
}

func (m *Photo) GetCommentsCount() *IntegerType {
	if m != nil {
		return m.CommentsCount
	}
	return nil
}

func (m *Photo) GetExif() *StringType {
	if m != nil {
		return m.Exif
	}
	return nil
}

func (m *Photo) GetIsIllegal() *BooleanType {
	if m != nil {
		return m.IsIllegal
	}
	return nil
}

type Note struct {
	ID        int64       `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID      string      `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt string      `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	UserUUID  string      `protobuf:"bytes,4,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	Title     *StringType `protobuf:"bytes,5,opt,name=Title,json=title" json:"Title,omitempty"`
	Content   *StringType `protobuf:"bytes,6,opt,name=Content,json=content" json:"Content,omitempty"`
	Style     *StringType `protobuf:"bytes,7,opt,name=Style,json=style" json:"Style,omitempty"`
	Timestamp string      `protobuf:"bytes,8,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto1.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *Note) GetTitle() *StringType {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *Note) GetContent() *StringType {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Note) GetStyle() *StringType {
	if m != nil {
		return m.Style
	}
	return nil
}

type File struct {
	ID             int64          `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID           string         `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt      string         `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	UserUUID       string         `protobuf:"bytes,4,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	Bucket         *StringType    `protobuf:"bytes,5,opt,name=Bucket,json=bucket" json:"Bucket,omitempty"`
	Key            *StringType    `protobuf:"bytes,6,opt,name=Key,json=key" json:"Key,omitempty"`
	Size           *IntegerType   `protobuf:"bytes,7,opt,name=Size,json=size" json:"Size,omitempty"`
	PersistentID   *StringType    `protobuf:"bytes,8,opt,name=PersistentID,json=persistentID" json:"PersistentID,omitempty"`
	PersistentType PersistentType `protobuf:"varint,9,opt,name=PersistentType,json=persistentType,enum=proto.PersistentType" json:"PersistentType,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto1.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *File) GetBucket() *StringType {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *File) GetKey() *StringType {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *File) GetSize() *IntegerType {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *File) GetPersistentID() *StringType {
	if m != nil {
		return m.PersistentID
	}
	return nil
}

type PhotoTimeAxis struct {
	Date []string `protobuf:"bytes,1,rep,name=Date,json=date" json:"Date,omitempty"`
}

func (m *PhotoTimeAxis) Reset()                    { *m = PhotoTimeAxis{} }
func (m *PhotoTimeAxis) String() string            { return proto1.CompactTextString(m) }
func (*PhotoTimeAxis) ProtoMessage()               {}
func (*PhotoTimeAxis) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func init() {
	proto1.RegisterType((*BasePhoto)(nil), "proto.BasePhoto")
	proto1.RegisterType((*Photo)(nil), "proto.Photo")
	proto1.RegisterType((*Note)(nil), "proto.Note")
	proto1.RegisterType((*File)(nil), "proto.File")
	proto1.RegisterType((*PhotoTimeAxis)(nil), "proto.PhotoTimeAxis")
	proto1.RegisterEnum("proto.PhotoType", PhotoType_name, PhotoType_value)
	proto1.RegisterEnum("proto.PhotoPrivacy", PhotoPrivacy_name, PhotoPrivacy_value)
	proto1.RegisterEnum("proto.PersistentType", PersistentType_name, PersistentType_value)
}

func init() { proto1.RegisterFile("photo.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xb1, 0xe3, 0xa4, 0xf1, 0x49, 0x9a, 0x75, 0xa7, 0xdb, 0x65, 0x5a, 0xba, 0x22, 0x64,
	0xf9, 0x08, 0x05, 0x15, 0xe8, 0x0a, 0x81, 0x90, 0xb8, 0x68, 0x93, 0x02, 0x16, 0xd5, 0x12, 0xb9,
	0x0d, 0x5c, 0x70, 0x51, 0xb9, 0xc9, 0x49, 0x3a, 0x5a, 0xc7, 0xb6, 0xec, 0xf1, 0xb6, 0xe6, 0x8e,
	0xe7, 0xe0, 0x35, 0xe0, 0x45, 0x78, 0x22, 0x34, 0x63, 0x3b, 0xb5, 0xb3, 0x9d, 0x6d, 0x2e, 0xd8,
	0xab, 0x64, 0xce, 0xf9, 0x79, 0x66, 0xce, 0xff, 0x7c, 0xd8, 0xd0, 0x0a, 0xaf, 0x03, 0x1e, 0x1c,
	0x86, 0x51, 0xc0, 0x03, 0x52, 0x97, 0x3f, 0x7b, 0x1d, 0x3f, 0xf1, 0x3c, 0x9e, 0x86, 0x98, 0x99,
	0xf7, 0x20, 0x89, 0x31, 0xca, 0xfe, 0xf7, 0xfe, 0xd1, 0xc1, 0x3c, 0x71, 0x63, 0x1c, 0x89, 0xc7,
	0x48, 0x07, 0x74, 0x7b, 0x48, 0xb5, 0xae, 0xd6, 0xaf, 0x39, 0x3a, 0x1b, 0x12, 0x02, 0xc6, 0x78,
	0x6c, 0x0f, 0xa9, 0xde, 0xd5, 0xfa, 0xa6, 0x63, 0x24, 0x63, 0x7b, 0x48, 0xfa, 0x50, 0xff, 0x8d,
	0x4d, 0xf9, 0x35, 0xad, 0x75, 0xb5, 0x7e, 0xeb, 0x88, 0x64, 0x1b, 0x1d, 0xda, 0x3e, 0xc7, 0x39,
	0x46, 0x17, 0x69, 0x88, 0x4e, 0xfd, 0x46, 0x00, 0xe4, 0x00, 0x1a, 0x3f, 0x21, 0x9b, 0x5f, 0x73,
	0x6a, 0x28, 0xd1, 0xc6, 0xb5, 0x24, 0xc8, 0x67, 0xb0, 0xf1, 0x03, 0xf3, 0x70, 0xec, 0x9c, 0xd1,
	0xba, 0x84, 0xb7, 0x72, 0xf8, 0x9c, 0x47, 0xcc, 0x9f, 0x4b, 0x76, 0x63, 0x96, 0x11, 0x64, 0x0f,
	0x9a, 0x12, 0x16, 0x57, 0x6b, 0xc8, 0xab, 0x35, 0x67, 0xf9, 0x9a, 0x1c, 0x66, 0xbe, 0x73, 0xf6,
	0x07, 0xd2, 0x0d, 0xe5, 0xb1, 0x92, 0x17, 0x0c, 0xf9, 0x1a, 0xda, 0xa3, 0x88, 0x2d, 0xdc, 0x28,
	0x1d, 0x04, 0x5e, 0x10, 0xd1, 0xa6, 0xea, 0xf4, 0x76, 0x58, 0xc2, 0x7a, 0x7f, 0x99, 0x50, 0x5f,
	0x5f, 0xb3, 0x7d, 0x30, 0x07, 0x11, 0xba, 0x1c, 0xa7, 0xc7, 0x5c, 0xea, 0x66, 0x3a, 0xe6, 0xa4,
	0x30, 0x90, 0xf7, 0xc1, 0x18, 0xc7, 0x18, 0xe5, 0x2a, 0xb5, 0xf2, 0xa3, 0x85, 0xc9, 0x31, 0x44,
	0xaa, 0x44, 0xbc, 0x62, 0x25, 0xb7, 0xad, 0x67, 0xf1, 0x26, 0xf9, 0x9a, 0x7c, 0x0e, 0x4d, 0xc7,
	0xbd, 0x91, 0x57, 0x91, 0x5a, 0xb4, 0x8e, 0xac, 0x7c, 0x83, 0x65, 0x5a, 0x9d, 0x66, 0x94, 0x13,
	0x77, 0xc9, 0xdb, 0x58, 0x3f, 0x79, 0xcd, 0x07, 0x93, 0x77, 0x04, 0x60, 0xfb, 0x23, 0x16, 0xa2,
	0xc7, 0x7c, 0xa4, 0x66, 0x85, 0x3f, 0x09, 0x02, 0x0f, 0x5d, 0x5f, 0xf2, 0xc0, 0x96, 0x14, 0x39,
	0x04, 0x53, 0x5e, 0x49, 0x38, 0x28, 0x74, 0xb5, 0x7e, 0x67, 0x79, 0xf1, 0xa5, 0xdd, 0x31, 0xc3,
	0xe2, 0x6f, 0x25, 0xe7, 0xad, 0x37, 0xe4, 0xbc, 0xbd, 0x46, 0xce, 0x4b, 0xc5, 0xb6, 0xf9, 0x60,
	0xb1, 0x7d, 0x05, 0x60, 0x4f, 0xd1, 0xe7, 0x6c, 0xc6, 0x30, 0xa2, 0x1d, 0x15, 0x0f, 0x6c, 0x09,
	0x91, 0x6f, 0xa0, 0x2d, 0x63, 0x18, 0x45, 0xec, 0x95, 0x3b, 0x49, 0xe9, 0x23, 0x19, 0xde, 0x76,
	0x39, 0xbc, 0xdc, 0xe5, 0xb4, 0xc3, 0xd2, 0xea, 0xb5, 0x62, 0xb4, 0xd6, 0x2a, 0x46, 0xf2, 0x1c,
	0x5a, 0x3f, 0x9e, 0xfe, 0x72, 0x16, 0x4c, 0x5c, 0xce, 0x02, 0x9f, 0x6e, 0xa9, 0x9e, 0x6a, 0xcd,
	0xef, 0x28, 0x51, 0x75, 0x2f, 0x02, 0x8e, 0x94, 0x54, 0xaa, 0x4e, 0x98, 0x1c, 0xc3, 0x0f, 0xb8,
	0x54, 0x5c, 0xac, 0xa4, 0xe2, 0xdb, 0x99, 0xe2, 0x7e, 0xbe, 0x26, 0xcf, 0xa0, 0x26, 0xd4, 0x7b,
	0xac, 0x3a, 0xa9, 0x96, 0x38, 0x67, 0xe4, 0x3b, 0xe8, 0x0c, 0x59, 0x1c, 0x7a, 0x6e, 0xfa, 0x2b,
	0x46, 0xb1, 0xb8, 0xd9, 0x8e, 0x32, 0x39, 0x9d, 0x69, 0x85, 0x24, 0x5f, 0x80, 0x79, 0x3a, 0x65,
	0x7c, 0xe4, 0x46, 0xee, 0x82, 0x3e, 0x51, 0x1d, 0x63, 0x62, 0xc1, 0x88, 0x1a, 0xb0, 0xe3, 0xe3,
	0x57, 0x2e, 0x77, 0x23, 0xfa, 0xae, 0xb2, 0x02, 0x9b, 0x2c, 0x67, 0x44, 0x7d, 0xdb, 0xf1, 0x45,
	0x12, 0x07, 0x94, 0x2a, 0xe9, 0x06, 0x93, 0x04, 0xf9, 0x16, 0x36, 0x07, 0xc1, 0x62, 0x81, 0x3e,
	0x8f, 0x07, 0x41, 0xe2, 0x73, 0xba, 0xab, 0x8c, 0x63, 0x73, 0x52, 0x06, 0x45, 0xe3, 0x5f, 0xb0,
	0x05, 0xc6, 0xdc, 0x5d, 0x84, 0x74, 0x2f, 0x6b, 0x7c, 0x5e, 0x18, 0xc8, 0x47, 0x60, 0x9c, 0xde,
	0xb2, 0x19, 0x7d, 0x4f, 0x15, 0x9f, 0x81, 0xb7, 0x6c, 0x46, 0xbe, 0x04, 0xd3, 0x8e, 0x6d, 0xcf,
	0xc3, 0xb9, 0xeb, 0xd1, 0x7d, 0xe5, 0x6d, 0x4d, 0x56, 0x40, 0xbd, 0x3f, 0xf5, 0x2c, 0xb9, 0xff,
	0xc3, 0x70, 0x2a, 0xcf, 0x1e, 0x63, 0x65, 0xf6, 0x7c, 0x02, 0xf5, 0x0b, 0xc6, 0x3d, 0x54, 0x8f,
	0xec, 0x3a, 0x17, 0x7e, 0xd1, 0x70, 0x83, 0xc0, 0xe7, 0xe8, 0xf3, 0x7c, 0x46, 0xdd, 0xd7, 0x70,
	0x93, 0x8c, 0x10, 0xbb, 0x9e, 0xf3, 0xd4, 0x2b, 0xc6, 0xf7, 0x7d, 0xbb, 0xc6, 0xc2, 0x5f, 0x15,
	0xb7, 0xb9, 0x22, 0x6e, 0xef, 0x5f, 0x1d, 0x0c, 0xd1, 0xe5, 0x6f, 0x59, 0x83, 0x4f, 0xa1, 0x71,
	0x92, 0x4c, 0x5e, 0x22, 0x57, 0x8b, 0xd0, 0xb8, 0x92, 0x80, 0x68, 0x9a, 0x9f, 0x31, 0x55, 0x2b,
	0x50, 0x7b, 0x89, 0x29, 0xf9, 0x18, 0x8c, 0x07, 0xde, 0x5d, 0x46, 0x5c, 0xbc, 0xb7, 0x44, 0xaf,
	0xc4, 0x42, 0x33, 0x7b, 0xf8, 0xa6, 0xf7, 0x56, 0x09, 0x23, 0xdf, 0x43, 0xe7, 0xee, 0x31, 0x39,
	0x7b, 0x4d, 0x39, 0x9c, 0x76, 0x8a, 0xe1, 0x54, 0x71, 0x3a, 0x9d, 0xb0, 0xb2, 0xee, 0x3d, 0x83,
	0xcd, 0x6c, 0x3a, 0xb3, 0x05, 0x1e, 0xdf, 0xb2, 0x58, 0x88, 0x39, 0x74, 0x39, 0x52, 0xad, 0x5b,
	0x13, 0x62, 0x4e, 0x5d, 0x8e, 0x07, 0x2f, 0x4a, 0xa3, 0x9d, 0x6c, 0xc3, 0x23, 0x39, 0xe2, 0x2e,
	0xc5, 0x07, 0xc8, 0xa5, 0xf8, 0x12, 0xb1, 0xde, 0x21, 0x3b, 0xb0, 0x55, 0x32, 0xc6, 0x69, 0xcc,
	0x71, 0x61, 0x69, 0x2b, 0x66, 0xd1, 0xdb, 0x38, 0xb5, 0xf4, 0x83, 0xdf, 0xab, 0xe3, 0x94, 0x3c,
	0x01, 0x92, 0x61, 0x61, 0x66, 0x28, 0x76, 0xa5, 0xf0, 0xb8, 0x6a, 0x0f, 0x93, 0x2b, 0x8f, 0x4d,
	0x2c, 0x8d, 0xec, 0xc2, 0xce, 0x8a, 0x47, 0xfc, 0x72, 0xb4, 0xf4, 0x83, 0xbf, 0xb5, 0x55, 0x45,
	0xe4, 0x3e, 0x4b, 0x4b, 0xe5, 0xde, 0x4f, 0x61, 0xf7, 0x35, 0x0f, 0xde, 0x5c, 0xca, 0xbd, 0x2d,
	0x8d, 0x7c, 0x08, 0xdd, 0xfb, 0xdc, 0x59, 0x20, 0x39, 0xa5, 0x93, 0x2e, 0xec, 0xaf, 0x52, 0x15,
	0xa2, 0x46, 0x3e, 0x80, 0xa7, 0x0a, 0xc2, 0x95, 0xc3, 0xcb, 0x32, 0xae, 0x1a, 0x32, 0x5d, 0xcf,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xad, 0xa4, 0x80, 0x3e, 0xf0, 0x09, 0x00, 0x00,
}
