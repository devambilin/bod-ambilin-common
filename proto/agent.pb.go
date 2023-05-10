// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/agent.proto

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

type AgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string     `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	Agent            *DataAgent `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *AgentRequest) Reset() {
	*x = AgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRequest) ProtoMessage() {}

func (x *AgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRequest.ProtoReflect.Descriptor instead.
func (*AgentRequest) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AgentRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AgentRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AgentRequest) GetAgent() *DataAgent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type AgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string       `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string       `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	DataAgent        *DataAgent   `protobuf:"bytes,5,opt,name=dataAgent,proto3" json:"dataAgent,omitempty"`
	ListAgent        []*DataAgent `protobuf:"bytes,6,rep,name=listAgent,proto3" json:"listAgent,omitempty"`
}

func (x *AgentResponse) Reset() {
	*x = AgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentResponse) ProtoMessage() {}

func (x *AgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentResponse.ProtoReflect.Descriptor instead.
func (*AgentResponse) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AgentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AgentResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AgentResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AgentResponse) GetDataAgent() *DataAgent {
	if x != nil {
		return x.DataAgent
	}
	return nil
}

func (x *AgentResponse) GetListAgent() []*DataAgent {
	if x != nil {
		return x.ListAgent
	}
	return nil
}

type DataAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created               int64   `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Updated               int64   `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	Removed               int64   `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	CreatedBy             string  `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy             string  `protobuf:"bytes,5,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	Id                    string  `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Ktp                   string  `protobuf:"bytes,7,opt,name=ktp,proto3" json:"ktp,omitempty"`
	FullName              string  `protobuf:"bytes,8,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Religion              string  `protobuf:"bytes,9,opt,name=religion,proto3" json:"religion,omitempty"`
	Gender                string  `protobuf:"bytes,10,opt,name=gender,proto3" json:"gender,omitempty"`
	DateOfBirth           string  `protobuf:"bytes,11,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	PlaceOfBirth          string  `protobuf:"bytes,12,opt,name=placeOfBirth,proto3" json:"placeOfBirth,omitempty"`
	Address               string  `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	ZipCode               int32   `protobuf:"varint,14,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	AddressDomicile       string  `protobuf:"bytes,15,opt,name=addressDomicile,proto3" json:"addressDomicile,omitempty"`
	ZipCodeDomicile       int32   `protobuf:"varint,16,opt,name=zipCodeDomicile,proto3" json:"zipCodeDomicile,omitempty"`
	Phone                 string  `protobuf:"bytes,17,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                 string  `protobuf:"bytes,18,opt,name=email,proto3" json:"email,omitempty"`
	Longitude             float32 `protobuf:"fixed32,19,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude              float32 `protobuf:"fixed32,20,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Image                 string  `protobuf:"bytes,21,opt,name=image,proto3" json:"image,omitempty"`
	Notification          bool    `protobuf:"varint,22,opt,name=notification,proto3" json:"notification,omitempty"`
	IsOnline              bool    `protobuf:"varint,23,opt,name=isOnline,proto3" json:"isOnline,omitempty"`
	CostInstant           string  `protobuf:"bytes,24,opt,name=costInstant,proto3" json:"costInstant,omitempty"`
	Status                int32   `protobuf:"varint,25,opt,name=status,proto3" json:"status,omitempty"`
	CostLater             string  `protobuf:"bytes,26,opt,name=costLater,proto3" json:"costLater,omitempty"`
	LimitPickup           string  `protobuf:"bytes,27,opt,name=limitPickup,proto3" json:"limitPickup,omitempty"`
	Mid                   string  `protobuf:"bytes,28,opt,name=mid,proto3" json:"mid,omitempty"`
	PhotoUrl              string  `protobuf:"bytes,29,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	AccessId              string  `protobuf:"bytes,30,opt,name=accessId,proto3" json:"accessId,omitempty"`
	KancaId               int32   `protobuf:"varint,31,opt,name=kancaId,proto3" json:"kancaId,omitempty"`
	PabId                 int32   `protobuf:"varint,32,opt,name=pabId,proto3" json:"pabId,omitempty"`
	FirstLogin            bool    `protobuf:"varint,33,opt,name=firstLogin,proto3" json:"firstLogin,omitempty"`
	LastUpdateDataBrilink int64   `protobuf:"varint,34,opt,name=lastUpdateDataBrilink,proto3" json:"lastUpdateDataBrilink,omitempty"`
	BranchCode            string  `protobuf:"bytes,35,opt,name=branchCode,proto3" json:"branchCode,omitempty"`
	BranchName            string  `protobuf:"bytes,36,opt,name=branchName,proto3" json:"branchName,omitempty"`
	RegionCode            string  `protobuf:"bytes,37,opt,name=regionCode,proto3" json:"regionCode,omitempty"`
	RegionName            string  `protobuf:"bytes,38,opt,name=regionName,proto3" json:"regionName,omitempty"`
	MerchantCode          string  `protobuf:"bytes,39,opt,name=merchantCode,proto3" json:"merchantCode,omitempty"`
	VehicleCode           string  `protobuf:"bytes,40,opt,name=vehicleCode,proto3" json:"vehicleCode,omitempty"`
	NumberPolice          string  `protobuf:"bytes,41,opt,name=numberPolice,proto3" json:"numberPolice,omitempty"`
	SetupLocation         bool    `protobuf:"varint,42,opt,name=setupLocation,proto3" json:"setupLocation,omitempty"`
}

func (x *DataAgent) Reset() {
	*x = DataAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAgent) ProtoMessage() {}

func (x *DataAgent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAgent.ProtoReflect.Descriptor instead.
func (*DataAgent) Descriptor() ([]byte, []int) {
	return file_proto_agent_proto_rawDescGZIP(), []int{2}
}

func (x *DataAgent) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *DataAgent) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *DataAgent) GetRemoved() int64 {
	if x != nil {
		return x.Removed
	}
	return 0
}

func (x *DataAgent) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DataAgent) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DataAgent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataAgent) GetKtp() string {
	if x != nil {
		return x.Ktp
	}
	return ""
}

func (x *DataAgent) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *DataAgent) GetReligion() string {
	if x != nil {
		return x.Religion
	}
	return ""
}

func (x *DataAgent) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *DataAgent) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *DataAgent) GetPlaceOfBirth() string {
	if x != nil {
		return x.PlaceOfBirth
	}
	return ""
}

func (x *DataAgent) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DataAgent) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

func (x *DataAgent) GetAddressDomicile() string {
	if x != nil {
		return x.AddressDomicile
	}
	return ""
}

func (x *DataAgent) GetZipCodeDomicile() int32 {
	if x != nil {
		return x.ZipCodeDomicile
	}
	return 0
}

func (x *DataAgent) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DataAgent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DataAgent) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DataAgent) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DataAgent) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DataAgent) GetNotification() bool {
	if x != nil {
		return x.Notification
	}
	return false
}

func (x *DataAgent) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *DataAgent) GetCostInstant() string {
	if x != nil {
		return x.CostInstant
	}
	return ""
}

func (x *DataAgent) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DataAgent) GetCostLater() string {
	if x != nil {
		return x.CostLater
	}
	return ""
}

func (x *DataAgent) GetLimitPickup() string {
	if x != nil {
		return x.LimitPickup
	}
	return ""
}

func (x *DataAgent) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *DataAgent) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *DataAgent) GetAccessId() string {
	if x != nil {
		return x.AccessId
	}
	return ""
}

func (x *DataAgent) GetKancaId() int32 {
	if x != nil {
		return x.KancaId
	}
	return 0
}

func (x *DataAgent) GetPabId() int32 {
	if x != nil {
		return x.PabId
	}
	return 0
}

func (x *DataAgent) GetFirstLogin() bool {
	if x != nil {
		return x.FirstLogin
	}
	return false
}

func (x *DataAgent) GetLastUpdateDataBrilink() int64 {
	if x != nil {
		return x.LastUpdateDataBrilink
	}
	return 0
}

func (x *DataAgent) GetBranchCode() string {
	if x != nil {
		return x.BranchCode
	}
	return ""
}

func (x *DataAgent) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *DataAgent) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *DataAgent) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *DataAgent) GetMerchantCode() string {
	if x != nil {
		return x.MerchantCode
	}
	return ""
}

func (x *DataAgent) GetVehicleCode() string {
	if x != nil {
		return x.VehicleCode
	}
	return ""
}

func (x *DataAgent) GetNumberPolice() string {
	if x != nil {
		return x.NumberPolice
	}
	return ""
}

func (x *DataAgent) GetSetupLocation() bool {
	if x != nil {
		return x.SetupLocation
	}
	return false
}

var File_proto_agent_proto protoreflect.FileDescriptor

var file_proto_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0xeb, 0x01,
	0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x6c,
	0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0xeb, 0x09, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x74, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x74, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x6d, 0x69, 0x63, 0x69,
	0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x44, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x7a, 0x69, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x6d, 0x69, 0x63,
	0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x61, 0x6e, 0x63, 0x61, 0x49, 0x64, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6b, 0x61, 0x6e, 0x63, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x62, 0x49, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61,
	0x62, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x72, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x72, 0x69, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x18,
	0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb5, 0x02, 0x0a, 0x0c, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x65, 0x76, 0x61, 0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x64, 0x2d, 0x61,
	0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_agent_proto_rawDescOnce sync.Once
	file_proto_agent_proto_rawDescData = file_proto_agent_proto_rawDesc
)

func file_proto_agent_proto_rawDescGZIP() []byte {
	file_proto_agent_proto_rawDescOnce.Do(func() {
		file_proto_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_agent_proto_rawDescData)
	})
	return file_proto_agent_proto_rawDescData
}

var file_proto_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_agent_proto_goTypes = []interface{}{
	(*AgentRequest)(nil),  // 0: proto.AgentRequest
	(*AgentResponse)(nil), // 1: proto.AgentResponse
	(*DataAgent)(nil),     // 2: proto.DataAgent
}
var file_proto_agent_proto_depIdxs = []int32{
	2, // 0: proto.AgentRequest.agent:type_name -> proto.DataAgent
	2, // 1: proto.AgentResponse.dataAgent:type_name -> proto.DataAgent
	2, // 2: proto.AgentResponse.listAgent:type_name -> proto.DataAgent
	0, // 3: proto.AgentService.GetAgents:input_type -> proto.AgentRequest
	0, // 4: proto.AgentService.GetAgent:input_type -> proto.AgentRequest
	0, // 5: proto.AgentService.CreateAgent:input_type -> proto.AgentRequest
	0, // 6: proto.AgentService.UpdateAgent:input_type -> proto.AgentRequest
	0, // 7: proto.AgentService.DeleteAgent:input_type -> proto.AgentRequest
	1, // 8: proto.AgentService.GetAgents:output_type -> proto.AgentResponse
	1, // 9: proto.AgentService.GetAgent:output_type -> proto.AgentResponse
	1, // 10: proto.AgentService.CreateAgent:output_type -> proto.AgentResponse
	1, // 11: proto.AgentService.UpdateAgent:output_type -> proto.AgentResponse
	1, // 12: proto.AgentService.DeleteAgent:output_type -> proto.AgentResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_agent_proto_init() }
func file_proto_agent_proto_init() {
	if File_proto_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRequest); i {
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
		file_proto_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentResponse); i {
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
		file_proto_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataAgent); i {
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
			RawDescriptor: file_proto_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_agent_proto_goTypes,
		DependencyIndexes: file_proto_agent_proto_depIdxs,
		MessageInfos:      file_proto_agent_proto_msgTypes,
	}.Build()
	File_proto_agent_proto = out.File
	file_proto_agent_proto_rawDesc = nil
	file_proto_agent_proto_goTypes = nil
	file_proto_agent_proto_depIdxs = nil
}
