// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: protos/user-service/user.proto

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

type UserModel struct {
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

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserModel) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *UserModel) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type RegisterUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterUserReq) Reset() {
	*x = RegisterUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserReq) ProtoMessage() {}

func (x *RegisterUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserReq.ProtoReflect.Descriptor instead.
func (*RegisterUserReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRes *UserModel `protobuf:"bytes,1,opt,name=UserRes,proto3" json:"UserRes,omitempty"`
}

func (x *RegisterUserRes) Reset() {
	*x = RegisterUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRes) ProtoMessage() {}

func (x *RegisterUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRes.ProtoReflect.Descriptor instead.
func (*RegisterUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterUserRes) GetUserRes() *UserModel {
	if x != nil {
		return x.UserRes
	}
	return nil
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRes *UserModel `protobuf:"bytes,1,opt,name=UserRes,proto3" json:"UserRes,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRes) GetUserRes() *UserModel {
	if x != nil {
		return x.UserRes
	}
	return nil
}

type ForgotPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ForgotPasswordReq) Reset() {
	*x = ForgotPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordReq) ProtoMessage() {}

func (x *ForgotPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordReq.ProtoReflect.Descriptor instead.
func (*ForgotPasswordReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{5}
}

func (x *ForgotPasswordReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ForgotPasswordRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ForgotPasswordRes) Reset() {
	*x = ForgotPasswordRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRes) ProtoMessage() {}

func (x *ForgotPasswordRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordRes.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{6}
}

func (x *ForgotPasswordRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetUserByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *GetUserByIDReq) Reset() {
	*x = GetUserByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDReq) ProtoMessage() {}

func (x *GetUserByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDReq.ProtoReflect.Descriptor instead.
func (*GetUserByIDReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserByIDReq) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type GetUserByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRes *UserModel `protobuf:"bytes,1,opt,name=UserRes,proto3" json:"UserRes,omitempty"`
}

func (x *GetUserByIDRes) Reset() {
	*x = GetUserByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRes) ProtoMessage() {}

func (x *GetUserByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRes.ProtoReflect.Descriptor instead.
func (*GetUserByIDRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserByIDRes) GetUserRes() *UserModel {
	if x != nil {
		return x.UserRes
	}
	return nil
}

type GetAllUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserReq *UserModel `protobuf:"bytes,1,opt,name=UserReq,proto3" json:"UserReq,omitempty"`
	Limit   int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32      `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetAllUserReq) Reset() {
	*x = GetAllUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserReq) ProtoMessage() {}

func (x *GetAllUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserReq.ProtoReflect.Descriptor instead.
func (*GetAllUserReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllUserReq) GetUserReq() *UserModel {
	if x != nil {
		return x.UserReq
	}
	return nil
}

func (x *GetAllUserReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllUserReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRes []*UserModel `protobuf:"bytes,1,rep,name=UserRes,proto3" json:"UserRes,omitempty"`
}

func (x *GetAllUserRes) Reset() {
	*x = GetAllUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRes) ProtoMessage() {}

func (x *GetAllUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRes.ProtoReflect.Descriptor instead.
func (*GetAllUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllUserRes) GetUserRes() []*UserModel {
	if x != nil {
		return x.UserRes
	}
	return nil
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserReq *UserModel `protobuf:"bytes,1,opt,name=UserReq,proto3" json:"UserReq,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserReq) GetUserReq() *UserModel {
	if x != nil {
		return x.UserReq
	}
	return nil
}

type UpdateUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRes *UserModel `protobuf:"bytes,1,opt,name=UserRes,proto3" json:"UserRes,omitempty"`
}

func (x *UpdateUserRes) Reset() {
	*x = UpdateUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRes) ProtoMessage() {}

func (x *UpdateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRes.ProtoReflect.Descriptor instead.
func (*UpdateUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateUserRes) GetUserRes() *UserModel {
	if x != nil {
		return x.UserRes
	}
	return nil
}

type DeleteUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteUserReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteUserRes) Reset() {
	*x = DeleteUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_service_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRes) ProtoMessage() {}

func (x *DeleteUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_service_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRes.ProtoReflect.Descriptor instead.
func (*DeleteUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_service_user_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteUserRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_user_service_user_proto protoreflect.FileDescriptor

var file_protos_user_service_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xcf, 0x01,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x57, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x3c, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x11, 0x46, 0x6f, 0x72,
	0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2d, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x22, 0x6f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x41, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfd, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0e, 0x46, 0x6f,
	0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_user_service_user_proto_rawDescOnce sync.Once
	file_protos_user_service_user_proto_rawDescData = file_protos_user_service_user_proto_rawDesc
)

func file_protos_user_service_user_proto_rawDescGZIP() []byte {
	file_protos_user_service_user_proto_rawDescOnce.Do(func() {
		file_protos_user_service_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_service_user_proto_rawDescData)
	})
	return file_protos_user_service_user_proto_rawDescData
}

var file_protos_user_service_user_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_protos_user_service_user_proto_goTypes = []interface{}{
	(*UserModel)(nil),         // 0: userservice.UserModel
	(*RegisterUserReq)(nil),   // 1: userservice.RegisterUserReq
	(*RegisterUserRes)(nil),   // 2: userservice.RegisterUserRes
	(*LoginReq)(nil),          // 3: userservice.LoginReq
	(*LoginRes)(nil),          // 4: userservice.LoginRes
	(*ForgotPasswordReq)(nil), // 5: userservice.ForgotPasswordReq
	(*ForgotPasswordRes)(nil), // 6: userservice.ForgotPasswordRes
	(*GetUserByIDReq)(nil),    // 7: userservice.GetUserByIDReq
	(*GetUserByIDRes)(nil),    // 8: userservice.GetUserByIDRes
	(*GetAllUserReq)(nil),     // 9: userservice.GetAllUserReq
	(*GetAllUserRes)(nil),     // 10: userservice.GetAllUserRes
	(*UpdateUserReq)(nil),     // 11: userservice.UpdateUserReq
	(*UpdateUserRes)(nil),     // 12: userservice.UpdateUserRes
	(*DeleteUserReq)(nil),     // 13: userservice.DeleteUserReq
	(*DeleteUserRes)(nil),     // 14: userservice.DeleteUserRes
}
var file_protos_user_service_user_proto_depIdxs = []int32{
	0,  // 0: userservice.RegisterUserRes.UserRes:type_name -> userservice.UserModel
	0,  // 1: userservice.LoginRes.UserRes:type_name -> userservice.UserModel
	0,  // 2: userservice.GetUserByIDRes.UserRes:type_name -> userservice.UserModel
	0,  // 3: userservice.GetAllUserReq.UserReq:type_name -> userservice.UserModel
	0,  // 4: userservice.GetAllUserRes.UserRes:type_name -> userservice.UserModel
	0,  // 5: userservice.UpdateUserReq.UserReq:type_name -> userservice.UserModel
	0,  // 6: userservice.UpdateUserRes.UserRes:type_name -> userservice.UserModel
	1,  // 7: userservice.UserService.RegisterUser:input_type -> userservice.RegisterUserReq
	3,  // 8: userservice.UserService.Login:input_type -> userservice.LoginReq
	5,  // 9: userservice.UserService.ForgotPassword:input_type -> userservice.ForgotPasswordReq
	7,  // 10: userservice.UserService.GetUserByID:input_type -> userservice.GetUserByIDReq
	9,  // 11: userservice.UserService.GetAllUser:input_type -> userservice.GetAllUserReq
	11, // 12: userservice.UserService.UpdateUser:input_type -> userservice.UpdateUserReq
	13, // 13: userservice.UserService.DeleteUser:input_type -> userservice.DeleteUserReq
	2,  // 14: userservice.UserService.RegisterUser:output_type -> userservice.RegisterUserRes
	4,  // 15: userservice.UserService.Login:output_type -> userservice.LoginRes
	6,  // 16: userservice.UserService.ForgotPassword:output_type -> userservice.ForgotPasswordRes
	8,  // 17: userservice.UserService.GetUserByID:output_type -> userservice.GetUserByIDRes
	10, // 18: userservice.UserService.GetAllUser:output_type -> userservice.GetAllUserRes
	12, // 19: userservice.UserService.UpdateUser:output_type -> userservice.UpdateUserRes
	14, // 20: userservice.UserService.DeleteUser:output_type -> userservice.DeleteUserRes
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_user_service_user_proto_init() }
func file_protos_user_service_user_proto_init() {
	if File_protos_user_service_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_service_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_protos_user_service_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserReq); i {
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
		file_protos_user_service_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserRes); i {
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
		file_protos_user_service_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_protos_user_service_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_protos_user_service_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordReq); i {
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
		file_protos_user_service_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordRes); i {
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
		file_protos_user_service_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDReq); i {
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
		file_protos_user_service_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDRes); i {
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
		file_protos_user_service_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserReq); i {
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
		file_protos_user_service_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserRes); i {
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
		file_protos_user_service_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
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
		file_protos_user_service_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRes); i {
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
		file_protos_user_service_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReq); i {
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
		file_protos_user_service_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRes); i {
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
			RawDescriptor: file_protos_user_service_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_service_user_proto_goTypes,
		DependencyIndexes: file_protos_user_service_user_proto_depIdxs,
		MessageInfos:      file_protos_user_service_user_proto_msgTypes,
	}.Build()
	File_protos_user_service_user_proto = out.File
	file_protos_user_service_user_proto_rawDesc = nil
	file_protos_user_service_user_proto_goTypes = nil
	file_protos_user_service_user_proto_depIdxs = nil
}
