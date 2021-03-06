// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: BLELocation.proto

package BLELocation

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

type BLELocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp       uint32  `protobuf:"fixed32,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                    // 定位消息的时间戳，单位0.001秒，Unix Epoch Time, UTC时区
	Id              []byte  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                                    //定位卡唯一ID
	MajorMinorArray []byte  `protobuf:"bytes,3,opt,name=major_minor_array,json=majorMinorArray,proto3" json:"major_minor_array,omitempty"` // 4个字节一组，共4组，每组格式major 2字节，minor 2字节，无效的用0表示
	AccuracyArray   []byte  `protobuf:"bytes,4,opt,name=accuracy_array,json=accuracyArray,proto3" json:"accuracy_array,omitempty"`         // BLE 1距离精度，无效用负值表示，典型值为-1
	Barometer       float32 `protobuf:"fixed32,5,opt,name=barometer,proto3" json:"barometer,omitempty"`                                    //定位卡电量：[0~1]，1表示100%，0.5表示50%
	Status          []byte  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`                                            // battery(uint8) / temperature(int8) / Bits : SOS + motion
}

func (x *BLELocation) Reset() {
	*x = BLELocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BLELocation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLELocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLELocation) ProtoMessage() {}

func (x *BLELocation) ProtoReflect() protoreflect.Message {
	mi := &file_BLELocation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLELocation.ProtoReflect.Descriptor instead.
func (*BLELocation) Descriptor() ([]byte, []int) {
	return file_BLELocation_proto_rawDescGZIP(), []int{0}
}

func (x *BLELocation) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BLELocation) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BLELocation) GetMajorMinorArray() []byte {
	if x != nil {
		return x.MajorMinorArray
	}
	return nil
}

func (x *BLELocation) GetAccuracyArray() []byte {
	if x != nil {
		return x.AccuracyArray
	}
	return nil
}

func (x *BLELocation) GetBarometer() float32 {
	if x != nil {
		return x.Barometer
	}
	return 0
}

func (x *BLELocation) GetStatus() []byte {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_BLELocation_proto protoreflect.FileDescriptor

var file_BLELocation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x4c, 0x45, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x42, 0x4c, 0x45, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72,
	0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6d, 0x61,
	0x6a, 0x6f, 0x72, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_BLELocation_proto_rawDescOnce sync.Once
	file_BLELocation_proto_rawDescData = file_BLELocation_proto_rawDesc
)

func file_BLELocation_proto_rawDescGZIP() []byte {
	file_BLELocation_proto_rawDescOnce.Do(func() {
		file_BLELocation_proto_rawDescData = protoimpl.X.CompressGZIP(file_BLELocation_proto_rawDescData)
	})
	return file_BLELocation_proto_rawDescData
}

var file_BLELocation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BLELocation_proto_goTypes = []interface{}{
	(*BLELocation)(nil), // 0: BLELocation
}
var file_BLELocation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BLELocation_proto_init() }
func file_BLELocation_proto_init() {
	if File_BLELocation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BLELocation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLELocation); i {
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
			RawDescriptor: file_BLELocation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BLELocation_proto_goTypes,
		DependencyIndexes: file_BLELocation_proto_depIdxs,
		MessageInfos:      file_BLELocation_proto_msgTypes,
	}.Build()
	File_BLELocation_proto = out.File
	file_BLELocation_proto_rawDesc = nil
	file_BLELocation_proto_goTypes = nil
	file_BLELocation_proto_depIdxs = nil
}
