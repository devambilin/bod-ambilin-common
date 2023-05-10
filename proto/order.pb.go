// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/order.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string    `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string    `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	User             *DataUser `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *OrderRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *OrderRequest) GetUser() *DataUser {
	if x != nil {
		return x.User
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string       `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string       `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	DataOrder        *DataOrder   `protobuf:"bytes,5,opt,name=dataOrder,proto3" json:"dataOrder,omitempty"`
	ListOrder        []*DataOrder `protobuf:"bytes,6,rep,name=listOrder,proto3" json:"listOrder,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *OrderResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *OrderResponse) GetDataOrder() *DataOrder {
	if x != nil {
		return x.DataOrder
	}
	return nil
}

func (x *OrderResponse) GetListOrder() []*DataOrder {
	if x != nil {
		return x.ListOrder
	}
	return nil
}

type DataOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created             int64   `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Updated             int64   `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	Removed             int64   `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	CreatedBy           string  `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy           string  `protobuf:"bytes,5,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	Id                  string  `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	UserId              string  `protobuf:"bytes,7,opt,name=userId,proto3" json:"userId,omitempty"`
	AgentId             string  `protobuf:"bytes,8,opt,name=agentId,proto3" json:"agentId,omitempty"`
	DepositAmount       float32 `protobuf:"fixed32,9,opt,name=depositAmount,proto3" json:"depositAmount,omitempty"`
	ShippingFee         float32 `protobuf:"fixed32,10,opt,name=shippingFee,proto3" json:"shippingFee,omitempty"`
	ShippingMethod      int32   `protobuf:"varint,11,opt,name=shippingMethod,proto3" json:"shippingMethod,omitempty"`
	OrderStatusCustomer int32   `protobuf:"varint,12,opt,name=orderStatusCustomer,proto3" json:"orderStatusCustomer,omitempty"`
	OrderType           string  `protobuf:"bytes,13,opt,name=orderType,proto3" json:"orderType,omitempty"`
	ProgressPickup      int32   `protobuf:"varint,14,opt,name=progressPickup,proto3" json:"progressPickup,omitempty"`
	OrderLater          string  `protobuf:"bytes,15,opt,name=orderLater,proto3" json:"orderLater,omitempty"`
	EstimatedArrival    string  `protobuf:"bytes,16,opt,name=estimatedArrival,proto3" json:"estimatedArrival,omitempty"`
	EstimatedPickup     string  `protobuf:"bytes,17,opt,name=estimatedPickup,proto3" json:"estimatedPickup,omitempty"`
	CustomerPlace       string  `protobuf:"bytes,18,opt,name=customerPlace,proto3" json:"customerPlace,omitempty"`
	CustomerAddress     string  `protobuf:"bytes,19,opt,name=customerAddress,proto3" json:"customerAddress,omitempty"`
	CustomerLatitude    float32 `protobuf:"fixed32,20,opt,name=customerLatitude,proto3" json:"customerLatitude,omitempty"`
	CustomerLongitude   float32 `protobuf:"fixed32,21,opt,name=customerLongitude,proto3" json:"customerLongitude,omitempty"`
	AgentLatitude       float32 `protobuf:"fixed32,22,opt,name=agentLatitude,proto3" json:"agentLatitude,omitempty"`
	AgentLongitude      float32 `protobuf:"fixed32,23,opt,name=agentLongitude,proto3" json:"agentLongitude,omitempty"`
	OrderStatusAgent    int32   `protobuf:"varint,24,opt,name=orderStatusAgent,proto3" json:"orderStatusAgent,omitempty"`
	CostOrder           float32 `protobuf:"fixed32,25,opt,name=costOrder,proto3" json:"costOrder,omitempty"`
	AmountStatus        bool    `protobuf:"varint,26,opt,name=amountStatus,proto3" json:"amountStatus,omitempty"`
	ExpiredTime         int64   `protobuf:"varint,27,opt,name=expiredTime,proto3" json:"expiredTime,omitempty"`
	CompletedTime       int64   `protobuf:"varint,28,opt,name=completedTime,proto3" json:"completedTime,omitempty"`
	Distance            float32 `protobuf:"fixed32,29,opt,name=distance,proto3" json:"distance,omitempty"`
	ConfirmButton       bool    `protobuf:"varint,30,opt,name=confirmButton,proto3" json:"confirmButton,omitempty"`
	ConfirmButtonAgent  bool    `protobuf:"varint,31,opt,name=confirmButtonAgent,proto3" json:"confirmButtonAgent,omitempty"`
	AmountStatusInvalid bool    `protobuf:"varint,32,opt,name=amountStatusInvalid,proto3" json:"amountStatusInvalid,omitempty"`
	OrderCode           string  `protobuf:"bytes,33,opt,name=orderCode,proto3" json:"orderCode,omitempty"`
	IsCountdown         bool    `protobuf:"varint,34,opt,name=isCountdown,proto3" json:"isCountdown,omitempty"`
	LocationDescription string  `protobuf:"bytes,35,opt,name=locationDescription,proto3" json:"locationDescription,omitempty"`
	OrderFinish         int64   `protobuf:"varint,36,opt,name=orderFinish,proto3" json:"orderFinish,omitempty"`
	ExpiredTimeExtend   string  `protobuf:"bytes,37,opt,name=expiredTimeExtend,proto3" json:"expiredTimeExtend,omitempty"`
	UrlInvoice          string  `protobuf:"bytes,38,opt,name=urlInvoice,proto3" json:"urlInvoice,omitempty"`
}

func (x *DataOrder) Reset() {
	*x = DataOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOrder) ProtoMessage() {}

