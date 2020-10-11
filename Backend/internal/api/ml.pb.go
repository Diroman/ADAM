// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: ml.proto

package api

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type Classes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes map[string]float32 `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Classes) Reset() {
	*x = Classes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ml_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Classes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Classes) ProtoMessage() {}

func (x *Classes) ProtoReflect() protoreflect.Message {
	mi := &file_ml_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Classes.ProtoReflect.Descriptor instead.
func (*Classes) Descriptor() ([]byte, []int) {
	return file_ml_proto_rawDescGZIP(), []int{1}
}

func (x *Classes) GetClasses() map[string]float32 {
	if x != nil {
		return x.Classes
	}
	return nil
}

var File_ml_proto protoreflect.FileDescriptor

var file_ml_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x76, 0x0a, 0x07, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0x2a, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x1b, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x06, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ml_proto_rawDescOnce sync.Once
	file_ml_proto_rawDescData = file_ml_proto_rawDesc
)

func file_ml_proto_rawDescGZIP() []byte {
	file_ml_proto_rawDescOnce.Do(func() {
		file_ml_proto_rawDescData = protoimpl.X.CompressGZIP(file_ml_proto_rawDescData)
	})
	return file_ml_proto_rawDescData
}

var file_ml_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ml_proto_goTypes = []interface{}{
	(*Image)(nil),   // 0: Image
	(*Classes)(nil), // 1: Classes
	nil,             // 2: Classes.ClassesEntry
}
var file_ml_proto_depIdxs = []int32{
	2, // 0: Classes.classes:type_name -> Classes.ClassesEntry
	0, // 1: CarDetector.predict:input_type -> Image
	1, // 2: CarDetector.predict:output_type -> Classes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ml_proto_init() }
func file_ml_proto_init() {
	if File_ml_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ml_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_ml_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Classes); i {
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
			RawDescriptor: file_ml_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ml_proto_goTypes,
		DependencyIndexes: file_ml_proto_depIdxs,
		MessageInfos:      file_ml_proto_msgTypes,
	}.Build()
	File_ml_proto = out.File
	file_ml_proto_rawDesc = nil
	file_ml_proto_goTypes = nil
	file_ml_proto_depIdxs = nil
}
