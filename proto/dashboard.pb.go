// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/dashboard.proto

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

type MasterConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string            `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string            `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	MasterConfig     *DataMasterConfig `protobuf:"bytes,3,opt,name=masterConfig,proto3" json:"masterConfig,omitempty"`
}

func (x *MasterConfigRequest) Reset() {
	*x = MasterConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterConfigRequest) ProtoMessage() {}

func (x *MasterConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterConfigRequest.ProtoReflect.Descriptor instead.
func (*MasterConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *MasterConfigRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MasterConfigRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *MasterConfigRequest) GetMasterConfig() *DataMasterConfig {
	if x != nil {
		return x.MasterConfig
	}
	return nil
}

type MasterConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32               `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string              `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string              `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	DataMasterConfig *DataMasterConfig   `protobuf:"bytes,5,opt,name=dataMasterConfig,proto3" json:"dataMasterConfig,omitempty"`
	ListMasterConfig []*DataMasterConfig `protobuf:"bytes,6,rep,name=listMasterConfig,proto3" json:"listMasterConfig,omitempty"`
}

func (x *MasterConfigResponse) Reset() {
	*x = MasterConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterConfigResponse) ProtoMessage() {}

func (x *MasterConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterConfigResponse.ProtoReflect.Descriptor instead.
func (*MasterConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *MasterConfigResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MasterConfigResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MasterConfigResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MasterConfigResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *MasterConfigResponse) GetDataMasterConfig() *DataMasterConfig {
	if x != nil {
		return x.DataMasterConfig
	}
	return nil
}

func (x *MasterConfigResponse) GetListMasterConfig() []*DataMasterConfig {
	if x != nil {
		return x.ListMasterConfig
	}
	return nil
}

type DataMasterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParamKey   string `protobuf:"bytes,2,opt,name=paramKey,proto3" json:"paramKey,omitempty"`
	ParamValue string `protobuf:"bytes,3,opt,name=paramValue,proto3" json:"paramValue,omitempty"`
}

func (x *DataMasterConfig) Reset() {
	*x = DataMasterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMasterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMasterConfig) ProtoMessage() {}

func (x *DataMasterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMasterConfig.ProtoReflect.Descriptor instead.
func (*DataMasterConfig) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *DataMasterConfig) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataMasterConfig) GetParamKey() string {
	if x != nil {
		return x.ParamKey
	}
	return ""
}

func (x *DataMasterConfig) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

type AdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string     `protobuf:"bytes,2,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	Admin            *DataAdmin `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *AdminRequest) Reset() {
	*x = AdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRequest) ProtoMessage() {}

func (x *AdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRequest.ProtoReflect.Descriptor instead.
func (*AdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *AdminRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AdminRequest) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AdminRequest) GetAdmin() *DataAdmin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type AdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           int32        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message          string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId        string       `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	RequestByService string       `protobuf:"bytes,4,opt,name=requestByService,proto3" json:"requestByService,omitempty"`
	DataAdmin        *DataAdmin   `protobuf:"bytes,5,opt,name=dataAdmin,proto3" json:"dataAdmin,omitempty"`
	ListAdmin        []*DataAdmin `protobuf:"bytes,6,rep,name=listAdmin,proto3" json:"listAdmin,omitempty"`
}

func (x *AdminResponse) Reset() {
	*x = AdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResponse) ProtoMessage() {}

func (x *AdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResponse.ProtoReflect.Descriptor instead.
func (*AdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *AdminResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdminResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AdminResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AdminResponse) GetRequestByService() string {
	if x != nil {
		return x.RequestByService
	}
	return ""
}

func (x *AdminResponse) GetDataAdmin() *DataAdmin {
	if x != nil {
		return x.DataAdmin
	}
	return nil
}

func (x *AdminResponse) GetListAdmin() []*DataAdmin {
	if x != nil {
		return x.ListAdmin
	}
	return nil
}

type DataAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created          int64  `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Updated          int64  `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	Removed          int64  `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	CreatedBy        string `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy        string `protobuf:"bytes,5,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	Id               string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	UserId           string `protobuf:"bytes,7,opt,name=userId,proto3" json:"userId,omitempty"`
	Name             string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Address          string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	Email            string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Password         string `protobuf:"bytes,11,opt,name=password,proto3" json:"password,omitempty"`
	Phone            string `protobuf:"bytes,12,opt,name=phone,proto3" json:"phone,omitempty"`
	PhotoUrl         string `protobuf:"bytes,13,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	AccessId         string `protobuf:"bytes,14,opt,name=accessId,proto3" json:"accessId,omitempty"`
	ResetCode        string `protobuf:"bytes,15,opt,name=resetCode,proto3" json:"resetCode,omitempty"`
	ResetExpiredTime int64  `protobuf:"varint,16,opt,name=resetExpiredTime,proto3" json:"resetExpiredTime,omitempty"`
	IdOfficeRegional int32  `protobuf:"varint,17,opt,name=idOfficeRegional,proto3" json:"idOfficeRegional,omitempty"`
	PersonalNumber   string `protobuf:"bytes,18,opt,name=personalNumber,proto3" json:"personalNumber,omitempty"`
	IsSuperAdmin     bool   `protobuf:"varint,19,opt,name=isSuperAdmin,proto3" json:"isSuperAdmin,omitempty"`
	Role             string `protobuf:"bytes,20,opt,name=role,proto3" json:"role,omitempty"`
	RoleId           int32  `protobuf:"varint,21,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *DataAdmin) Reset() {
	*x = DataAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dashboard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAdmin) ProtoMessage() {}

