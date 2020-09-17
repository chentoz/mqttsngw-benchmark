// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: Config.proto

package Config

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StartFrequency  uint32   `protobuf:"fixed32,2,opt,name=StartFrequency,proto3" json:"StartFrequency,omitempty"`
	FrequencyStep   uint32   `protobuf:"fixed32,3,opt,name=FrequencyStep,proto3" json:"FrequencyStep,omitempty"`
	ControlChannel  uint32   `protobuf:"fixed32,4,opt,name=ControlChannel,proto3" json:"ControlChannel,omitempty"`
	ParentChannel   uint32   `protobuf:"fixed32,5,opt,name=ParentChannel,proto3" json:"ParentChannel,omitempty"`
	NeighborChannel []uint32 `protobuf:"fixed32,6,rep,packed,name=NeighborChannel,proto3" json:"NeighborChannel,omitempty"`
	UUID            string   `protobuf:"bytes,7,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_Config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_Config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Config) GetStartFrequency() uint32 {
	if x != nil {
		return x.StartFrequency
	}
	return 0
}

func (x *Config) GetFrequencyStep() uint32 {
	if x != nil {
		return x.FrequencyStep
	}
	return 0
}

func (x *Config) GetControlChannel() uint32 {
	if x != nil {
		return x.ControlChannel
	}
	return 0
}

func (x *Config) GetParentChannel() uint32 {
	if x != nil {
		return x.ParentChannel
	}
	return 0
}

func (x *Config) GetNeighborChannel() []uint32 {
	if x != nil {
		return x.NeighborChannel
	}
	return nil
}

func (x *Config) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

var File_Config_proto protoreflect.FileDescriptor

var file_Config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x07, 0x52, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74,
	0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0d, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x53, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x07, 0x52, 0x0f,
	0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Config_proto_rawDescOnce sync.Once
	file_Config_proto_rawDescData = file_Config_proto_rawDesc
)

func file_Config_proto_rawDescGZIP() []byte {
	file_Config_proto_rawDescOnce.Do(func() {
		file_Config_proto_rawDescData = protoimpl.X.CompressGZIP(file_Config_proto_rawDescData)
	})
	return file_Config_proto_rawDescData
}

var file_Config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: Config
}
var file_Config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Config_proto_init() }
func file_Config_proto_init() {
	if File_Config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_Config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Config_proto_goTypes,
		DependencyIndexes: file_Config_proto_depIdxs,
		MessageInfos:      file_Config_proto_msgTypes,
	}.Build()
	File_Config_proto = out.File
	file_Config_proto_rawDesc = nil
	file_Config_proto_goTypes = nil
	file_Config_proto_depIdxs = nil
}