// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: newOrg.proto

package server

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

type Msp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultHost string `protobuf:"bytes,1,opt,name=vaultHost,proto3" json:"vaultHost,omitempty"`
	Org       string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
	OutDir    string `protobuf:"bytes,3,opt,name=outDir,proto3" json:"outDir,omitempty"`
}

func (x *Msp) Reset() {
	*x = Msp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newOrg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msp) ProtoMessage() {}

func (x *Msp) ProtoReflect() protoreflect.Message {
	mi := &file_newOrg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msp.ProtoReflect.Descriptor instead.
func (*Msp) Descriptor() ([]byte, []int) {
	return file_newOrg_proto_rawDescGZIP(), []int{0}
}

func (x *Msp) GetVaultHost() string {
	if x != nil {
		return x.VaultHost
	}
	return ""
}

func (x *Msp) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *Msp) GetOutDir() string {
	if x != nil {
		return x.OutDir
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newOrg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_newOrg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_newOrg_proto_rawDescGZIP(), []int{1}
}

var File_newOrg_proto protoreflect.FileDescriptor

var file_newOrg_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x4f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x03, 0x4d, 0x73, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x72, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x44, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x44, 0x69, 0x72, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x32, 0x0a,
	0x06, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x67, 0x12, 0x28, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x53, 0x50, 0x12, 0x0b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x73,
	0x70, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newOrg_proto_rawDescOnce sync.Once
	file_newOrg_proto_rawDescData = file_newOrg_proto_rawDesc
)

func file_newOrg_proto_rawDescGZIP() []byte {
	file_newOrg_proto_rawDescOnce.Do(func() {
		file_newOrg_proto_rawDescData = protoimpl.X.CompressGZIP(file_newOrg_proto_rawDescData)
	})
	return file_newOrg_proto_rawDescData
}

var file_newOrg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_newOrg_proto_goTypes = []interface{}{
	(*Msp)(nil),  // 0: server.Msp
	(*Void)(nil), // 1: server.Void
}
var file_newOrg_proto_depIdxs = []int32{
	0, // 0: server.NewOrg.CreateMSP:input_type -> server.Msp
	1, // 1: server.NewOrg.CreateMSP:output_type -> server.Void
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newOrg_proto_init() }
func file_newOrg_proto_init() {
	if File_newOrg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newOrg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msp); i {
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
		file_newOrg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_newOrg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_newOrg_proto_goTypes,
		DependencyIndexes: file_newOrg_proto_depIdxs,
		MessageInfos:      file_newOrg_proto_msgTypes,
	}.Build()
	File_newOrg_proto = out.File
	file_newOrg_proto_rawDesc = nil
	file_newOrg_proto_goTypes = nil
	file_newOrg_proto_depIdxs = nil
}