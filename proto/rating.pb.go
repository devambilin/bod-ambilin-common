// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/rating.proto

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

type AgentRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId        string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	AgentId        string `protobuf:"bytes,2,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Rating         int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	KetepatanWaktu bool   `protobuf:"varint,4,opt,name=ketepatanWaktu,proto3" json:"ketepatanWaktu,omitempty"`
	Kesopanan      bool   `protobuf:"varint,5,opt,name=kesopanan,proto3" json:"kesopanan,omitempty"`
	Ketelitian     bool   `protobuf:"varint,6,opt,name=ketelitian,proto3" json:"ketelitian,omitempty"`
	Kerapihan      bool   `protobuf:"varint,7,opt,name=kerapihan,proto3" json:"kerapihan,omitempty"`
	Keterangan     string `protobuf:"bytes,8,opt,name=keterangan,proto3" json:"keterangan,omitempty"`
}

func (x *AgentRating) Reset() {
	*x = AgentRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRating) ProtoMessage() {}

func (x *AgentRating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRating.ProtoReflect.Descriptor instead.
func (*AgentRating) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{0}
}

func (x *AgentRating) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *AgentRating) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentRating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AgentRating) GetKetepatanWaktu() bool {
	if x != nil {
		return x.KetepatanWaktu
	}
	return false
}

func (x *AgentRating) GetKesopanan() bool {
	if x != nil {
		return x.Kesopanan
	}
	return false
}

func (x *AgentRating) GetKetelitian() bool {
	if x != nil {
		return x.Ketelitian
	}
	return false
}

func (x *AgentRating) GetKerapihan() bool {
	if x != nil {
		return x.Kerapihan
	}
	return false
}

func (x *AgentRating) GetKeterangan() string {
	if x != nil {
		return x.Keterangan
	}
	return ""
}

type UserRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Rating  int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *UserRating) Reset() {
	*x = UserRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRating) ProtoMessage() {}

func (x *UserRating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRating.ProtoReflect.Descriptor instead.
func (*UserRating) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{1}
}

func (x *UserRating) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UserRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AgentRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string       `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string       `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	Rating           *AgentRating `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *AgentRatingRequest) Reset() {
	*x = AgentRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRatingRequest) ProtoMessage() {}

func (x *AgentRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRatingRequest.ProtoReflect.Descriptor instead.
func (*AgentRatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{2}
}

func (x *AgentRatingRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AgentRatingRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AgentRatingRequest) GetRating() *AgentRating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type UserRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string      `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string      `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	Rating           *UserRating `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *UserRatingRequest) Reset() {
	*x = UserRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRatingRequest) ProtoMessage() {}

func (x *UserRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRatingRequest.ProtoReflect.Descriptor instead.
func (*UserRatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{3}
}

func (x *UserRatingRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UserRatingRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *UserRatingRequest) GetRating() *UserRating {
	if x != nil {
		return x.Rating
	}
	return nil
}

type AgentRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string            `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string            `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	RatingList       []*AgentRating    `protobuf:"bytes,6,rep,name=ratingList,proto3" json:"ratingList,omitempty"`
	RatingTotal      *AgentRatingTotal `protobuf:"bytes,7,opt,name=ratingTotal,proto3" json:"ratingTotal,omitempty"`
}

func (x *AgentRatingResponse) Reset() {
	*x = AgentRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRatingResponse) ProtoMessage() {}

func (x *AgentRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRatingResponse.ProtoReflect.Descriptor instead.
func (*AgentRatingResponse) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{4}
}

func (x *AgentRatingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AgentRatingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AgentRatingResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AgentRatingResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AgentRatingResponse) GetRatingList() []*AgentRating {
	if x != nil {
		return x.RatingList
	}
	return nil
}

func (x *AgentRatingResponse) GetRatingTotal() *AgentRatingTotal {
	if x != nil {
		return x.RatingTotal
	}
	return nil
}

type UserRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string           `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string           `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	RatingList       []*UserRating    `protobuf:"bytes,6,rep,name=ratingList,proto3" json:"ratingList,omitempty"`
	RatingTotal      *UserRatingTotal `protobuf:"bytes,7,opt,name=ratingTotal,proto3" json:"ratingTotal,omitempty"`
}

func (x *UserRatingResponse) Reset() {
	*x = UserRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRatingResponse) ProtoMessage() {}

func (x *UserRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRatingResponse.ProtoReflect.Descriptor instead.
func (*UserRatingResponse) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{5}
}

func (x *UserRatingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserRatingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserRatingResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *UserRatingResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *UserRatingResponse) GetRatingList() []*UserRating {
	if x != nil {
		return x.RatingList
	}
	return nil
}

func (x *UserRatingResponse) GetRatingTotal() *UserRatingTotal {
	if x != nil {
		return x.RatingTotal
	}
	return nil
}

type AgentRatingTotal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId        string  `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	TotalRating    float64 `protobuf:"fixed64,2,opt,name=totalRating,proto3" json:"totalRating,omitempty"`
	TotalReviewer  float64 `protobuf:"fixed64,3,opt,name=totalReviewer,proto3" json:"totalReviewer,omitempty"`
	CoverageRating float64 `protobuf:"fixed64,4,opt,name=coverageRating,proto3" json:"coverageRating,omitempty"`
}

