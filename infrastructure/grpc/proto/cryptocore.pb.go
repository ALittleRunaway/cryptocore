// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: infrastructure/grpc/proto/cryptocore.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Rule  int64  `protobuf:"varint,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *MineRequest) Reset() {
	*x = MineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MineRequest) ProtoMessage() {}

func (x *MineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MineRequest.ProtoReflect.Descriptor instead.
func (*MineRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_grpc_proto_cryptocore_proto_rawDescGZIP(), []int{0}
}

func (x *MineRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *MineRequest) GetRule() int64 {
	if x != nil {
		return x.Rule
	}
	return 0
}

type MineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pow int64 `protobuf:"varint,1,opt,name=pow,proto3" json:"pow,omitempty"`
}

func (x *MineResponse) Reset() {
	*x = MineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MineResponse) ProtoMessage() {}

func (x *MineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MineResponse.ProtoReflect.Descriptor instead.
func (*MineResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_grpc_proto_cryptocore_proto_rawDescGZIP(), []int{1}
}

func (x *MineResponse) GetPow() int64 {
	if x != nil {
		return x.Pow
	}
	return 0
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Pow   int64  `protobuf:"varint,3,opt,name=pow,proto3" json:"pow,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_grpc_proto_cryptocore_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *ValidateRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ValidateRequest) GetPow() int64 {
	if x != nil {
		return x.Pow
	}
	return 0
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_grpc_proto_cryptocore_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_infrastructure_grpc_proto_cryptocore_proto protoreflect.FileDescriptor

var file_infrastructure_grpc_proto_cryptocore_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x68,
	0x6f, 0x75, 0x72, 0x75, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x6f, 0x72, 0x65,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x6f, 0x77, 0x22, 0x4d, 0x0a, 0x0f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x6f, 0x77, 0x22, 0x28, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x32, 0xb2, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x6f, 0x72,
	0x65, 0x12, 0x4b, 0x0a, 0x04, 0x4d, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x68, 0x6f, 0x75,
	0x72, 0x75, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x68, 0x6f,
	0x75, 0x72, 0x75, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x68, 0x6f,
	0x75, 0x72, 0x75, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x72, 0x75, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_grpc_proto_cryptocore_proto_rawDescOnce sync.Once
	file_infrastructure_grpc_proto_cryptocore_proto_rawDescData = file_infrastructure_grpc_proto_cryptocore_proto_rawDesc
)

func file_infrastructure_grpc_proto_cryptocore_proto_rawDescGZIP() []byte {
	file_infrastructure_grpc_proto_cryptocore_proto_rawDescOnce.Do(func() {
		file_infrastructure_grpc_proto_cryptocore_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_grpc_proto_cryptocore_proto_rawDescData)
	})
	return file_infrastructure_grpc_proto_cryptocore_proto_rawDescData
}

var file_infrastructure_grpc_proto_cryptocore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infrastructure_grpc_proto_cryptocore_proto_goTypes = []interface{}{
	(*MineRequest)(nil),      // 0: thourus.cryptocore.MineRequest
	(*MineResponse)(nil),     // 1: thourus.cryptocore.MineResponse
	(*ValidateRequest)(nil),  // 2: thourus.cryptocore.ValidateRequest
	(*ValidateResponse)(nil), // 3: thourus.cryptocore.ValidateResponse
}
var file_infrastructure_grpc_proto_cryptocore_proto_depIdxs = []int32{
	0, // 0: thourus.cryptocore.CryptoCore.Mine:input_type -> thourus.cryptocore.MineRequest
	2, // 1: thourus.cryptocore.CryptoCore.Validate:input_type -> thourus.cryptocore.ValidateRequest
	1, // 2: thourus.cryptocore.CryptoCore.Mine:output_type -> thourus.cryptocore.MineResponse
	3, // 3: thourus.cryptocore.CryptoCore.Validate:output_type -> thourus.cryptocore.ValidateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infrastructure_grpc_proto_cryptocore_proto_init() }
func file_infrastructure_grpc_proto_cryptocore_proto_init() {
	if File_infrastructure_grpc_proto_cryptocore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MineRequest); i {
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
		file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MineResponse); i {
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
		file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_infrastructure_grpc_proto_cryptocore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
			RawDescriptor: file_infrastructure_grpc_proto_cryptocore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_grpc_proto_cryptocore_proto_goTypes,
		DependencyIndexes: file_infrastructure_grpc_proto_cryptocore_proto_depIdxs,
		MessageInfos:      file_infrastructure_grpc_proto_cryptocore_proto_msgTypes,
	}.Build()
	File_infrastructure_grpc_proto_cryptocore_proto = out.File
	file_infrastructure_grpc_proto_cryptocore_proto_rawDesc = nil
	file_infrastructure_grpc_proto_cryptocore_proto_goTypes = nil
	file_infrastructure_grpc_proto_cryptocore_proto_depIdxs = nil
}