func (x *DataOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOrder.ProtoReflect.Descriptor instead.
func (*DataOrder) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *DataOrder) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *DataOrder) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *DataOrder) GetRemoved() int64 {
	if x != nil {
		return x.Removed
	}
	return 0
}

func (x *DataOrder) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DataOrder) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DataOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataOrder) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DataOrder) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *DataOrder) GetDepositAmount() float32 {
	if x != nil {
		return x.DepositAmount
	}
	return 0
}

func (x *DataOrder) GetShippingFee() float32 {
	if x != nil {
		return x.ShippingFee
	}
	return 0
}

func (x *DataOrder) GetShippingMethod() int32 {
	if x != nil {
		return x.ShippingMethod
	}
	return 0
}

func (x *DataOrder) GetOrderStatusCustomer() int32 {
	if x != nil {
		return x.OrderStatusCustomer
	}
	return 0
}

func (x *DataOrder) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *DataOrder) GetProgressPickup() int32 {
	if x != nil {
		return x.ProgressPickup
	}
	return 0
}

func (x *DataOrder) GetOrderLater() string {
	if x != nil {
		return x.OrderLater
	}
	return ""
}

func (x *DataOrder) GetEstimatedArrival() string {
	if x != nil {
		return x.EstimatedArrival
	}
	return ""
}

func (x *DataOrder) GetEstimatedPickup() string {
	if x != nil {
		return x.EstimatedPickup
	}
	return ""
}

func (x *DataOrder) GetCustomerPlace() string {
	if x != nil {
		return x.CustomerPlace
	}
	return ""
}

func (x *DataOrder) GetCustomerAddress() string {
	if x != nil {
		return x.CustomerAddress
	}
	return ""
}

func (x *DataOrder) GetCustomerLatitude() float32 {
	if x != nil {
		return x.CustomerLatitude
	}
	return 0
}

func (x *DataOrder) GetCustomerLongitude() float32 {
	if x != nil {
		return x.CustomerLongitude
	}
	return 0
}

func (x *DataOrder) GetAgentLatitude() float32 {
	if x != nil {
		return x.AgentLatitude
	}
	return 0
}

func (x *DataOrder) GetAgentLongitude() float32 {
	if x != nil {
		return x.AgentLongitude
	}
	return 0
}

func (x *DataOrder) GetOrderStatusAgent() int32 {
	if x != nil {
		return x.OrderStatusAgent
	}
	return 0
}

func (x *DataOrder) GetCostOrder() float32 {
	if x != nil {
		return x.CostOrder
	}
	return 0
}

func (x *DataOrder) GetAmountStatus() bool {
	if x != nil {
		return x.AmountStatus
	}
	return false
}

func (x *DataOrder) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *DataOrder) GetCompletedTime() int64 {
	if x != nil {
		return x.CompletedTime
	}
	return 0
}

func (x *DataOrder) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *DataOrder) GetConfirmButton() bool {
	if x != nil {
		return x.ConfirmButton
	}
	return false
}

func (x *DataOrder) GetConfirmButtonAgent() bool {
	if x != nil {
		return x.ConfirmButtonAgent
	}
	return false
}

func (x *DataOrder) GetAmountStatusInvalid() bool {
	if x != nil {
		return x.AmountStatusInvalid
	}
	return false
}

func (x *DataOrder) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *DataOrder) GetIsCountdown() bool {
	if x != nil {
		return x.IsCountdown
	}
	return false
}

func (x *DataOrder) GetLocationDescription() string {
	if x != nil {
		return x.LocationDescription
	}
	return ""
}

func (x *DataOrder) GetOrderFinish() int64 {
	if x != nil {
		return x.OrderFinish
	}
	return 0
}

func (x *DataOrder) GetExpiredTimeExtend() string {
	if x != nil {
		return x.ExpiredTimeExtend
	}
	return ""
}

func (x *DataOrder) GetUrlInvoice() string {
	if x != nil {
		return x.UrlInvoice
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xeb, 0x01, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xe9, 0x0a, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x50, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x18, 0x24, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x25, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x72, 0x6c, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x72, 0x6c, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x32, 0xb5, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x61,
	0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x64, 0x2d, 0x61, 0x6d, 0x62, 0x69, 0x6c,
	0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_order_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),  // 0: proto.OrderRequest
	(*OrderResponse)(nil), // 1: proto.OrderResponse
	(*DataOrder)(nil),     // 2: proto.DataOrder
	(*DataUser)(nil),      // 3: proto.DataUser
}
var file_proto_order_proto_depIdxs = []int32{
	3, // 0: proto.OrderRequest.user:type_name -> proto.DataUser
	2, // 1: proto.OrderResponse.dataOrder:type_name -> proto.DataOrder
	2, // 2: proto.OrderResponse.listOrder:type_name -> proto.DataOrder
	0, // 3: proto.OrderService.GetOrders:input_type -> proto.OrderRequest
	0, // 4: proto.OrderService.GetOrder:input_type -> proto.OrderRequest
	0, // 5: proto.OrderService.CreateOrder:input_type -> proto.OrderRequest
	0, // 6: proto.OrderService.UpdateOrder:input_type -> proto.OrderRequest
	0, // 7: proto.OrderService.DeleteOrder:input_type -> proto.OrderRequest
	1, // 8: proto.OrderService.GetOrders:output_type -> proto.OrderResponse
	1, // 9: proto.OrderService.GetOrder:output_type -> proto.OrderResponse
	1, // 10: proto.OrderService.CreateOrder:output_type -> proto.OrderResponse
	1, // 11: proto.OrderService.UpdateOrder:output_type -> proto.OrderResponse
	1, // 12: proto.OrderService.DeleteOrder:output_type -> proto.OrderResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	file_proto_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