func (x *AgentRatingTotal) Reset() {
	*x = AgentRatingTotal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRatingTotal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRatingTotal) ProtoMessage() {}

func (x *AgentRatingTotal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRatingTotal.ProtoReflect.Descriptor instead.
func (*AgentRatingTotal) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{6}
}

func (x *AgentRatingTotal) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentRatingTotal) GetTotalRating() float64 {
	if x != nil {
		return x.TotalRating
	}
	return 0
}

func (x *AgentRatingTotal) GetTotalReviewer() float64 {
	if x != nil {
		return x.TotalReviewer
	}
	return 0
}

func (x *AgentRatingTotal) GetCoverageRating() float64 {
	if x != nil {
		return x.CoverageRating
	}
	return 0
}

type UserRatingTotal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TotalRating    float64 `protobuf:"fixed64,2,opt,name=totalRating,proto3" json:"totalRating,omitempty"`
	TotalReviewer  float64 `protobuf:"fixed64,3,opt,name=totalReviewer,proto3" json:"totalReviewer,omitempty"`
	CoverageRating float64 `protobuf:"fixed64,4,opt,name=coverageRating,proto3" json:"coverageRating,omitempty"`
}

func (x *UserRatingTotal) Reset() {
	*x = UserRatingTotal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rating_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRatingTotal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRatingTotal) ProtoMessage() {}

func (x *UserRatingTotal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rating_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRatingTotal.ProtoReflect.Descriptor instead.
func (*UserRatingTotal) Descriptor() ([]byte, []int) {
	return file_proto_rating_proto_rawDescGZIP(), []int{7}
}

func (x *UserRatingTotal) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRatingTotal) GetTotalRating() float64 {
	if x != nil {
		return x.TotalRating
	}
	return 0
}

func (x *UserRatingTotal) GetTotalReviewer() float64 {
	if x != nil {
		return x.TotalReviewer
	}
	return 0
}

func (x *UserRatingTotal) GetCoverageRating() float64 {
	if x != nil {
		return x.CoverageRating
	}
	return 0
}

var File_proto_rating_proto protoreflect.FileDescriptor

var file_proto_rating_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01,
	0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x65, 0x74,
	0x65, 0x70, 0x61, 0x74, 0x61, 0x6e, 0x57, 0x61, 0x6b, 0x74, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x6b, 0x65, 0x74, 0x65, 0x70, 0x61, 0x74, 0x61, 0x6e, 0x57, 0x61, 0x6b, 0x74,
	0x75, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x73, 0x6f, 0x70, 0x61, 0x6e, 0x61, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6b, 0x65, 0x73, 0x6f, 0x70, 0x61, 0x6e, 0x61, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x6c, 0x69, 0x74, 0x69, 0x61, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x6c, 0x69, 0x74, 0x69, 0x61, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x72, 0x61, 0x70, 0x69, 0x68, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x6b, 0x65, 0x72, 0x61, 0x70, 0x69, 0x68, 0x61, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x22, 0x56, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x88, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x80, 0x02,
	0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0xfd, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x9c, 0x01, 0x0a, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x99, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x32, 0xfc, 0x06, 0x0a, 0x0d,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x12,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x19, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x61, 0x6d, 0x62, 0x69,
	0x6c, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x64, 0x2d, 0x61, 0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rating_proto_rawDescOnce sync.Once
	file_proto_rating_proto_rawDescData = file_proto_rating_proto_rawDesc
)

func file_proto_rating_proto_rawDescGZIP() []byte {
	file_proto_rating_proto_rawDescOnce.Do(func() {
		file_proto_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rating_proto_rawDescData)
	})
	return file_proto_rating_proto_rawDescData
}