func (x *DataAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dashboard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAdmin.ProtoReflect.Descriptor instead.
func (*DataAdmin) Descriptor() ([]byte, []int) {
	return file_proto_dashboard_proto_rawDescGZIP(), []int{5}
}

func (x *DataAdmin) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *DataAdmin) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *DataAdmin) GetRemoved() int64 {
	if x != nil {
		return x.Removed
	}
	return 0
}

func (x *DataAdmin) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DataAdmin) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DataAdmin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataAdmin) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DataAdmin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataAdmin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DataAdmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DataAdmin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DataAdmin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DataAdmin) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *DataAdmin) GetAccessId() string {
	if x != nil {
		return x.AccessId
	}
	return ""
}

func (x *DataAdmin) GetResetCode() string {
	if x != nil {
		return x.ResetCode
	}
	return ""
}

func (x *DataAdmin) GetResetExpiredTime() int64 {
	if x != nil {
		return x.ResetExpiredTime
	}
	return 0
}

func (x *DataAdmin) GetIdOfficeRegional() int32 {
	if x != nil {
		return x.IdOfficeRegional
	}
	return 0
}

func (x *DataAdmin) GetPersonalNumber() string {
	if x != nil {
		return x.PersonalNumber
	}
	return ""
}

func (x *DataAdmin) GetIsSuperAdmin() bool {
	if x != nil {
		return x.IsSuperAdmin
	}
	return false
}

func (x *DataAdmin) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *DataAdmin) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

var File_proto_dashboard_proto protoreflect.FileDescriptor

var file_proto_dashboard_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9c, 0x02,
	0x0a, 0x14, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x6c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5e, 0x0a, 0x10,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0xeb, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2e, 0x0a,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xd9, 0x04,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x73, 0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x69, 0x64, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x32, 0x87, 0x03, 0x0a, 0x10, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x65, 0x76, 0x61, 0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x64,
	0x2d, 0x61, 0x6d, 0x62, 0x69, 0x6c, 0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dashboard_proto_rawDescOnce sync.Once
	file_proto_dashboard_proto_rawDescData = file_proto_dashboard_proto_rawDesc
)

func file_proto_dashboard_proto_rawDescGZIP() []byte {
	file_proto_dashboard_proto_rawDescOnce.Do(func() {
		file_proto_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dashboard_proto_rawDescData)
	})
	return file_proto_dashboard_proto_rawDescData
}

var file_proto_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_dashboard_proto_goTypes = []interface{}{
	(*MasterConfigRequest)(nil),  // 0: proto.MasterConfigRequest
	(*MasterConfigResponse)(nil), // 1: proto.MasterConfigResponse
	(*DataMasterConfig)(nil),     // 2: proto.DataMasterConfig
	(*AdminRequest)(nil),         // 3: proto.AdminRequest
	(*AdminResponse)(nil),        // 4: proto.AdminResponse
	(*DataAdmin)(nil),            // 5: proto.DataAdmin
}
var file_proto_dashboard_proto_depIdxs = []int32{
	2,  // 0: proto.MasterConfigRequest.masterConfig:type_name -> proto.DataMasterConfig
	2,  // 1: proto.MasterConfigResponse.dataMasterConfig:type_name -> proto.DataMasterConfig
	2,  // 2: proto.MasterConfigResponse.listMasterConfig:type_name -> proto.DataMasterConfig
	5,  // 3: proto.AdminRequest.admin:type_name -> proto.DataAdmin
	5,  // 4: proto.AdminResponse.dataAdmin:type_name -> proto.DataAdmin
	5,  // 5: proto.AdminResponse.listAdmin:type_name -> proto.DataAdmin
	0,  // 6: proto.DashboardService.GetMasterConfig:input_type -> proto.MasterConfigRequest
	3,  // 7: proto.DashboardService.GetAdmins:input_type -> proto.AdminRequest
	3,  // 8: proto.DashboardService.GetAdmin:input_type -> proto.AdminRequest
	3,  // 9: proto.DashboardService.CreateAdmin:input_type -> proto.AdminRequest
	3,  // 10: proto.DashboardService.UpdateAdmin:input_type -> proto.AdminRequest
	3,  // 11: proto.DashboardService.DeleteAdmin:input_type -> proto.AdminRequest
	1,  // 12: proto.DashboardService.GetMasterConfig:output_type -> proto.MasterConfigResponse
	4,  // 13: proto.DashboardService.GetAdmins:output_type -> proto.AdminResponse
	4,  // 14: proto.DashboardService.GetAdmin:output_type -> proto.AdminResponse
	4,  // 15: proto.DashboardService.CreateAdmin:output_type -> proto.AdminResponse
	4,  // 16: proto.DashboardService.UpdateAdmin:output_type -> proto.AdminResponse
	4,  // 17: proto.DashboardService.DeleteAdmin:output_type -> proto.AdminResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_dashboard_proto_init() }
func file_proto_dashboard_proto_init() {
	if File_proto_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterConfigRequest); i {
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
		file_proto_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterConfigResponse); i {
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
		file_proto_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMasterConfig); i {
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
		file_proto_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRequest); i {
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
		file_proto_dashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminResponse); i {
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
		file_proto_dashboard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataAdmin); i {
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
			RawDescriptor: file_proto_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dashboard_proto_goTypes,
		DependencyIndexes: file_proto_dashboard_proto_depIdxs,
		MessageInfos:      file_proto_dashboard_proto_msgTypes,
	}.Build()
	File_proto_dashboard_proto = out.File
	file_proto_dashboard_proto_rawDesc = nil
	file_proto_dashboard_proto_goTypes = nil
	file_proto_dashboard_proto_depIdxs = nil
}
