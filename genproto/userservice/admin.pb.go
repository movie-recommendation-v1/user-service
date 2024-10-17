// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: protos/user-service/admin.proto

package userservice

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

type AdminModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Role      string `protobuf:"bytes,8,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *AdminModel) Reset() {
	*x = AdminModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminModel) ProtoMessage() {}

func (x *AdminModel) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminModel.ProtoReflect.Descriptor instead.
func (*AdminModel) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AdminModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminModel) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *AdminModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AdminModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AdminModel) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminName     string `protobuf:"bytes,1,opt,name=admin_name,json=adminName,proto3" json:"admin_name,omitempty"`
	AdminEmail    string `protobuf:"bytes,2,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
	AdminPassword string `protobuf:"bytes,3,opt,name=admin_password,json=adminPassword,proto3" json:"admin_password,omitempty"`
}

func (x *CreateAdminReq) Reset() {
	*x = CreateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReq) ProtoMessage() {}

func (x *CreateAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminReq.ProtoReflect.Descriptor instead.
func (*CreateAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdminReq) GetAdminName() string {
	if x != nil {
		return x.AdminName
	}
	return ""
}

func (x *CreateAdminReq) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

func (x *CreateAdminReq) GetAdminPassword() string {
	if x != nil {
		return x.AdminPassword
	}
	return ""
}

type CreateAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateAdminRes) Reset() {
	*x = CreateAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminRes) ProtoMessage() {}

