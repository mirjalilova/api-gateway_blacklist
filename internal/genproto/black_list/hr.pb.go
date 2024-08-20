// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: hr.proto

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	FullName    string `protobuf:"bytes,3,opt,name=FullName,proto3" json:"FullName,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Position    string `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
	HrId        string `protobuf:"bytes,6,opt,name=hr_id,json=hrId,proto3" json:"hr_id,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Employee) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Employee) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Employee) GetHrId() string {
	if x != nil {
		return x.HrId
	}
	return ""
}

type EmployeeCreateBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *EmployeeCreateBody) Reset() {
	*x = EmployeeCreateBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeCreateBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeCreateBody) ProtoMessage() {}

func (x *EmployeeCreateBody) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeCreateBody.ProtoReflect.Descriptor instead.
func (*EmployeeCreateBody) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeCreateBody) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EmployeeCreateBody) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type EmployeeCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	HrId     string `protobuf:"bytes,4,opt,name=hr_id,json=hrId,proto3" json:"hr_id,omitempty"`
}

func (x *EmployeeCreate) Reset() {
	*x = EmployeeCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeCreate) ProtoMessage() {}

func (x *EmployeeCreate) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeCreate.ProtoReflect.Descriptor instead.
func (*EmployeeCreate) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{2}
}

func (x *EmployeeCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EmployeeCreate) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *EmployeeCreate) GetHrId() string {
	if x != nil {
		return x.HrId
	}
	return ""
}

type UpdateReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	HrId     string `protobuf:"bytes,2,opt,name=hr_id,json=hrId,proto3" json:"hr_id,omitempty"`
}

func (x *UpdateReqBody) Reset() {
	*x = UpdateReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReqBody) ProtoMessage() {}

func (x *UpdateReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReqBody.ProtoReflect.Descriptor instead.
func (*UpdateReqBody) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReqBody) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *UpdateReqBody) GetHrId() string {
	if x != nil {
		return x.HrId
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	HrId     string `protobuf:"bytes,3,opt,name=hr_id,json=hrId,proto3" json:"hr_id,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateReq) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *UpdateReq) GetHrId() string {
	if x != nil {
		return x.HrId
	}
	return ""
}

type ListEmployeeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position string  `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Filter   *Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListEmployeeReq) Reset() {
	*x = ListEmployeeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmployeeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmployeeReq) ProtoMessage() {}

func (x *ListEmployeeReq) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmployeeReq.ProtoReflect.Descriptor instead.
func (*ListEmployeeReq) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{5}
}

func (x *ListEmployeeReq) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *ListEmployeeReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListEmployeeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListEmployeeRes) Reset() {
	*x = ListEmployeeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmployeeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmployeeRes) ProtoMessage() {}

func (x *ListEmployeeRes) ProtoReflect() protoreflect.Message {
	mi := &file_hr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmployeeRes.ProtoReflect.Descriptor instead.
func (*ListEmployeeRes) Descriptor() ([]byte, []int) {
	return file_hr_proto_rawDescGZIP(), []int{6}
}

func (x *ListEmployeeRes) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

func (x *ListEmployeeRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_hr_proto protoreflect.FileDescriptor

var file_hr_proto_rawDesc = []byte{
	0x0a, 0x08, 0x68, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x13, 0x0a, 0x05, 0x68, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5a, 0x0a, 0x0e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x68, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x68, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x49, 0x64, 0x22,
	0x4c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x68, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa7, 0x02, 0x0a, 0x09, 0x48, 0x52, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42,
	0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hr_proto_rawDescOnce sync.Once
	file_hr_proto_rawDescData = file_hr_proto_rawDesc
)

func file_hr_proto_rawDescGZIP() []byte {
	file_hr_proto_rawDescOnce.Do(func() {
		file_hr_proto_rawDescData = protoimpl.X.CompressGZIP(file_hr_proto_rawDescData)
	})
	return file_hr_proto_rawDescData
}

var file_hr_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hr_proto_goTypes = []any{
	(*Employee)(nil),           // 0: black_list.Employee
	(*EmployeeCreateBody)(nil), // 1: black_list.EmployeeCreateBody
	(*EmployeeCreate)(nil),     // 2: black_list.EmployeeCreate
	(*UpdateReqBody)(nil),      // 3: black_list.UpdateReqBody
	(*UpdateReq)(nil),          // 4: black_list.UpdateReq
	(*ListEmployeeReq)(nil),    // 5: black_list.ListEmployeeReq
	(*ListEmployeeRes)(nil),    // 6: black_list.ListEmployeeRes
	(*Filter)(nil),             // 7: black_list.Filter
	(*GetById)(nil),            // 8: black_list.GetById
	(*Void)(nil),               // 9: black_list.Void
}
var file_hr_proto_depIdxs = []int32{
	7, // 0: black_list.ListEmployeeReq.filter:type_name -> black_list.Filter
	0, // 1: black_list.ListEmployeeRes.employees:type_name -> black_list.Employee
	2, // 2: black_list.HRService.Create:input_type -> black_list.EmployeeCreate
	8, // 3: black_list.HRService.Get:input_type -> black_list.GetById
	5, // 4: black_list.HRService.GetAll:input_type -> black_list.ListEmployeeReq
	4, // 5: black_list.HRService.Update:input_type -> black_list.UpdateReq
	8, // 6: black_list.HRService.Delete:input_type -> black_list.GetById
	9, // 7: black_list.HRService.Create:output_type -> black_list.Void
	0, // 8: black_list.HRService.Get:output_type -> black_list.Employee
	6, // 9: black_list.HRService.GetAll:output_type -> black_list.ListEmployeeRes
	9, // 10: black_list.HRService.Update:output_type -> black_list.Void
	9, // 11: black_list.HRService.Delete:output_type -> black_list.Void
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hr_proto_init() }
func file_hr_proto_init() {
	if File_hr_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hr_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Employee); i {
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
		file_hr_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EmployeeCreateBody); i {
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
		file_hr_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EmployeeCreate); i {
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
		file_hr_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateReqBody); i {
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
		file_hr_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateReq); i {
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
		file_hr_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListEmployeeReq); i {
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
		file_hr_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListEmployeeRes); i {
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
			RawDescriptor: file_hr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hr_proto_goTypes,
		DependencyIndexes: file_hr_proto_depIdxs,
		MessageInfos:      file_hr_proto_msgTypes,
	}.Build()
	File_hr_proto = out.File
	file_hr_proto_rawDesc = nil
	file_hr_proto_goTypes = nil
	file_hr_proto_depIdxs = nil
}
