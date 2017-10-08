// Code generated by protoc-gen-go.
// source: github.com/micro/go-os/router/proto/router.proto
// DO NOT EDIT!

/*
Package go_micro_os_router is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-os/router/proto/router.proto

It has these top-level messages:
	Stats
	Service
	Node
	Endpoint
	Value
	Selected
	EndpointStats
	Metrics
*/
package go_micro_os_router

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

type Stats struct {
	// service for these stats match
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// client making the requests
	Client *Service `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
	// time of the sample
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// duration of the sample
	Duration int64 `protobuf:"varint,4,opt,name=duration" json:"duration,omitempty"`
	// Selected stats
	Selected *Selected `protobuf:"bytes,5,opt,name=selected" json:"selected,omitempty"`
	// Endpoint stats
	EndpointStats []*EndpointStats `protobuf:"bytes,6,rep,name=endpoint_stats,json=endpointStats" json:"endpoint_stats,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Stats) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Stats) GetClient() *Service {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Stats) GetSelected() *Selected {
	if m != nil {
		return m.Selected
	}
	return nil
}

func (m *Stats) GetEndpointStats() []*EndpointStats {
	if m != nil {
		return m.EndpointStats
	}
	return nil
}

type Service struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,3,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Endpoints []*Endpoint       `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
	Nodes     []*Node           `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port     int64             `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Endpoint struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Request  *Value            `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Response *Value            `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Endpoint) GetRequest() *Value {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Endpoint) GetResponse() *Value {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Value struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type   string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Values []*Value `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Value) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type Selected struct {
	Errors  *Metrics `protobuf:"bytes,7,opt,name=errors" json:"errors,omitempty"`
	Success *Metrics `protobuf:"bytes,5,opt,name=success" json:"success,omitempty"`
	Dropped *Metrics `protobuf:"bytes,6,opt,name=dropped" json:"dropped,omitempty"`
}

func (m *Selected) Reset()                    { *m = Selected{} }
func (m *Selected) String() string            { return proto.CompactTextString(m) }
func (*Selected) ProtoMessage()               {}
func (*Selected) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Selected) GetErrors() *Metrics {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Selected) GetSuccess() *Metrics {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *Selected) GetDropped() *Metrics {
	if m != nil {
		return m.Dropped
	}
	return nil
}

type EndpointStats struct {
	// Name of the endpoint
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Success and error rates
	Errors  *Metrics `protobuf:"bytes,2,opt,name=errors" json:"errors,omitempty"`
	Success *Metrics `protobuf:"bytes,3,opt,name=success" json:"success,omitempty"`
	Dropped *Metrics `protobuf:"bytes,4,opt,name=dropped" json:"dropped,omitempty"`
}

func (m *EndpointStats) Reset()                    { *m = EndpointStats{} }
func (m *EndpointStats) String() string            { return proto.CompactTextString(m) }
func (*EndpointStats) ProtoMessage()               {}
func (*EndpointStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EndpointStats) GetErrors() *Metrics {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *EndpointStats) GetSuccess() *Metrics {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *EndpointStats) GetDropped() *Metrics {
	if m != nil {
		return m.Dropped
	}
	return nil
}

type Metrics struct {
	Count   int64   `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Mean    float64 `protobuf:"fixed64,2,opt,name=mean" json:"mean,omitempty"`
	StdDev  float64 `protobuf:"fixed64,3,opt,name=std_dev,json=stdDev" json:"std_dev,omitempty"`
	Upper95 float64 `protobuf:"fixed64,4,opt,name=upper95" json:"upper95,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Stats)(nil), "go.micro.os.router.Stats")
	proto.RegisterType((*Service)(nil), "go.micro.os.router.Service")
	proto.RegisterType((*Node)(nil), "go.micro.os.router.Node")
	proto.RegisterType((*Endpoint)(nil), "go.micro.os.router.Endpoint")
	proto.RegisterType((*Value)(nil), "go.micro.os.router.Value")
	proto.RegisterType((*Selected)(nil), "go.micro.os.router.Selected")
	proto.RegisterType((*EndpointStats)(nil), "go.micro.os.router.EndpointStats")
	proto.RegisterType((*Metrics)(nil), "go.micro.os.router.Metrics")
}

