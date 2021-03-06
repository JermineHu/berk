// Code generated by protoc-gen-go.
// source: dod.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_NULL    OrderStatus = 0
	OrderStatus_ORDER_STATUS_NOTPAID OrderStatus = 1
	OrderStatus_ORDER_STATUS_PAYING  OrderStatus = 2
	OrderStatus_ORDER_STATUS_PAYED   OrderStatus = 3
	OrderStatus_ORDER_STATUS_CLOSED  OrderStatus = 4
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_STATUS_NULL",
	1: "ORDER_STATUS_NOTPAID",
	2: "ORDER_STATUS_PAYING",
	3: "ORDER_STATUS_PAYED",
	4: "ORDER_STATUS_CLOSED",
}
var OrderStatus_value = map[string]int32{
	"ORDER_STATUS_NULL":    0,
	"ORDER_STATUS_NOTPAID": 1,
	"ORDER_STATUS_PAYING":  2,
	"ORDER_STATUS_PAYED":   3,
	"ORDER_STATUS_CLOSED":  4,
}

func (x OrderStatus) String() string {
	return proto1.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type PayFrom int32

const (
	PayFrom_ORDER_PAYFRM_NULL   PayFrom = 0
	PayFrom_ORDER_PAYFRM_WECHAT PayFrom = 1
	PayFrom_ORDER_PAYFRM_ALIPAY PayFrom = 2
)

var PayFrom_name = map[int32]string{
	0: "ORDER_PAYFRM_NULL",
	1: "ORDER_PAYFRM_WECHAT",
	2: "ORDER_PAYFRM_ALIPAY",
}
var PayFrom_value = map[string]int32{
	"ORDER_PAYFRM_NULL":   0,
	"ORDER_PAYFRM_WECHAT": 1,
	"ORDER_PAYFRM_ALIPAY": 2,
}

func (x PayFrom) String() string {
	return proto1.EnumName(PayFrom_name, int32(x))
}
func (PayFrom) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type ItemType int32

const (
	ItemType_ITEM_TYPE_NULL   ItemType = 0
	ItemType_ITEM_TYPE_FILTER ItemType = 1
)

var ItemType_name = map[int32]string{
	0: "ITEM_TYPE_NULL",
	1: "ITEM_TYPE_FILTER",
}
var ItemType_value = map[string]int32{
	"ITEM_TYPE_NULL":   0,
	"ITEM_TYPE_FILTER": 1,
}

func (x ItemType) String() string {
	return proto1.EnumName(ItemType_name, int32(x))
}
func (ItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type OrderList struct {
	Data []*Order `protobuf:"bytes,1,rep,name=Data,json=data" json:"Data,omitempty"`
}

func (m *OrderList) Reset()                    { *m = OrderList{} }
func (m *OrderList) String() string            { return proto1.CompactTextString(m) }
func (*OrderList) ProtoMessage()               {}
func (*OrderList) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *OrderList) GetData() []*Order {
	if m != nil {
		return m.Data
	}
	return nil
}

type BatchSetOrderStatus struct {
	OrderUUIDs  []string    `protobuf:"bytes,1,rep,name=OrderUUIDs,json=orderUUIDs" json:"OrderUUIDs,omitempty"`
	OrderStatus OrderStatus `protobuf:"varint,2,opt,name=OrderStatus,json=orderStatus,enum=proto.OrderStatus" json:"OrderStatus,omitempty"`
}

func (m *BatchSetOrderStatus) Reset()                    { *m = BatchSetOrderStatus{} }
func (m *BatchSetOrderStatus) String() string            { return proto1.CompactTextString(m) }
func (*BatchSetOrderStatus) ProtoMessage()               {}
func (*BatchSetOrderStatus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type Order struct {
	ID            int64        `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID          string       `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	CreatedAt     string       `protobuf:"bytes,3,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	User          *User        `protobuf:"bytes,4,opt,name=User,json=user" json:"User,omitempty"`
	UserUUID      string       `protobuf:"bytes,5,opt,name=UserUUID,json=userUUID" json:"UserUUID,omitempty"`
	Number        *StringType  `protobuf:"bytes,6,opt,name=Number,json=number" json:"Number,omitempty"`
	Items         []*OrderItem `protobuf:"bytes,7,rep,name=Items,json=items" json:"Items,omitempty"`
	PrimeCost     float32      `protobuf:"fixed32,8,opt,name=PrimeCost,json=primeCost" json:"PrimeCost,omitempty"`
	Currency      *StringType  `protobuf:"bytes,9,opt,name=Currency,json=currency" json:"Currency,omitempty"`
	Status        OrderStatus  `protobuf:"varint,10,opt,name=Status,json=status,enum=proto.OrderStatus" json:"Status,omitempty"`
	PayFrom       PayFrom      `protobuf:"varint,11,opt,name=PayFrom,json=payFrom,enum=proto.PayFrom" json:"PayFrom,omitempty"`
	TransactionID *StringType  `protobuf:"bytes,12,opt,name=TransactionID,json=transactionID" json:"TransactionID,omitempty"`
	Timestamp     string       `protobuf:"bytes,13,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	UpdatedAt     string       `protobuf:"bytes,14,opt,name=UpdatedAt,json=updatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt     string       `protobuf:"bytes,15,opt,name=DeletedAt,json=deletedAt" json:"DeletedAt,omitempty"`
	TotalPrice    float32      `protobuf:"fixed32,16,opt,name=TotalPrice,json=totalPrice" json:"TotalPrice,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto1.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Order) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Order) GetNumber() *StringType {
	if m != nil {
		return m.Number
	}
	return nil
}

func (m *Order) GetItems() []*OrderItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetCurrency() *StringType {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *Order) GetTransactionID() *StringType {
	if m != nil {
		return m.TransactionID
	}
	return nil
}

type OrderItem struct {
	ID           int64                `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID         string               `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	UpdatedAt    string               `protobuf:"bytes,3,opt,name=UpdatedAt,json=updatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt    string               `protobuf:"bytes,4,opt,name=DeletedAt,json=deletedAt" json:"DeletedAt,omitempty"`
	CreatedAt    string               `protobuf:"bytes,5,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	Snapshots    []*OrderItemSnapshot `protobuf:"bytes,6,rep,name=Snapshots,json=snapshots" json:"Snapshots,omitempty"`
	PrimeCost    float32              `protobuf:"fixed32,7,opt,name=PrimeCost,json=primeCost" json:"PrimeCost,omitempty"`
	RealPrice    float32              `protobuf:"fixed32,8,opt,name=RealPrice,json=realPrice" json:"RealPrice,omitempty"`
	Currency     *StringType          `protobuf:"bytes,9,opt,name=Currency,json=currency" json:"Currency,omitempty"`
	OrderUUID    string               `protobuf:"bytes,10,opt,name=OrderUUID,json=orderUUID" json:"OrderUUID,omitempty"`
	ItemUUID     *StringType          `protobuf:"bytes,11,opt,name=ItemUUID,json=itemUUID" json:"ItemUUID,omitempty"`
	ItemType     ItemType             `protobuf:"varint,12,opt,name=ItemType,json=itemType,enum=proto.ItemType" json:"ItemType,omitempty"`
	ItemName     *StringType          `protobuf:"bytes,13,opt,name=ItemName,json=itemName" json:"ItemName,omitempty"`
	ItemImageUrl *StringType          `protobuf:"bytes,14,opt,name=ItemImageUrl,json=itemImageUrl" json:"ItemImageUrl,omitempty"`
	Count        *IntegerType         `protobuf:"bytes,15,opt,name=Count,json=count" json:"Count,omitempty"`
}

func (m *OrderItem) Reset()                    { *m = OrderItem{} }
func (m *OrderItem) String() string            { return proto1.CompactTextString(m) }
func (*OrderItem) ProtoMessage()               {}
func (*OrderItem) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *OrderItem) GetSnapshots() []*OrderItemSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func (m *OrderItem) GetCurrency() *StringType {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *OrderItem) GetItemUUID() *StringType {
	if m != nil {
		return m.ItemUUID
	}
	return nil
}

func (m *OrderItem) GetItemName() *StringType {
	if m != nil {
		return m.ItemName
	}
	return nil
}

func (m *OrderItem) GetItemImageUrl() *StringType {
	if m != nil {
		return m.ItemImageUrl
	}
	return nil
}

func (m *OrderItem) GetCount() *IntegerType {
	if m != nil {
		return m.Count
	}
	return nil
}

type OrderItemSnapshot struct {
	ID            int64       `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	UUID          string      `protobuf:"bytes,2,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
	UpdatedAt     string      `protobuf:"bytes,3,opt,name=UpdatedAt,json=updatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt     string      `protobuf:"bytes,4,opt,name=DeletedAt,json=deletedAt" json:"DeletedAt,omitempty"`
	CreatedAt     string      `protobuf:"bytes,5,opt,name=CreatedAt,json=createdAt" json:"CreatedAt,omitempty"`
	OrderItemUUID string      `protobuf:"bytes,6,opt,name=OrderItemUUID,json=orderItemUUID" json:"OrderItemUUID,omitempty"`
	Key           *StringType `protobuf:"bytes,7,opt,name=Key,json=key" json:"Key,omitempty"`
	Value         *StringType `protobuf:"bytes,8,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *OrderItemSnapshot) Reset()                    { *m = OrderItemSnapshot{} }
func (m *OrderItemSnapshot) String() string            { return proto1.CompactTextString(m) }
func (*OrderItemSnapshot) ProtoMessage()               {}
func (*OrderItemSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *OrderItemSnapshot) GetKey() *StringType {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *OrderItemSnapshot) GetValue() *StringType {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterType((*OrderList)(nil), "proto.OrderList")
	proto1.RegisterType((*BatchSetOrderStatus)(nil), "proto.BatchSetOrderStatus")
	proto1.RegisterType((*Order)(nil), "proto.Order")
	proto1.RegisterType((*OrderItem)(nil), "proto.OrderItem")
	proto1.RegisterType((*OrderItemSnapshot)(nil), "proto.OrderItemSnapshot")
	proto1.RegisterEnum("proto.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto1.RegisterEnum("proto.PayFrom", PayFrom_name, PayFrom_value)
	proto1.RegisterEnum("proto.ItemType", ItemType_name, ItemType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DodServices service

type DodServicesClient interface {
	NewOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Order, error)
	NewOrders(ctx context.Context, in *OrderList, opts ...client.CallOption) (*OrderList, error)
	DeleteOrdersByUUIDs(ctx context.Context, in *StringIDs, opts ...client.CallOption) (*Bool, error)
	GetOrderByUUIDs(ctx context.Context, in *StringIDs, opts ...client.CallOption) (*OrderList, error)
	UpdateOrdersStatus(ctx context.Context, in *BatchSetOrderStatus, opts ...client.CallOption) (*Bool, error)
	FindOrdersByParams(ctx context.Context, in *RequestParams, opts ...client.CallOption) (*PageModel, error)
}

type dodServicesClient struct {
	c           client.Client
	serviceName string
}

func NewDodServicesClient(serviceName string, c client.Client) DodServicesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &dodServicesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dodServicesClient) NewOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Order, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.NewOrder", in)
	out := new(Order)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dodServicesClient) NewOrders(ctx context.Context, in *OrderList, opts ...client.CallOption) (*OrderList, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.NewOrders", in)
	out := new(OrderList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dodServicesClient) DeleteOrdersByUUIDs(ctx context.Context, in *StringIDs, opts ...client.CallOption) (*Bool, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.DeleteOrdersByUUIDs", in)
	out := new(Bool)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dodServicesClient) GetOrderByUUIDs(ctx context.Context, in *StringIDs, opts ...client.CallOption) (*OrderList, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.GetOrderByUUIDs", in)
	out := new(OrderList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dodServicesClient) UpdateOrdersStatus(ctx context.Context, in *BatchSetOrderStatus, opts ...client.CallOption) (*Bool, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.UpdateOrdersStatus", in)
	out := new(Bool)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dodServicesClient) FindOrdersByParams(ctx context.Context, in *RequestParams, opts ...client.CallOption) (*PageModel, error) {
	req := c.c.NewRequest(c.serviceName, "DodServices.FindOrdersByParams", in)
	out := new(PageModel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DodServices service

type DodServicesHandler interface {
	NewOrder(context.Context, *Order, *Order) error
	NewOrders(context.Context, *OrderList, *OrderList) error
	DeleteOrdersByUUIDs(context.Context, *StringIDs, *Bool) error
	GetOrderByUUIDs(context.Context, *StringIDs, *OrderList) error
	UpdateOrdersStatus(context.Context, *BatchSetOrderStatus, *Bool) error
	FindOrdersByParams(context.Context, *RequestParams, *PageModel) error
}

func RegisterDodServicesHandler(s server.Server, hdlr DodServicesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DodServices{hdlr}, opts...))
}

type DodServices struct {
	DodServicesHandler
}

func (h *DodServices) NewOrder(ctx context.Context, in *Order, out *Order) error {
	return h.DodServicesHandler.NewOrder(ctx, in, out)
}

func (h *DodServices) NewOrders(ctx context.Context, in *OrderList, out *OrderList) error {
	return h.DodServicesHandler.NewOrders(ctx, in, out)
}

func (h *DodServices) DeleteOrdersByUUIDs(ctx context.Context, in *StringIDs, out *Bool) error {
	return h.DodServicesHandler.DeleteOrdersByUUIDs(ctx, in, out)
}

func (h *DodServices) GetOrderByUUIDs(ctx context.Context, in *StringIDs, out *OrderList) error {
	return h.DodServicesHandler.GetOrderByUUIDs(ctx, in, out)
}

func (h *DodServices) UpdateOrdersStatus(ctx context.Context, in *BatchSetOrderStatus, out *Bool) error {
	return h.DodServicesHandler.UpdateOrdersStatus(ctx, in, out)
}

func (h *DodServices) FindOrdersByParams(ctx context.Context, in *RequestParams, out *PageModel) error {
	return h.DodServicesHandler.FindOrdersByParams(ctx, in, out)
}

func init() { proto1.RegisterFile("dod.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 945 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xce, 0x8f, 0x9d, 0xc6, 0xc7, 0x49, 0xea, 0x4e, 0x0b, 0x58, 0x51, 0x05, 0x51, 0x40, 0x10,
	0x8a, 0x76, 0x25, 0xc2, 0xb2, 0x7b, 0x07, 0x4a, 0xe3, 0x74, 0xb1, 0x48, 0x53, 0xcb, 0x71, 0x40,
	0xbd, 0xaa, 0x66, 0xe3, 0x51, 0xd7, 0xda, 0xd8, 0xce, 0x8e, 0xc7, 0x45, 0x79, 0x00, 0x24, 0xde,
	0x86, 0x3b, 0x9e, 0x86, 0x87, 0x41, 0xf3, 0xe3, 0xfc, 0xb5, 0x41, 0x70, 0xb5, 0x57, 0xf5, 0xf9,
	0xbe, 0xef, 0x9c, 0x39, 0x3e, 0xf3, 0xc5, 0xa7, 0x60, 0x84, 0x69, 0xf8, 0x7c, 0x49, 0x53, 0x96,
	0x22, 0x5d, 0xfc, 0x69, 0xb7, 0x92, 0x7c, 0xb1, 0x60, 0xab, 0x25, 0x91, 0x70, 0x1b, 0xf2, 0x8c,
	0x50, 0xf5, 0x6c, 0x86, 0x94, 0xe0, 0x58, 0x05, 0x8d, 0x68, 0x8e, 0x69, 0x9e, 0xa9, 0xe8, 0xec,
	0x7d, 0x4e, 0xe8, 0x6a, 0x89, 0x29, 0x8e, 0x09, 0x2b, 0x12, 0xba, 0xcf, 0xc0, 0xb8, 0xa1, 0x21,
	0xa1, 0xe3, 0x28, 0x63, 0xa8, 0x03, 0x9a, 0x83, 0x19, 0xb6, 0xcb, 0x9d, 0x6a, 0xcf, 0xec, 0x37,
	0xa4, 0xe4, 0xb9, 0xe0, 0x7d, 0x2d, 0xc4, 0x0c, 0x77, 0xdf, 0xc1, 0xe9, 0x25, 0x66, 0xf3, 0xb7,
	0x53, 0xc2, 0x04, 0x3c, 0x65, 0x98, 0xe5, 0x19, 0xfa, 0x14, 0x40, 0x84, 0xb3, 0x99, 0xeb, 0x64,
	0x22, 0xdd, 0xf0, 0x21, 0x5d, 0x23, 0xe8, 0x05, 0x98, 0x5b, 0x72, 0xbb, 0xd2, 0x29, 0xf7, 0x5a,
	0x7d, 0xb4, 0x5d, 0x5f, 0x32, 0xbe, 0x99, 0x6e, 0x82, 0xee, 0x9f, 0x1a, 0xe8, 0x82, 0x44, 0x2d,
	0xa8, 0xb8, 0x8e, 0x5d, 0xee, 0x94, 0x7b, 0x55, 0xbf, 0x12, 0x39, 0x08, 0x81, 0xc6, 0x0b, 0x8b,
	0x42, 0x86, 0xaf, 0xe5, 0x33, 0xd7, 0x41, 0xe7, 0x60, 0x0c, 0x29, 0xc1, 0x8c, 0x84, 0x03, 0x66,
	0x57, 0x05, 0x61, 0xcc, 0x0b, 0x00, 0x7d, 0x06, 0xda, 0x2c, 0x23, 0xd4, 0xd6, 0x3a, 0xe5, 0x9e,
	0xd9, 0x37, 0xd5, 0xd1, 0x1c, 0xf2, 0x35, 0x3e, 0x3f, 0xd4, 0x86, 0x3a, 0x8f, 0x44, 0x59, 0x5d,
	0x64, 0xd7, 0x73, 0x15, 0xa3, 0xaf, 0xa1, 0x36, 0xc9, 0xe3, 0x37, 0x84, 0xda, 0x35, 0x91, 0x7e,
	0xa2, 0xd2, 0xa7, 0x8c, 0x46, 0xc9, 0x7d, 0xb0, 0x5a, 0x12, 0xbf, 0x96, 0x08, 0x01, 0xfa, 0x12,
	0x74, 0x97, 0x91, 0x38, 0xb3, 0x8f, 0xc4, 0x0c, 0xad, 0xed, 0x77, 0xe4, 0x84, 0xaf, 0x47, 0x9c,
	0xe6, 0xdd, 0x7a, 0x34, 0x8a, 0xc9, 0x30, 0xcd, 0x98, 0x5d, 0xef, 0x94, 0x7b, 0x15, 0xdf, 0x58,
	0x16, 0x00, 0x7a, 0x06, 0xf5, 0x61, 0x4e, 0x29, 0x49, 0xe6, 0x2b, 0xdb, 0x38, 0x74, 0x64, 0x7d,
	0xae, 0x24, 0xe8, 0x02, 0x6a, 0x6a, 0xb2, 0x70, 0x70, 0xb2, 0xb5, 0x4c, 0x5e, 0x55, 0x0f, 0x8e,
	0x3c, 0xbc, 0xba, 0xa2, 0x69, 0x6c, 0x9b, 0x42, 0xdc, 0x52, 0x62, 0x85, 0xfa, 0x47, 0x4b, 0xf9,
	0x80, 0x5e, 0x41, 0x33, 0xa0, 0x38, 0xc9, 0xf0, 0x9c, 0x45, 0x69, 0xe2, 0x3a, 0x76, 0xe3, 0x50,
	0x27, 0x4d, 0xb6, 0xad, 0xe3, 0xef, 0x16, 0x44, 0x31, 0xc9, 0x18, 0x8e, 0x97, 0x76, 0x53, 0xde,
	0x04, 0x2b, 0x00, 0xce, 0xce, 0x96, 0xa1, 0xba, 0xa7, 0x96, 0x64, 0xf3, 0x02, 0xe0, 0xac, 0x43,
	0x16, 0x44, 0xb2, 0xc7, 0x92, 0x0d, 0x0b, 0x80, 0xfb, 0x2c, 0x48, 0x19, 0x5e, 0x78, 0x34, 0x9a,
	0x13, 0xdb, 0x12, 0x63, 0x03, 0xb6, 0x46, 0xba, 0x7f, 0x69, 0xca, 0xce, 0x7c, 0xd4, 0xff, 0xd5,
	0x35, 0x9b, 0x6e, 0xaa, 0xff, 0xda, 0x8d, 0xb6, 0xdf, 0xcd, 0x8e, 0xe3, 0xf4, 0x7d, 0xc7, 0xbd,
	0x04, 0x63, 0x9a, 0xe0, 0x65, 0xf6, 0x36, 0x65, 0x99, 0x5d, 0x13, 0x6e, 0xb0, 0xf7, 0xdd, 0x50,
	0x08, 0x7c, 0x23, 0x2b, 0xa4, 0xbb, 0xce, 0x38, 0xda, 0x77, 0xc6, 0x39, 0x18, 0x3e, 0x29, 0x06,
	0xa0, 0x7c, 0x43, 0x0b, 0xe0, 0xff, 0xfa, 0xe6, 0x5c, 0x4d, 0x4b, 0x4c, 0x05, 0xe4, 0x0b, 0xac,
	0x7f, 0xb5, 0xbc, 0x18, 0xef, 0x51, 0x90, 0xe6, 0xc1, 0x62, 0x91, 0x92, 0xa0, 0x6f, 0xa4, 0x9c,
	0xa3, 0xc2, 0x29, 0xad, 0xfe, 0xb1, 0x92, 0x17, 0xb0, 0x14, 0xf3, 0xa7, 0xa2, 0xf6, 0x04, 0xc7,
	0x44, 0x38, 0xe4, 0x70, 0x6d, 0x2e, 0x41, 0xdf, 0x43, 0x83, 0xcb, 0xdd, 0x18, 0xdf, 0x93, 0x19,
	0x5d, 0x08, 0xdb, 0x3c, 0x99, 0xd2, 0x88, 0xb6, 0x64, 0xa8, 0x07, 0xfa, 0x30, 0xcd, 0x13, 0x69,
	0x24, 0x73, 0xfd, 0xb3, 0x70, 0x13, 0x46, 0xee, 0x09, 0x15, 0x09, 0xfa, 0x9c, 0x0b, 0xba, 0xbf,
	0x57, 0xe0, 0xe4, 0xd1, 0xad, 0x7c, 0x60, 0x03, 0x7d, 0x01, 0xcd, 0x75, 0x4b, 0xe2, 0xd8, 0x9a,
	0x50, 0x34, 0xd3, 0x6d, 0x10, 0x7d, 0x0e, 0xd5, 0x9f, 0xc9, 0x4a, 0x18, 0xe5, 0xc9, 0x89, 0x54,
	0xdf, 0x91, 0x15, 0xfa, 0x0a, 0xf4, 0x5f, 0xf0, 0x22, 0x97, 0x8e, 0x79, 0x52, 0xa6, 0x3f, 0x70,
	0xfe, 0xe2, 0x8f, 0xf2, 0xce, 0x97, 0x1a, 0x7d, 0x04, 0x27, 0x37, 0xbe, 0x33, 0xf2, 0xef, 0xa6,
	0xc1, 0x20, 0x98, 0x4d, 0xef, 0x26, 0xb3, 0xf1, 0xd8, 0x2a, 0x21, 0x1b, 0xce, 0x76, 0xe1, 0x9b,
	0xc0, 0x1b, 0xb8, 0x8e, 0x55, 0x46, 0x9f, 0xc0, 0xe9, 0x0e, 0xe3, 0x0d, 0x6e, 0xdd, 0xc9, 0x6b,
	0xab, 0x82, 0x3e, 0x06, 0xb4, 0x4f, 0x8c, 0x1c, 0xab, 0xfa, 0x28, 0x61, 0x38, 0xbe, 0x99, 0x8e,
	0x1c, 0x4b, 0xbb, 0xf0, 0xd7, 0x1f, 0xaa, 0x4d, 0x17, 0xde, 0xe0, 0xf6, 0xca, 0xbf, 0x2e, 0xba,
	0x58, 0xa7, 0x2a, 0xf8, 0xd7, 0xd1, 0xf0, 0xa7, 0x41, 0xb0, 0xdd, 0x84, 0x22, 0x06, 0x63, 0xd7,
	0x1b, 0xdc, 0x5a, 0x95, 0x8b, 0x17, 0x1b, 0x8f, 0x22, 0x04, 0x2d, 0x37, 0x18, 0x5d, 0xdf, 0x05,
	0xb7, 0xde, 0xa8, 0xa8, 0x78, 0x06, 0xd6, 0x06, 0xbb, 0x72, 0xc7, 0xc1, 0xc8, 0xb7, 0xca, 0xfd,
	0xbf, 0x2b, 0x60, 0x3a, 0x69, 0x38, 0x25, 0xf4, 0x21, 0x9a, 0x13, 0xfe, 0x09, 0xad, 0x4f, 0xc8,
	0x6f, 0x72, 0x33, 0xed, 0x2c, 0xc9, 0xf6, 0x4e, 0xd4, 0x2d, 0xa1, 0x6f, 0xc1, 0x28, 0x94, 0x19,
	0xda, 0xd9, 0x05, 0x7c, 0xdf, 0xb6, 0x1f, 0x21, 0xdd, 0x12, 0x7a, 0x09, 0xa7, 0xd2, 0x31, 0x32,
	0xeb, 0x72, 0x25, 0x37, 0xa8, 0xb5, 0x73, 0x65, 0xae, 0x93, 0xb5, 0x8b, 0x1d, 0x76, 0x99, 0xa6,
	0x8b, 0x6e, 0x09, 0xbd, 0x82, 0xe3, 0xd7, 0x6a, 0x29, 0x1f, 0xce, 0x79, 0xea, 0xc0, 0x1f, 0x01,
	0x49, 0x03, 0xcb, 0x03, 0xd5, 0xc5, 0xb7, 0x8b, 0xea, 0x8f, 0xb7, 0xfd, 0xfe, 0xc9, 0x3f, 0x00,
	0xba, 0x8a, 0x92, 0xb0, 0xe8, 0xd7, 0xe3, 0xff, 0x61, 0x64, 0xe8, 0x4c, 0x89, 0x7c, 0xf2, 0x3e,
	0x27, 0x19, 0x93, 0xe8, 0xba, 0x01, 0x0f, 0xdf, 0x93, 0xeb, 0x34, 0x24, 0x8b, 0x6e, 0xe9, 0x4d,
	0x4d, 0x40, 0xdf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x63, 0xfe, 0x25, 0xea, 0x08, 0x00,
	0x00,
}
