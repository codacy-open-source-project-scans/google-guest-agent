// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: telemetry.proto

package telemetry

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

type OSInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OsType        *string `protobuf:"bytes,1,opt,name=os_type,json=osType,proto3,oneof" json:"os_type,omitempty"`
	LongName      *string `protobuf:"bytes,2,opt,name=long_name,json=longName,proto3,oneof" json:"long_name,omitempty"`
	ShortName     *string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3,oneof" json:"short_name,omitempty"`
	Version       *string `protobuf:"bytes,4,opt,name=version,proto3,oneof" json:"version,omitempty"`
	KernelVersion *string `protobuf:"bytes,5,opt,name=kernel_version,json=kernelVersion,proto3,oneof" json:"kernel_version,omitempty"`
	KernelRelease *string `protobuf:"bytes,6,opt,name=kernel_release,json=kernelRelease,proto3,oneof" json:"kernel_release,omitempty"`
}

func (x *OSInfo) Reset() {
	*x = OSInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSInfo) ProtoMessage() {}

func (x *OSInfo) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSInfo.ProtoReflect.Descriptor instead.
func (*OSInfo) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *OSInfo) GetOsType() string {
	if x != nil && x.OsType != nil {
		return *x.OsType
	}
	return ""
}

func (x *OSInfo) GetLongName() string {
	if x != nil && x.LongName != nil {
		return *x.LongName
	}
	return ""
}

func (x *OSInfo) GetShortName() string {
	if x != nil && x.ShortName != nil {
		return *x.ShortName
	}
	return ""
}

func (x *OSInfo) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *OSInfo) GetKernelVersion() string {
	if x != nil && x.KernelVersion != nil {
		return *x.KernelVersion
	}
	return ""
}

func (x *OSInfo) GetKernelRelease() string {
	if x != nil && x.KernelRelease != nil {
		return *x.KernelRelease
	}
	return ""
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Version      *string `protobuf:"bytes,2,opt,name=version,proto3,oneof" json:"version,omitempty"`
	Architecture *string `protobuf:"bytes,3,opt,name=architecture,proto3,oneof" json:"architecture,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *AgentInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AgentInfo) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *AgentInfo) GetArchitecture() string {
	if x != nil && x.Architecture != nil {
		return *x.Architecture
	}
	return ""
}

var File_telemetry_proto protoreflect.FileDescriptor

var file_telemetry_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0xbe, 0x02, 0x0a,
	0x06, 0x4f, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x92, 0x01,
	0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_proto_rawDescOnce sync.Once
	file_telemetry_proto_rawDescData = file_telemetry_proto_rawDesc
)

func file_telemetry_proto_rawDescGZIP() []byte {
	file_telemetry_proto_rawDescOnce.Do(func() {
		file_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_proto_rawDescData)
	})
	return file_telemetry_proto_rawDescData
}

var file_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_telemetry_proto_goTypes = []interface{}{
	(*OSInfo)(nil),    // 0: telemetry.OSInfo
	(*AgentInfo)(nil), // 1: telemetry.AgentInfo
}
var file_telemetry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_telemetry_proto_init() }
func file_telemetry_proto_init() {
	if File_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSInfo); i {
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
		file_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
	file_telemetry_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_telemetry_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telemetry_proto_goTypes,
		DependencyIndexes: file_telemetry_proto_depIdxs,
		MessageInfos:      file_telemetry_proto_msgTypes,
	}.Build()
	File_telemetry_proto = out.File
	file_telemetry_proto_rawDesc = nil
	file_telemetry_proto_goTypes = nil
	file_telemetry_proto_depIdxs = nil
}
