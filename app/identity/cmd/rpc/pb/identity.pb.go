// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: identity.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GenerateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GenerateTokenReq) Reset() {
	*x = GenerateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenReq) ProtoMessage() {}

func (x *GenerateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateTokenReq) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GenerateTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refreshAfter,proto3" json:"refreshAfter,omitempty"`
}

func (x *GenerateTokenResp) Reset() {
	*x = GenerateTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResp) ProtoMessage() {}

func (x *GenerateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResp.ProtoReflect.Descriptor instead.
func (*GenerateTokenResp) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GenerateTokenResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *GenerateTokenResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

type ClearTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ClearTokenReq) Reset() {
	*x = ClearTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearTokenReq) ProtoMessage() {}

func (x *ClearTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearTokenReq.ProtoReflect.Descriptor instead.
func (*ClearTokenReq) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{2}
}

func (x *ClearTokenReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ClearTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ClearTokenResp) Reset() {
	*x = ClearTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearTokenResp) ProtoMessage() {}

func (x *ClearTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearTokenResp.ProtoReflect.Descriptor instead.
func (*ClearTokenResp) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{3}
}

func (x *ClearTokenResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ValidateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateTokenReq) Reset() {
	*x = ValidateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenReq) ProtoMessage() {}

func (x *ValidateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenReq.ProtoReflect.Descriptor instead.
func (*ValidateTokenReq) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateTokenReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ValidateTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ValidateTokenResp) Reset() {
	*x = ValidateTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResp) ProtoMessage() {}

func (x *ValidateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResp.ProtoReflect.Descriptor instead.
func (*ValidateTokenResp) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateTokenResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_identity_proto protoreflect.FileDescriptor

var file_identity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x20,
	0x0a, 0x0e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x40, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xdf, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f,
	0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x48, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_proto_rawDescOnce sync.Once
	file_identity_proto_rawDescData = file_identity_proto_rawDesc
)

func file_identity_proto_rawDescGZIP() []byte {
	file_identity_proto_rawDescOnce.Do(func() {
		file_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_proto_rawDescData)
	})
	return file_identity_proto_rawDescData
}

var file_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_identity_proto_goTypes = []interface{}{
	(*GenerateTokenReq)(nil),  // 0: identity.GenerateTokenReq
	(*GenerateTokenResp)(nil), // 1: identity.GenerateTokenResp
	(*ClearTokenReq)(nil),     // 2: identity.ClearTokenReq
	(*ClearTokenResp)(nil),    // 3: identity.ClearTokenResp
	(*ValidateTokenReq)(nil),  // 4: identity.ValidateTokenReq
	(*ValidateTokenResp)(nil), // 5: identity.ValidateTokenResp
}
var file_identity_proto_depIdxs = []int32{
	0, // 0: identity.Identity.generateToken:input_type -> identity.GenerateTokenReq
	2, // 1: identity.Identity.clearToken:input_type -> identity.ClearTokenReq
	4, // 2: identity.Identity.validateToken:input_type -> identity.ValidateTokenReq
	1, // 3: identity.Identity.generateToken:output_type -> identity.GenerateTokenResp
	3, // 4: identity.Identity.clearToken:output_type -> identity.ClearTokenResp
	5, // 5: identity.Identity.validateToken:output_type -> identity.ValidateTokenResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identity_proto_init() }
func file_identity_proto_init() {
	if File_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenReq); i {
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
		file_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResp); i {
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
		file_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearTokenReq); i {
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
		file_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearTokenResp); i {
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
		file_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenReq); i {
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
		file_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenResp); i {
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
			RawDescriptor: file_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_proto_goTypes,
		DependencyIndexes: file_identity_proto_depIdxs,
		MessageInfos:      file_identity_proto_msgTypes,
	}.Build()
	File_identity_proto = out.File
	file_identity_proto_rawDesc = nil
	file_identity_proto_goTypes = nil
	file_identity_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	// 生成token, 只让用户服务访问
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	// 清理token, 只让用户服务访问
	ClearToken(ctx context.Context, in *ClearTokenReq, opts ...grpc.CallOption) (*ClearTokenResp, error)
	ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*ValidateTokenResp, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, "/identity.Identity/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ClearToken(ctx context.Context, in *ClearTokenReq, opts ...grpc.CallOption) (*ClearTokenResp, error) {
	out := new(ClearTokenResp)
	err := c.cc.Invoke(ctx, "/identity.Identity/clearToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*ValidateTokenResp, error) {
	out := new(ValidateTokenResp)
	err := c.cc.Invoke(ctx, "/identity.Identity/validateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	// 生成token, 只让用户服务访问
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	// 清理token, 只让用户服务访问
	ClearToken(context.Context, *ClearTokenReq) (*ClearTokenResp, error)
	ValidateToken(context.Context, *ValidateTokenReq) (*ValidateTokenResp, error)
}

// UnimplementedIdentityServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (*UnimplementedIdentityServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (*UnimplementedIdentityServer) ClearToken(context.Context, *ClearTokenReq) (*ClearTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearToken not implemented")
}
func (*UnimplementedIdentityServer) ValidateToken(context.Context, *ValidateTokenReq) (*ValidateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.Identity/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ClearToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ClearToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.Identity/ClearToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ClearToken(ctx, req.(*ClearTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.Identity/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ValidateToken(ctx, req.(*ValidateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identity.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateToken",
			Handler:    _Identity_GenerateToken_Handler,
		},
		{
			MethodName: "clearToken",
			Handler:    _Identity_ClearToken_Handler,
		},
		{
			MethodName: "validateToken",
			Handler:    _Identity_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}