func init() { proto.RegisterFile("github.com/micro/go-os/router/proto/router.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x8e, 0xd3, 0x3e,
	0x10, 0x56, 0x92, 0xe6, 0x4f, 0xe7, 0xa7, 0xae, 0x7e, 0xb2, 0x90, 0x08, 0xcb, 0x1e, 0x96, 0x1e,
	0x10, 0x20, 0x91, 0x02, 0xd5, 0x4a, 0xcb, 0x72, 0x43, 0x14, 0x71, 0x81, 0x83, 0x57, 0xe2, 0xba,
	0x4a, 0x63, 0x53, 0x22, 0x9a, 0x38, 0xd8, 0x4e, 0xa5, 0x3e, 0x05, 0x8f, 0xc1, 0x89, 0x03, 0x0f,
	0x01, 0xcf, 0x85, 0xed, 0xd8, 0xdd, 0x5d, 0x6d, 0xda, 0x85, 0x45, 0xdc, 0x66, 0xdc, 0xef, 0xb3,
	0xe7, 0xfb, 0x32, 0x33, 0x85, 0x27, 0x8b, 0x52, 0x7e, 0x6c, 0xe7, 0x59, 0xc1, 0xaa, 0x49, 0x55,
	0x16, 0x9c, 0x4d, 0x16, 0xec, 0x31, 0x13, 0x13, 0xce, 0x5a, 0x49, 0xf9, 0xa4, 0xe1, 0x4c, 0x32,
	0x9b, 0x64, 0x26, 0x41, 0x68, 0xc1, 0x32, 0x83, 0xcc, 0x98, 0xc8, 0xba, 0x5f, 0xc6, 0xdf, 0x7d,
	0x08, 0x4f, 0x65, 0x2e, 0x05, 0x3a, 0x82, 0x58, 0x50, 0xbe, 0x2a, 0x0b, 0x9a, 0x7a, 0x87, 0xde,
	0x83, 0xff, 0x9e, 0xdd, 0xcd, 0xae, 0xe2, 0xb3, 0xd3, 0x0e, 0x82, 0x1d, 0x16, 0x4d, 0x21, 0x2a,
	0x96, 0x25, 0xad, 0x65, 0xea, 0x5f, 0xcf, 0xb2, 0x50, 0x74, 0x00, 0x43, 0x59, 0x56, 0x54, 0xc8,
	0xbc, 0x6a, 0xd2, 0x40, 0xf1, 0x02, 0x7c, 0x7e, 0x80, 0xf6, 0x21, 0x21, 0x2d, 0xcf, 0x65, 0xc9,
	0xea, 0x74, 0x60, 0x7e, 0xdc, 0xe4, 0xe8, 0x18, 0x12, 0x41, 0x97, 0xb4, 0x90, 0x94, 0xa4, 0xa1,
	0x79, 0xf0, 0xa0, 0xff, 0xc1, 0x0e, 0x83, 0x37, 0x68, 0xf4, 0x06, 0xf6, 0x68, 0x4d, 0x1a, 0x56,
	0xd6, 0xf2, 0x4c, 0x68, 0xc5, 0x69, 0x74, 0x18, 0x28, 0xfe, 0xbd, 0x3e, 0xfe, 0xcc, 0x22, 0x8d,
	0x35, 0x78, 0x44, 0x2f, 0xa6, 0xe3, 0xaf, 0x3e, 0xc4, 0x56, 0x11, 0x42, 0x30, 0xa8, 0xf3, 0xaa,
	0xb3, 0x6c, 0x88, 0x4d, 0x8c, 0x52, 0x88, 0x57, 0x94, 0x0b, 0x5d, 0xbe, 0x6f, 0x8e, 0x5d, 0x8a,
	0x66, 0x90, 0x54, 0x54, 0xe6, 0x24, 0x97, 0xb9, 0x92, 0xad, 0x5f, 0x7f, 0xb8, 0xc3, 0xae, 0xec,
	0xad, 0xc5, 0xce, 0x6a, 0xc9, 0xd7, 0x78, 0x43, 0x45, 0x27, 0x30, 0x74, 0x15, 0x09, 0xe5, 0x50,
	0xb0, 0xcd, 0x05, 0xa7, 0x02, 0x9f, 0xc3, 0x51, 0x06, 0x61, 0xcd, 0x08, 0x15, 0xca, 0x3d, 0xcd,
	0x4b, 0xfb, 0x78, 0xef, 0x14, 0x00, 0x77, 0xb0, 0xfd, 0x17, 0x30, 0xba, 0x54, 0x06, 0xfa, 0x1f,
	0x82, 0x4f, 0x74, 0x6d, 0x05, 0xeb, 0x10, 0xdd, 0x82, 0x70, 0x95, 0x2f, 0x5b, 0x6a, 0xd5, 0x76,
	0xc9, 0x89, 0x7f, 0xec, 0x8d, 0x7f, 0x78, 0x30, 0xd0, 0x97, 0xa1, 0x3d, 0xf0, 0x4b, 0x62, 0x39,
	0x2a, 0xd2, 0x16, 0xe5, 0x84, 0x70, 0x2a, 0x84, 0xb3, 0xc8, 0xa6, 0xda, 0xd0, 0x86, 0x71, 0x69,
	0xbb, 0xc2, 0xc4, 0xe8, 0xe5, 0x05, 0xdb, 0x3a, 0xb9, 0xf7, 0xb7, 0x95, 0xbd, 0xcd, 0xb3, 0xbf,
	0xd3, 0xf1, 0xc5, 0x87, 0xc4, 0x99, 0xd9, 0xfb, 0xc9, 0xa7, 0x10, 0x73, 0xfa, 0xb9, 0x55, 0x0d,
	0x6c, 0xc7, 0xe0, 0x4e, 0x5f, 0x81, 0xef, 0xf5, 0x85, 0xd8, 0x21, 0xd5, 0xc4, 0x25, 0x4a, 0x72,
	0xc3, 0x6a, 0x41, 0x8d, 0xdc, 0x9d, 0xac, 0x0d, 0x14, 0xbd, 0xbe, 0xe2, 0xc6, 0xa3, 0x5d, 0x1f,
	0xff, 0xdf, 0x38, 0x32, 0x87, 0xd0, 0xd4, 0xd5, 0xeb, 0x86, 0x3a, 0x93, 0xeb, 0xc6, 0xb1, 0x4c,
	0x8c, 0x9e, 0x42, 0x64, 0xd8, 0xc2, 0x36, 0xfe, 0x0e, 0xa9, 0x16, 0x38, 0xfe, 0xe6, 0x41, 0xe2,
	0x06, 0x59, 0xef, 0x19, 0xca, 0x39, 0xe3, 0x22, 0x8d, 0xb7, 0xef, 0x19, 0xa5, 0x87, 0x97, 0x85,
	0xc0, 0x16, 0x6a, 0x76, 0x5a, 0x5b, 0x14, 0xba, 0xcd, 0xc2, 0xeb, 0x59, 0x0e, 0xab, 0x69, 0x84,
	0xb3, 0xa6, 0x51, 0x3b, 0x26, 0xfa, 0x0d, 0x9a, 0xc5, 0x8e, 0x7f, 0x7a, 0x30, 0xba, 0xb4, 0x38,
	0xb6, 0xb4, 0x8a, 0x13, 0xe2, 0xdf, 0x48, 0x48, 0x70, 0x33, 0x21, 0x83, 0x3f, 0x10, 0xf2, 0x01,
	0x62, 0x7b, 0xa6, 0x3b, 0xa0, 0x60, 0xad, 0xda, 0xee, 0x9e, 0x99, 0xc7, 0x2e, 0xd1, 0xba, 0x2a,
	0x9a, 0x77, 0xeb, 0xcd, 0xc3, 0x26, 0x46, 0xb7, 0x55, 0x89, 0x92, 0x9c, 0x11, 0xba, 0x32, 0x25,
	0x7a, 0x38, 0x52, 0xe9, 0x2b, 0xba, 0xd2, 0xb3, 0xde, 0xaa, 0x6b, 0xf9, 0xf3, 0x23, 0x53, 0x84,
	0x87, 0x5d, 0x3a, 0x8f, 0xcc, 0xff, 0xd2, 0xf4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x94,
	0x83, 0xfa, 0xcb, 0x06, 0x00, 0x00,
}