var file_proto_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_rating_proto_goTypes = []interface{}{
	(*AgentRating)(nil),         // 0: proto.AgentRating
	(*UserRating)(nil),          // 1: proto.UserRating
	(*AgentRatingRequest)(nil),  // 2: proto.AgentRatingRequest
	(*UserRatingRequest)(nil),   // 3: proto.UserRatingRequest
	(*AgentRatingResponse)(nil), // 4: proto.AgentRatingResponse
	(*UserRatingResponse)(nil),  // 5: proto.UserRatingResponse
	(*AgentRatingTotal)(nil),    // 6: proto.AgentRatingTotal
	(*UserRatingTotal)(nil),     // 7: proto.UserRatingTotal
	(*BaseRequest)(nil),         // 8: proto.BaseRequest
	(*BaseResponse)(nil),        // 9: proto.BaseResponse
}
var file_proto_rating_proto_depIdxs = []int32{
	0,  // 0: proto.AgentRatingRequest.rating:type_name -> proto.AgentRating
	1,  // 1: proto.UserRatingRequest.rating:type_name -> proto.UserRating
	0,  // 2: proto.AgentRatingResponse.ratingList:type_name -> proto.AgentRating
	6,  // 3: proto.AgentRatingResponse.ratingTotal:type_name -> proto.AgentRatingTotal
	1,  // 4: proto.UserRatingResponse.ratingList:type_name -> proto.UserRating
	7,  // 5: proto.UserRatingResponse.ratingTotal:type_name -> proto.UserRatingTotal
	2,  // 6: proto.RatingService.CreateAgentRating:input_type -> proto.AgentRatingRequest
	3,  // 7: proto.RatingService.CreateCustomerRating:input_type -> proto.UserRatingRequest
	2,  // 8: proto.RatingService.GetAgentRatings:input_type -> proto.AgentRatingRequest
	2,  // 9: proto.RatingService.GetAgentRatingTotal:input_type -> proto.AgentRatingRequest
	3,  // 10: proto.RatingService.GetCustomerRatings:input_type -> proto.UserRatingRequest
	3,  // 11: proto.RatingService.GetCustomerRatingTotal:input_type -> proto.UserRatingRequest
	8,  // 12: proto.RatingService.CustomerCreateAgentRating:input_type -> proto.BaseRequest
	8,  // 13: proto.RatingService.CustomerGetRatings:input_type -> proto.BaseRequest
	8,  // 14: proto.RatingService.CustomerGetRatingTotal:input_type -> proto.BaseRequest
	8,  // 15: proto.RatingService.AgentCreateCustomerRating:input_type -> proto.BaseRequest
	8,  // 16: proto.RatingService.AgentGetRatings:input_type -> proto.BaseRequest
	8,  // 17: proto.RatingService.AgentGetRatingTotal:input_type -> proto.BaseRequest
	4,  // 18: proto.RatingService.CreateAgentRating:output_type -> proto.AgentRatingResponse
	5,  // 19: proto.RatingService.CreateCustomerRating:output_type -> proto.UserRatingResponse
	4,  // 20: proto.RatingService.GetAgentRatings:output_type -> proto.AgentRatingResponse
	4,  // 21: proto.RatingService.GetAgentRatingTotal:output_type -> proto.AgentRatingResponse
	5,  // 22: proto.RatingService.GetCustomerRatings:output_type -> proto.UserRatingResponse
	5,  // 23: proto.RatingService.GetCustomerRatingTotal:output_type -> proto.UserRatingResponse
	9,  // 24: proto.RatingService.CustomerCreateAgentRating:output_type -> proto.BaseResponse
	9,  // 25: proto.RatingService.CustomerGetRatings:output_type -> proto.BaseResponse
	9,  // 26: proto.RatingService.CustomerGetRatingTotal:output_type -> proto.BaseResponse
	9,  // 27: proto.RatingService.AgentCreateCustomerRating:output_type -> proto.BaseResponse
	9,  // 28: proto.RatingService.AgentGetRatings:output_type -> proto.BaseResponse
	9,  // 29: proto.RatingService.AgentGetRatingTotal:output_type -> proto.BaseResponse
	18, // [18:30] is the sub-list for method output_type
	6,  // [6:18] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_rating_proto_init() }
func file_proto_rating_proto_init() {
	if File_proto_rating_proto != nil {
		return
	}
	file_proto_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRating); i {
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
		file_proto_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRating); i {
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
		file_proto_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRatingRequest); i {
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
		file_proto_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRatingRequest); i {
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
		file_proto_rating_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRatingResponse); i {
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
		file_proto_rating_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRatingResponse); i {
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
		file_proto_rating_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRatingTotal); i {
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
		file_proto_rating_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRatingTotal); i {
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
			RawDescriptor: file_proto_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rating_proto_goTypes,
		DependencyIndexes: file_proto_rating_proto_depIdxs,
		MessageInfos:      file_proto_rating_proto_msgTypes,
	}.Build()
	File_proto_rating_proto = out.File
	file_proto_rating_proto_rawDesc = nil
	file_proto_rating_proto_goTypes = nil
	file_proto_rating_proto_depIdxs = nil
}
