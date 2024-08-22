// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: admin.proto

package black_list

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

type CreateHR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ApprovedBy string `protobuf:"bytes,2,opt,name=approved_by,json=approvedBy,proto3" json:"approved_by,omitempty"`
}

func (x *CreateHR) Reset() {
	*x = CreateHR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHR) ProtoMessage() {}

func (x *CreateHR) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHR.ProtoReflect.Descriptor instead.
func (*CreateHR) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHR) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateHR) GetApprovedBy() string {
	if x != nil {
		return x.ApprovedBy
	}
	return ""
}

type GetAllHRRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hr    []*Hr `protobuf:"bytes,1,rep,name=hr,proto3" json:"hr,omitempty"`
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllHRRes) Reset() {
	*x = GetAllHRRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHRRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHRRes) ProtoMessage() {}

func (x *GetAllHRRes) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHRRes.ProtoReflect.Descriptor instead.
func (*GetAllHRRes) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllHRRes) GetHr() []*Hr {
	if x != nil {
		return x.Hr
	}
	return nil
}

func (x *GetAllHRRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Hr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	FullName    string `protobuf:"bytes,3,opt,name=FullName,proto3" json:"FullName,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Hr) Reset() {
	*x = Hr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hr) ProtoMessage() {}

func (x *Hr) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hr.ProtoReflect.Descriptor instead.
func (*Hr) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *Hr) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hr) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Hr) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Hr) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Hr) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ListUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Filter   *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListUserReq) Reset() {
	*x = ListUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReq) ProtoMessage() {}

func (x *ListUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReq.ProtoReflect.Descriptor instead.
func (*ListUserReq) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListUserReq) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *ListUserReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserRes `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Count int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListUserRes) Reset() {
	*x = ListUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRes) ProtoMessage() {}

func (x *ListUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRes.ProtoReflect.Descriptor instead.
func (*ListUserRes) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserRes) GetUsers() []*UserRes {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUserRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	FullName    string `protobuf:"bytes,4,opt,name=FullName,proto3" json:"FullName,omitempty"`
	DateOfBirth string `protobuf:"bytes,5,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Role        string `protobuf:"bytes,6,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *UserRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRes) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserRes) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *UserRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x52, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x22, 0x43, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x52, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x02,
	0x68, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x48, 0x72, 0x52, 0x02, 0x68, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x02, 0x48, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x72, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x4e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x29, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x9d, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x32, 0xf0, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x52, 0x1a, 0x10, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x52,
	0x12, 0x12, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x52, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10,
	0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x61,
	0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_proto_goTypes = []any{
	(*CreateHR)(nil),    // 0: black_list.CreateHR
	(*GetAllHRRes)(nil), // 1: black_list.GetAllHRRes
	(*Hr)(nil),          // 2: black_list.Hr
	(*ListUserReq)(nil), // 3: black_list.ListUserReq
	(*ListUserRes)(nil), // 4: black_list.ListUserRes
	(*UserRes)(nil),     // 5: black_list.UserRes
	(*Filter)(nil),      // 6: black_list.Filter
	(*GetById)(nil),     // 7: black_list.GetById
	(*Void)(nil),        // 8: black_list.Void
}
var file_admin_proto_depIdxs = []int32{
	2, // 0: black_list.GetAllHRRes.hr:type_name -> black_list.Hr
	6, // 1: black_list.ListUserReq.filter:type_name -> black_list.Filter
	5, // 2: black_list.ListUserRes.users:type_name -> black_list.UserRes
	0, // 3: black_list.AdminService.Approve:input_type -> black_list.CreateHR
	6, // 4: black_list.AdminService.ListHR:input_type -> black_list.Filter
	7, // 5: black_list.AdminService.Delete:input_type -> black_list.GetById
	3, // 6: black_list.AdminService.GetAllUsers:input_type -> black_list.ListUserReq
	8, // 7: black_list.AdminService.Approve:output_type -> black_list.Void
	1, // 8: black_list.AdminService.ListHR:output_type -> black_list.GetAllHRRes
	8, // 9: black_list.AdminService.Delete:output_type -> black_list.Void
	4, // 10: black_list.AdminService.GetAllUsers:output_type -> black_list.ListUserRes
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateHR); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllHRRes); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Hr); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserReq); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserRes); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserRes); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
