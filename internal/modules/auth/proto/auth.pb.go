// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: auth.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"` // JWT токен, если статус "success"
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	mi := &file_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"` // JWT токен, если статус "success"
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	mi := &file_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RegistrationResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a,
	0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x4d, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x14,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x32, 0x86, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []any{
	(*LoginRequest)(nil),         // 0: auth.LoginRequest
	(*LoginResponse)(nil),        // 1: auth.LoginResponse
	(*RegistrationRequest)(nil),  // 2: auth.RegistrationRequest
	(*RegistrationResponse)(nil), // 3: auth.RegistrationResponse
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.Login:input_type -> auth.LoginRequest
	2, // 1: auth.AuthService.Registration:input_type -> auth.RegistrationRequest
	1, // 2: auth.AuthService.Login:output_type -> auth.LoginResponse
	3, // 3: auth.AuthService.Registration:output_type -> auth.RegistrationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