func (x *CreateAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminRes.ProtoReflect.Descriptor instead.
func (*CreateAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAdminRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteAdminCreateAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAdminCreateAdminRes) Reset() {
	*x = DeleteAdminCreateAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminCreateAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminCreateAdminRes) ProtoMessage() {}

func (x *DeleteAdminCreateAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminCreateAdminRes.ProtoReflect.Descriptor instead.
func (*DeleteAdminCreateAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAdminCreateAdminRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminReq *AdminModel `protobuf:"bytes,1,opt,name=admin_req,json=adminReq,proto3" json:"admin_req,omitempty"`
}

func (x *UpdateAdminReq) Reset() {
	*x = UpdateAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReq) ProtoMessage() {}

func (x *UpdateAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminReq.ProtoReflect.Descriptor instead.
func (*UpdateAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAdminReq) GetAdminReq() *AdminModel {
	if x != nil {
		return x.AdminReq
	}
	return nil
}

type UpdateAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminRes *AdminModel `protobuf:"bytes,1,opt,name=admin_res,json=adminRes,proto3" json:"admin_res,omitempty"`
}

func (x *UpdateAdminRes) Reset() {
	*x = UpdateAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminRes) ProtoMessage() {}

func (x *UpdateAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminRes.ProtoReflect.Descriptor instead.
func (*UpdateAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAdminRes) GetAdminRes() *AdminModel {
	if x != nil {
		return x.AdminRes
	}
	return nil
}

type GetAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId string `protobuf:"bytes,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
}

func (x *GetAdminReq) Reset() {
	*x = GetAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReq) ProtoMessage() {}

func (x *GetAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminReq.ProtoReflect.Descriptor instead.
func (*GetAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{6}
}

func (x *GetAdminReq) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

type GetAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminRes *AdminModel `protobuf:"bytes,1,opt,name=admin_res,json=adminRes,proto3" json:"admin_res,omitempty"`
}

func (x *GetAdminRes) Reset() {
	*x = GetAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRes) ProtoMessage() {}

func (x *GetAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRes.ProtoReflect.Descriptor instead.
func (*GetAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{7}
}

func (x *GetAdminRes) GetAdminRes() *AdminModel {
	if x != nil {
		return x.AdminRes
	}
	return nil
}

type ForgetPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminEmail string `protobuf:"bytes,1,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
}

func (x *ForgetPasswordReq) Reset() {
	*x = ForgetPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordReq) ProtoMessage() {}

func (x *ForgetPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordReq.ProtoReflect.Descriptor instead.
func (*ForgetPasswordReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{8}
}

func (x *ForgetPasswordReq) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

type ForgetPasswordRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ForgetPasswordRes) Reset() {
	*x = ForgetPasswordRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordRes) ProtoMessage() {}

func (x *ForgetPasswordRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordRes.ProtoReflect.Descriptor instead.
func (*ForgetPasswordRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{9}
}

func (x *ForgetPasswordRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminReq *AdminModel `protobuf:"bytes,1,opt,name=admin_req,json=adminReq,proto3" json:"admin_req,omitempty"`
	Limit    int32       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32       `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetAllAdminReq) Reset() {
	*x = GetAllAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAdminReq) ProtoMessage() {}

func (x *GetAllAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAdminReq.ProtoReflect.Descriptor instead.
func (*GetAllAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllAdminReq) GetAdminReq() *AdminModel {
	if x != nil {
		return x.AdminReq
	}
	return nil
}

func (x *GetAllAdminReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllAdminReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminRes []*AdminModel `protobuf:"bytes,1,rep,name=admin_res,json=adminRes,proto3" json:"admin_res,omitempty"`
}

func (x *GetAllAdminRes) Reset() {
	*x = GetAllAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAdminRes) ProtoMessage() {}

func (x *GetAllAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAdminRes.ProtoReflect.Descriptor instead.
func (*GetAllAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{11}
}

func (x *GetAllAdminRes) GetAdminRes() []*AdminModel {
	if x != nil {
		return x.AdminRes
	}
	return nil
}

type DeleteAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId string `protobuf:"bytes,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
}

func (x *DeleteAdminReq) Reset() {
	*x = DeleteAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReq) ProtoMessage() {}

func (x *DeleteAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminReq.ProtoReflect.Descriptor instead.
func (*DeleteAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteAdminReq) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

type DeleteAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAdminRes) Reset() {
	*x = DeleteAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_admin_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminRes) ProtoMessage() {}

func (x *DeleteAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_admin_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminRes.ProtoReflect.Descriptor instead.
func (*DeleteAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_admin_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteAdminRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_user_service_admin_proto protoreflect.FileDescriptor

var file_protos_user_service_admin_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xd0,
	0x01, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x77, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x34, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x28, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x11,
	0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2d, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22,
	0x2b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc4, 0x03, 0x0a, 0x0c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x47, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0e, 0x46,
	0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x67,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x67,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x42,
	0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_user_service_admin_proto_rawDescOnce sync.Once
	file_protos_user_service_admin_proto_rawDescData = file_protos_user_service_admin_proto_rawDesc
)

func file_protos_user_service_admin_proto_rawDescGZIP() []byte {
	file_protos_user_service_admin_proto_rawDescOnce.Do(func() {
		file_protos_user_service_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_service_admin_proto_rawDescData)
	})
	return file_protos_user_service_admin_proto_rawDescData
}

var file_protos_user_service_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_protos_user_service_admin_proto_goTypes = []interface{}{
	(*AdminModel)(nil),                // 0: userservice.AdminModel
	(*CreateAdminReq)(nil),            // 1: userservice.CreateAdminReq
	(*CreateAdminRes)(nil),            // 2: userservice.CreateAdminRes
	(*DeleteAdminCreateAdminRes)(nil), // 3: userservice.DeleteAdminCreateAdminRes
	(*UpdateAdminReq)(nil),            // 4: userservice.UpdateAdminReq
	(*UpdateAdminRes)(nil),            // 5: userservice.UpdateAdminRes
	(*GetAdminReq)(nil),               // 6: userservice.GetAdminReq
	(*GetAdminRes)(nil),               // 7: userservice.GetAdminRes
	(*ForgetPasswordReq)(nil),         // 8: userservice.ForgetPasswordReq
	(*ForgetPasswordRes)(nil),         // 9: userservice.ForgetPasswordRes
	(*GetAllAdminReq)(nil),            // 10: userservice.GetAllAdminReq
	(*GetAllAdminRes)(nil),            // 11: userservice.GetAllAdminRes
	(*DeleteAdminReq)(nil),            // 12: userservice.DeleteAdminReq
	(*DeleteAdminRes)(nil),            // 13: userservice.DeleteAdminRes
}
var file_protos_user_service_admin_proto_depIdxs = []int32{
	0,  // 0: userservice.UpdateAdminReq.admin_req:type_name -> userservice.AdminModel
	0,  // 1: userservice.UpdateAdminRes.admin_res:type_name -> userservice.AdminModel
	0,  // 2: userservice.GetAdminRes.admin_res:type_name -> userservice.AdminModel
	0,  // 3: userservice.GetAllAdminReq.admin_req:type_name -> userservice.AdminModel
	0,  // 4: userservice.GetAllAdminRes.admin_res:type_name -> userservice.AdminModel
	1,  // 5: userservice.AdminService.CreateAdmin:input_type -> userservice.CreateAdminReq
	4,  // 6: userservice.AdminService.UpdateAdmin:input_type -> userservice.UpdateAdminReq
	6,  // 7: userservice.AdminService.GetAdmin:input_type -> userservice.GetAdminReq
	8,  // 8: userservice.AdminService.ForgetPassword:input_type -> userservice.ForgetPasswordReq
	10, // 9: userservice.AdminService.GetAllAdmin:input_type -> userservice.GetAllAdminReq
	12, // 10: userservice.AdminService.DeleteAdmin:input_type -> userservice.DeleteAdminReq
	2,  // 11: userservice.AdminService.CreateAdmin:output_type -> userservice.CreateAdminRes
	5,  // 12: userservice.AdminService.UpdateAdmin:output_type -> userservice.UpdateAdminRes
	7,  // 13: userservice.AdminService.GetAdmin:output_type -> userservice.GetAdminRes
	9,  // 14: userservice.AdminService.ForgetPassword:output_type -> userservice.ForgetPasswordRes
	11, // 15: userservice.AdminService.GetAllAdmin:output_type -> userservice.GetAllAdminRes
	13, // 16: userservice.AdminService.DeleteAdmin:output_type -> userservice.DeleteAdminRes
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_user_service_admin_proto_init() }
func file_protos_user_service_admin_proto_init() {
	if File_protos_user_service_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_service_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminModel); i {
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
		file_protos_user_service_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminCreateAdminRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetPasswordReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetPasswordRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAdminReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAdminRes); i {
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
		file_protos_user_service_admin_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminReq); i {
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
		file_protos_user_service_admin_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminRes); i {
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
			RawDescriptor: file_protos_user_service_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_service_admin_proto_goTypes,
		DependencyIndexes: file_protos_user_service_admin_proto_depIdxs,
		MessageInfos:      file_protos_user_service_admin_proto_msgTypes,
	}.Build()
	File_protos_user_service_admin_proto = out.File
	file_protos_user_service_admin_proto_rawDesc = nil
	file_protos_user_service_admin_proto_goTypes = nil
	file_protos_user_service_admin_proto_depIdxs = nil
}
